package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	id              int
	items           []int
	inspectOperator string
	inspectOperand  string
	testOperand     int
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
		m.id, m.items, m.inspectOperator, m.inspectOperand, m.testOperand, m.testTrueDest, m.testFalseDest, m.inspectedCount,
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
			items := make([]int, 0)
			for _, item := range lineFields[2:] {
				item := strings.TrimSuffix(item, ",")
				itemValue, _ := strconv.Atoi(item)
				items = append(items, itemValue)
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
			tmpMonkey.testOperand = testOperandValue
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

	for i := range m.items {
		m.inspectedCount++
		// update worry level
		switch {
		case m.inspectOperand == "old":
			m.items[i] = m.items[i] * m.items[i]
		case m.inspectOperator == "*":
			operand, _ := strconv.Atoi(m.inspectOperand)
			m.items[i] *= operand
		case m.inspectOperator == "+":
			operand, _ := strconv.Atoi(m.inspectOperand)
			m.items[i] += operand
		}
		// calculate worry relief (worry / 3)
		m.items[i] /= 3
		// check if worry level is divisible by test operand
		switch {
		case m.items[i]%m.testOperand == 0:
			// assign item to true monkey
			(*monkeys)[m.testTrueDest].items = append((*monkeys)[m.testTrueDest].items, m.items[i])
		case m.items[i]%m.testOperand != 0:
			// assign to false monkey's items
			(*monkeys)[m.testFalseDest].items = append((*monkeys)[m.testFalseDest].items, m.items[i])
		}
	}
	m.items = []int{}
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
	// for _, monkey := range monkeys[:1] {
	// 	fmt.Println(monkey.ToString())
	// 	fmt.Println()
	// }
	// fmt.Println(monkeys[0].ToString())
	// monkeys[0].DoOneTurn(&monkeys)
	// fmt.Println()
	// fmt.Println(monkeys[0].ToString())
	// fmt.Println()
	// fmt.Println(monkeys[3].ToString())
	for i := 0; i < 20; i++ {
		for j, _ := range monkeys {
			monkeys[j].DoOneTurn(&monkeys)
		}
	}

	PrintMonkeyStatus(monkeys)
	PrintMonkeyInspectedCounts(monkeys)
	fmt.Println(GetLevelOfMonkeyBusiness(monkeys))
}
