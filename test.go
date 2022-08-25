package main

import "fmt"

func main()  {

	fmt.Println("Start.......")
	var abc ABC
	abc.Name = "NAME1"
	abc.Value = "VALUE1"
	fmt.Println(abc.getName("-----subname"))
}

type ABC struct {
	Name string
	Value string
}

func (abc *ABC)getName(xyz string) string {
	return abc.Name + xyz
}
