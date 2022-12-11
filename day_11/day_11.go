package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	id              int
	items           []big.Int
	inspectOperator string
	inspectOperand  string
	testOperand     big.Int
	testTrueDest    int
	testFalseDest   int
	inspectedCount  int
}

func (m Monkey) ToString() string {
	return fmt.Sprintf(`Monkey\n
id: %v
items: %v
inspectOperator: %v
inspectOperand: %v
testOperand: %v
testTrueDest: %v
testFalseDest %v
inspectedCount %v`,
		m.id, m.items,
		m.inspectOperator,
		m.inspectOperand,
		m.testOperand,
		m.testTrueDest,
		m.testFalseDest,
		m.inspectedCount,
	)
}

func makeMonkeys(filname string) (monkeyList []Monkey) {
	file, err := os.Open(filname)
	if err != nil {
		fmt.Println("Error: Bad Monkeys")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tmpMonkey := Monkey{}

	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)
		switch {
		case line == "":
			monkeyList = append(monkeyList, tmpMonkey)
			tmpMonkey = Monkey{}
		case lineFields[0] == "Monkey":
			id, _ := strconv.Atoi(strings.TrimSuffix(lineFields[1], ":"))
			tmpMonkey.id = id
		case lineFields[0] == "Starting":
			items := make([]big.Int, 0)
			for _, item := range lineFields[2:] {
				item := strings.TrimSuffix(item, ",")
				itemValue, _ := strconv.Atoi(item)
				itemBig := big.NewInt(int64(itemValue))
				items = append(items, *itemBig)
			}
			tmpMonkey.items = items
		case lineFields[0] == "Operation:":
			tmpMonkey.inspectOperator = lineFields[4]
			tmpMonkey.inspectOperand = lineFields[5]
		case lineFields[0] == "Test:":
			testOperandValue, err := strconv.Atoi(lineFields[3])
			if err != nil {
				fmt.Println(err)
			}
			testOperandBig := big.NewInt(int64(testOperandValue))
			tmpMonkey.testOperand = *testOperandBig
		case strings.HasPrefix(line, "    If true"):
			testTrueDest, _ := strconv.Atoi(lineFields[5])
			tmpMonkey.testTrueDest = testTrueDest
		case strings.HasPrefix(line, "    If false"):
			testFalseDest, _ := strconv.Atoi(lineFields[5])
			tmpMonkey.testFalseDest = testFalseDest
		}

	}
	monkeyList = append(monkeyList, tmpMonkey)
	return monkeyList
}

func (m *Monkey) DoOneTurn(monkeys *[]Monkey) {
	zeroBig := big.NewInt(0)
	for i := range m.items {
		m.inspectedCount++
		// update worry level
		switch {
		case m.inspectOperand == "old":
			m.items[i] = *m.items[i].Mul(&m.items[i], &m.items[i])
		case m.inspectOperator == "*":
			operand, _ := strconv.Atoi(m.inspectOperand)
			operandBig := big.NewInt(int64(operand))
			if m.items[i].Mul(&m.items[i], operandBig).Cmp(zeroBig) < 0 {
				fmt.Println("before")
				fmt.Println(m.ToString())
				fmt.Println("item:", m.items[i])
				fmt.Println("operand:", operand)
			}

			m.items[i].Mul(&m.items[i], operandBig)
			if m.items[i].Cmp(zeroBig) < 0 {
				fmt.Println("illegal monkeybusiness!")
				fmt.Println("after")
				fmt.Println(m.ToString())
				panic("cannot abide bad monkey business")
			}
		case m.inspectOperator == "+":
			operand, _ := strconv.Atoi(m.inspectOperand)
			operandBig := big.NewInt(int64(operand))
			m.items[i].Add(&m.items[i], operandBig)
		}
		// calculate worry relief (worry / 3)
		// m.items[i] /= 10
		// check if worry level is divisible by test operand
		rem := big.NewInt(0)
		rem.Rem(&m.items[i], &m.testOperand)
		fmt.Printf("val %v, testOperand %v, remainder %v\n", m.items[i], m.testOperand, rem)
		switch {
		case rem == zeroBig:
			fmt.Println("assigning true monkey")
			// assign item to true monkey
			(*monkeys)[m.testTrueDest].items = append((*monkeys)[m.testTrueDest].items, m.items[i])
		case rem != zeroBig:
			fmt.Println("assigning false monkey")

			// assign to false monkey's items
			(*monkeys)[m.testFalseDest].items = append((*monkeys)[m.testFalseDest].items, m.items[i])
		}
	}
	m.items = []big.Int{}
}

func PrintMonkeyStatus(monkeys []Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %v: %v\n\n", monkey.id, monkey.items)
	}
}

func PrintMonkeyInspectedCounts(monkeys []Monkey) {
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %v inspected items %v times\n\n", monkey.id, monkey.inspectedCount)
	}
}

func GetLevelOfMonkeyBusiness(monkeys []Monkey) (lomb int) {
	l1, l2 := 0, 0
	for _, monkey := range monkeys {
		if l1 < monkey.inspectedCount {
			l2 = l1
			l1 = monkey.inspectedCount
		} else if l2 < monkey.inspectedCount {
			l2 = monkey.inspectedCount
		}
	}
	return l1 * l2
}

func main() {
	monkeys := makeMonkeys("small_input.txt")
	for i := 0; i < 20; i++ {
		for j, _ := range monkeys {
			monkeys[j].DoOneTurn(&monkeys)
		}
	}

	// PrintMonkeyStatus(monkeys)
	PrintMonkeyInspectedCounts(monkeys)
	fmt.Println(GetLevelOfMonkeyBusiness(monkeys))
	// 2637590098 too low for large input

}
