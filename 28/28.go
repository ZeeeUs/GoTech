package main

import "fmt"

type client struct {
}

type card interface {
	pay()
}

type sber struct {
}

type city struct {}

type cityAdapter struct {
	cityCard *city
}

func (c *client) payByCard(visa card) {
	fmt.Println("Клиент рассчитывается")
	visa.pay()
}

func (s *sber) pay() {
	fmt.Println("Оплачиваем рублями")
}

func (cit *city) pay(){
	fmt.Println("Оплачиваем долларами")
}

func (a *cityAdapter) pay() {
	fmt.Println("Конвертация валюты из рублей в доллары")
	a.cityCard.pay()
}

func main(){
	client := &client{}
	sber := &sber{}

	client.payByCard(sber)

	cityCard := &city{}
	cityCardAdapter := &cityAdapter{
		cityCard: cityCard,
	}

	client.payByCard(cityCardAdapter)
}

