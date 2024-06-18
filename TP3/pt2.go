package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func getInput() int {
	var input string
	fmt.Scanln(&input)
	command, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ingrese un numero valido")
		return 0
	}
	return command
}

func producer(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		fmt.Println("Ingrese un numero")
		input := getInput()
		c <- input
		time.Sleep(time.Millisecond * 100)
	}
}

func duplicate(num int) {
	fmt.Printf("Se duplicara el numero %d \n", num)
	num = num * 2
	fmt.Printf("El resultado es %d \n", num)
}

func divide(num int) {
	fmt.Printf("Se dividira en dos el numero %d \n", num)
	num = num / 2
	fmt.Printf("El resultado es %d \n", num)
}

func raiseTo2(num int) {
	fmt.Printf("Se elevara a la potencia 2 el numero %d \n", num)
	numTo2 := math.Pow(float64(num), 2.0)
	fmt.Printf("El resultado es %f \n", numTo2)
}

func consumer(c chan int) {
	process := rand.Intn(3) + 1
	for i := range c {
		switch process {
		case 1:
			duplicate(i)
			time.Sleep(time.Millisecond * 1000)
		case 2:
			divide(i)
			time.Sleep(time.Millisecond * 1000)
		case 3:
			raiseTo2(i)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func main() {
	channel := make(chan int)
	go producer(channel)
	for i := 0; i < 3; i++ {
		go consumer(channel)
	}
	time.Sleep(time.Second * 20)
}
