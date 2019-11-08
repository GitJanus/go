package main

import (
	"fmt"
	"github.com/ZZMarquis/gm/sm3"
)

func main() {
	var Data = "111111"
	dataByte := []byte(Data)
	digest := sm3.New()
	digest.Update(dataByte, 0)
	fmt.Println(sm3.Sum(dataByte))
}
