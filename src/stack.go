package main

import (
	"errors"
	"fmt"
)


type Stack struct {
	values []interface{}

}

func (stack Stack) size() int{
	return len(stack.values)
}

func (stack Stack) empty() bool{
	return stack.size() == 0
}

func (stack *Stack) stacking(value interface{}) {
	stack.values = append(stack.values, value)	
}

func (stack *Stack) unstack() (interface{}, error) {
	if stack.empty() {
		return nil, errors.New("stack empty!")
	}
	value := stack.values[stack.size()-1]
	stack.values = stack.values[:stack.size()-1]
	return value, nil
}

func main() {
	stack := Stack{}
	
	fmt.Println("Stack created with size ", stack.size())
	fmt.Println("Empty? ", stack.empty())

	stack.stacking("GO ")
	stack.stacking(2020)
	stack.stacking("FURGBOT")
	stack.stacking("ABzin")

	fmt.Println("Stack after 4 values ", stack.size())
	fmt.Println("Empty? ", stack.empty())

	for !stack.empty() {
		v, _ := stack.unstack()
		fmt.Println("unstack ", v)
		fmt.Println("size ", stack.size())
		fmt.Println("empty ", stack.empty())

	}

	_, err := stack.unstack()
	if err != nil{
		fmt.Println("Erro: ",err)
	}

}
