package main

import (
	"fmt"
	
	"github.com/Phogheus/GoLearnGo/greetings"
)

func main() {
	message:= greetings.Hello("Awesome Name")
	fmt.Println(message)
}
