package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func EvenOrOdd(num int) (string, error) {
	if num < 0 {
		return "negative", errors.New("negative number passed")
	}
	if num%2 == 0 {
		return "even", nil
	} else {
		return "odd", nil
	}
}

func main() {
	result, err := EvenOrOdd(-9)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)

	numbers := [4]int{1, 2, 3, 4}
	itemSaleCount := map[string]int{"Item A": 12, "Item B": 56}

	for key, value := range itemSaleCount {
		fmt.Println(key, value)
	}

	for item := range numbers {
		fmt.Println(numbers[item])
	}
	username := "Raghu Ballu"
	fullName := strings.Split(username, " ")
	fmt.Print(reflect.TypeOf(fullName).Kind(), fullName)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
