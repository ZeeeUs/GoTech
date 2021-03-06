package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1 // выделяем в памяти место и по определённому адресу кладём туда 1
		p = &a // переменная p ссылается на тот же адрес в памяти, что и переменная a
	)
	fmt.Println(*p) // указываем на переменную p, который ссылается на переменную а
	update(p) // передаём в функцию переменную, а не ссылку на данную переменную, т.е. в функции создаётся копия переменной p
	fmt.Println(*p)
}
