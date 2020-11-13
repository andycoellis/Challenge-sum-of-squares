package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // Pop the number of problems involved

	if err != nil {
		log.Fatalf("ERROR IN STDIN [%s]", err)
	}
	problemSize, _ := strconv.Atoi(strings.TrimSpace(input)) //Trim any '\n' attached
	bytes, err := ioutil.ReadAll(reader) // Store all remaining problem strings into one byte array

	if err != nil {
		log.Fatalf("ERROR AT READALL [%s]", err)
	}
	var result = process(problemSize, strings.Split(string(bytes), "\n"), 0)

	fmt.Println(result)
}

func process(size int, problem [] string, count int) string{

	nSize, _ := strconv.Atoi(problem[2*count])	//Returns the size of indv. problem
	allInputs := problem[(2*count) + 1]	//We return the index which contains problem set
	inputs := strings.Split(allInputs, " ")	//Create a splice of problem set 

	sum := sumOfSquares(inputs, nSize, 0) // Function found below

	if count == size-1 {

		return strconv.Itoa(sum)
	}

	return strconv.Itoa(sum) + "\n" + process(size, problem, count+1)
}

func sumOfSquares(problem [] string, size int, count int) int{
	
	number, _ := strconv.Atoi(problem[count])
	var square float64

	if number > 0{

		square = math.Pow(float64(number), 2)

	} else {
		square = 0
	}

	if count == size-1 {

		return int(square)
	}

	return int(square) + sumOfSquares(problem, size, count+1)
}