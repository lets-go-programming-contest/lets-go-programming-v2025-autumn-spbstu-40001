package main

import (
	"fmt"
)

type Dept struct {
	minLevel int
	maxLevel int
}

func NewDept() Dept {
	return Dept{minLevel: 15, maxLevel: 30}
}

func (department *Dept) Update(operator string, num int) {
	switch operator {
	case ">=":
		department.minLevel = max(department.minLevel, num)
	case "<=":
		department.maxLevel = min(department.maxLevel, num)
	}
}

func (department Dept) Result() int {
	if department.minLevel <= department.maxLevel {
		return department.minLevel
	}
	return -1
}

func main() {
	var (
		department int
		workers    int
		num        int
		operator   string
	)

	_, err := fmt.Scan(&department)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range department {
		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		dept := NewDept()

		for range workers {
			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Println("Invalid operator")

				return
			}

			_, err = fmt.Scan(&num)
			if err != nil {
				fmt.Println("Invalid temperature value")

				return
			}

			dept.Update(operator, num)

			fmt.Println(dept.Result())
		}
	}
}
