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
	inspectedCount  big.Int
}

func (m Monkey) ToString() string {
	return fmt.Sprintf(`Monkey
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

func MakeMonkeys(filename string) (monkeyList []Monkey) {
	file, err := os.Open(filename)
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

func (m *Monkey) DoOneTurn(monkeys *[]Monkey, reliefModifier int64, compress bool) {
	zeroBig := big.NewInt(0)
	mmd := GetMultiMonkeyDivisor(*monkeys)
	for i := range m.items {
		bigOne := big.NewInt(1)
		m.inspectedCount.Add(&m.inspectedCount, bigOne)
		// update worry level
		switch {
		case m.inspectOperand == "old":
			m.items[i] = *m.items[i].Mul(&m.items[i], &m.items[i])
		case m.inspectOperator == "*":
			operand, _ := strconv.Atoi(m.inspectOperand)
			operandBig := big.NewInt(int64(operand))
			m.items[i].Mul(&m.items[i], operandBig)
		case m.inspectOperator == "+":
			operand, _ := strconv.Atoi(m.inspectOperand)
			operandBig := big.NewInt(int64(operand))
			m.items[i].Add(&m.items[i], operandBig)
		}

		// calculate worry relief
		if !compress {
			relief := big.NewInt(reliefModifier)
			m.items[i].Div(&m.items[i], relief)
		} else if compress {
			relief := big.NewInt(reliefModifier)
			bigRelief := relief.Mul(relief, &mmd)
			m.items[i].Mod(&m.items[i], bigRelief)
		}

		// check if worry level is divisible by test operand
		rem := big.NewInt(0)
		rem.Rem(&m.items[i], &m.testOperand)
		switch {
		case rem.Cmp(zeroBig) == 0:
			// assign item to true monkey
			(*monkeys)[m.testTrueDest].items = append((*monkeys)[m.testTrueDest].items, m.items[i])
		case rem.Cmp(zeroBig) != 0:
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

func GetLevelOfMonkeyBusiness(monkeys []Monkey) (lomb *big.Int) {
	l1, l2 := big.NewInt(0), big.NewInt(0)
	for _, monkey := range monkeys {
		switch {
		case monkey.inspectedCount.Cmp(l1) == 1:
			l2.Set(l1)
			l1.Set(&monkey.inspectedCount)
		case monkey.inspectedCount.Cmp(l2) == 1:
			l2.Set(&monkey.inspectedCount)
		}
	}
	lomb = l2.Mul(l2, l1)
	return lomb
}

func GetMultiMonkeyDivisor(monkeys []Monkey) (mmd big.Int) {
	mmd = *big.NewInt(1)
	for _, monkey := range monkeys {
		mmd.Mul(&mmd, &monkey.testOperand)
	}
	return mmd
}

func DoRounds(monkeys []Monkey, numRounds int, reliefFactor int, compress bool, debug bool) (lomb *big.Int) {
	for i := 0; i < numRounds; i++ {
		if debug {
			fmt.Println("round", i)
		}
		for j := range monkeys {
			monkeys[j].DoOneTurn(&monkeys, int64(reliefFactor), compress)
		}
	}
	lomb = GetLevelOfMonkeyBusiness(monkeys)
	if debug {
		PrintMonkeyStatus(monkeys)
		PrintMonkeyInspectedCounts(monkeys)
		fmt.Println("level of monkey business", lomb)
	}
	return lomb
}
