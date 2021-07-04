package main

import (
	"encoding/binary"
	"fmt"
)

func main(){
	dNum := int64(1230)
	i := 1 // индекс бита, который надо заменить
	buf := make([]byte, binary.MaxVarintLen64)

	binary.LittleEndian.PutUint64(buf, uint64(dNum))
	fmt.Printf("%+v\n", buf)

	fmt.Printf("Исходное число в 10: %d \nИсходное число в 2: %b \n", dNum, dNum)

	if buf[i] == 1 {
		buf[i] = 0
	} else {
		buf[i] = 1
	}

	fmt.Printf("%+v\n", buf)

	newDecimal := int64(binary.LittleEndian.Uint64(buf))

	fmt.Printf("Новое число в 10: %d \nНовое чилсо в 2: %b \n", newDecimal, newDecimal)
}
