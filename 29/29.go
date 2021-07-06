package main

import (
	"fmt"
	"math/big"
)

func plus(x, y *big.Int) *big.Int{
	return big.NewInt(0).Add(x, y)
}

func minus(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func div(x, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}

func mult(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func main(){
	var x big.Int
	var y big.Int

	x.SetInt64(1 << 62)
	y.SetInt64(1 << 23)

	fmt.Println("Сложени: ", plus(&x, &y))
	fmt.Println("Вычитание: ", minus(&x, &y))
	fmt.Println("Деление: ", div(&x, &y))
	fmt.Println("Умножение: ", mult(&x, &y))
}
