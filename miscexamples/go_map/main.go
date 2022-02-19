package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ffgacv",
		"green": "#e345fgh",
		"blue":  "#rt145gg",
	}
	fmt.Println(colors)
	deleteItem("green", colors)
	print(colors)

}

func deleteItem(key string, m map[string]string) {
	delete(m, key)
	fmt.Println("delete successful")
}

func print(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, "value is ", v)
	}

}
