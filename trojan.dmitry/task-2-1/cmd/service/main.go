package main

import (
	"fmt"
)

type Dept struct {
	minLevel int
	maxLevel int
}

func NewDept(defaultMinLevel, defaultMaxLevel int) Dept {

	return Dept{minLevel: defaultMinLevel, maxLevel: defaultMaxLevel}
}

func (department *Dept) Update(operator string, num int) error {
	switch operator {
	case ">=":
		department.minLevel = max(department.minLevel, num)
		return nil
	case "<=":
		department.maxLevel = min(department.maxLevel, num)
		return nil
	default:
		return fmt.Errorf("Unknown operator")
	}

}

func (department *Dept) Result() int {
	if department.minLevel <= department.maxLevel {
		return department.minLevel
	}

	return -1
}

func main() {
	var department int

	_, err := fmt.Scan(&department)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range department {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		const (
			defaultMinLevel = 15
			defaultMaxLevel = 30
		)

		dept := NewDept(defaultMinLevel, defaultMaxLevel)

		for range workers {
			var operator string

			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Invalid operator")

				return
			}

			var num int

			_, err = fmt.Scan(&num)
			if err != nil {
				fmt.Println("Invalid temperature value")

				return
			}

			err = dept.Update(operator, num)
			if err != nil {
				fmt.Println("Error updating department")
			}

			fmt.Println(dept.Result())
		}
	}
}
