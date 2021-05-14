package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var name string
	var userRating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi,Please Enter Your Full Name")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please Give Us Rating From 1 To 5")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strnconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("%v,%v", name, myNumRating)

}
