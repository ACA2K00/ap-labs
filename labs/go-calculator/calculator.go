package main

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
import(
	"fmt"
	"os"
	"strconv"
)

func calc(operator int, values []int) int {
	var res int
	if operator == 1{
		res = 0
		for i := 0; i < len(values); i++{
			res += values[i]
		}
	}else if operator == 2{
		res = values[0]
		for i := 1; i < len(values); i++{
			res -= values[i]
		}
	}else{
		res = 1
		for i := 0; i < len(values); i++{
			res *= values[i]
		}
	}
	return res
}

func main() {
	strOperator := os.Args[1]
	var r int
	var operator int
	var nValues int = len(os.Args) - 2
	values := make([]int, nValues)
	k := 0


	for i := 2; i < len(os.Args); i++ {
		intVar, err := strconv.Atoi(os.Args[i])
		//values = append(values, intVar)
		values[k] = intVar
		k++
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}

	if strOperator == "add"{
		operator = 1
	}else if strOperator == "sub"{
		operator = 2
	}else if strOperator == "mult"{
		operator = 3
	}else{
		fmt.Printf("Invalid Operator: %s\n", strOperator)
			os.Exit(3)
	}

	r = calc(operator, values)
	fmt.Printf("Result: %v\n", r)

	/*
	fmt.Printf("Values: %v\n", values)
	fmt.Printf("String Operator: %v\n", strOperator)
	fmt.Printf("Operator: %v\n", operator)
	fmt.Printf("inputs: %v\n", len(os.Args))*/
}
