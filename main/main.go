package main

import (
	"fmt"
	"log"
	"os"
	
	"example.com/operator"
)

func getFloat() (float32, error) {
	var number float32

	fmt.Print("Enter a decimal number: ")
	_, err := fmt.Scanln(&number)
	
	if err != nil {
		return 0, fmt.Errorf("getFloat: invalid decimal number", err)
	}
	
	return number, nil
}

func getOperation() (string, error){
	var op string

	fmt.Print("Enter any sign among the following: '+', '-', '*', '/': ")
	_, err := fmt.Scanf("%s", &op)
	
	if err != nil {
		return "", fmt.Errorf("getOperation: unknown operator: %q", op)
	}
	
	return op, nil
}

func getResult(x float32, op string, y float32) (float32, error){
	switch op {
	case "+":
		return operator.Add(x,y), nil
	case "-":
		return operator.Sub(x,y), nil
	case "/":
		if y == 0{
			return 0, fmt.Errorf("getResult: cannot divide by zero")
		}
		return operator.Div(x, y), nil
	case "*":
		return operator.Mult(x,y), nil
	default:
		return 0, fmt.Errorf("getResult: invalid operation")
	}	
}

func main (){
	fmt.Println("A simple two number Calculator")
	
	x, err := getFloat()
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	
	y, err := getFloat()
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	
	op, err := getOperation()
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	
	result, err := getResult(x, op, y)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("The result of %v %v %v = %v", x, op, y, result)
}
