package main

import "fmt"

func Run() error {
	fmt.Println("startig up our appplication")
	return nil
}

func main() {
	fmt.Println("Go REST API movie")
	if err := Run(); err != nil{
		fmt.Println(err)
	}
}
