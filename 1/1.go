package main

import "fmt"

type Human struct {
	Age int
	Action Action
}

type Action struct {
	Read bool
	Swim bool
	RideBike bool
}

func (a *Human) Bday () *Human {
	a.Age++
	return a
}

func (act *Human) Learn (action string) *Human {
	switch action {
	case "Read":
		act.Action.Read = true
	case "Swim":
		act.Action.Swim = true
	case "RideBike":
		act.Action.RideBike = true
	default:
		fmt.Println("This Action is not support")
	}
	return act
}

func main() {
	action := Action{Read: true, Swim: false, RideBike: true}
	Ivan := Human{Age: 23, Action: action}
	fmt.Printf("%+v \n", Ivan)
	Ivan.Bday()
	fmt.Printf("%+v \n", Ivan)
	Ivan.Learn("Swim")
	fmt.Printf("%+v \n", Ivan)
	Vasya := Human{Age: 18, Action: Action{Read: true, Swim: true, RideBike: false}}
	fmt.Printf("%+v \n", Vasya)
}
