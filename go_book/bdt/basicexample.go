package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("basic example demo golang")

	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var r rune = 127
	fmt.Println(r, r+1, r*r)

	medals := []string{"gold", "silver", "bronze"}

	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var p int16 = 5
	var q int32 = 10

	fmt.Println(int32(p) + q)

	fmt.Println("Complex number ==")
	var x = complex(1, 2)
	var y = complex(3, 4)

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(true == true)

	s := "hello , world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])
	fmt.Println("goodbye " + s[:])

	s1 := "left foot"
	t1 := s1
	s1 += ", right foot"

	fmt.Println(s1)
	fmt.Println(t1)
	// not allowed as string are immutable s[0]='L'
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basenamestring("a/b/de.go"))
	fmt.Println(comma("1234567"))
	s3 := "abc"
	b := []byte(s3)
	s4 := string(b)
	fmt.Println(s4)

	fmt.Println("conversion between String and Numbers")

	x1 := 123
	y1 := fmt.Sprintf("%d", x1)
	fmt.Println(y1, strconv.Itoa(x1))

	e := fmt.Sprintf("x1=%b", x1)
	fmt.Println(e)

	const (
		pi = 3.14159
		e1 = 12493802
	)

	fmt.Println(pi, e1)

	const (
		a2 = 1
		b2
		c2 = 3
		d2
	)

	fmt.Println(a2, b2, c2, d2)

	//fmt.Println(BasicT())

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thrusday
		Friday
		Saturday
	)
	fmt.Println(Sunday)

	var arr [3]int

	var arr1 = []int{3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	fmt.Println(arr[len(arr)-1])
	arr[0] = 2
	arr[1] = 3
	fmt.Println(arr[0])

	for i, v := range arr {
		fmt.Printf("key= %d value= %d\n", i, v)
	}

	qa := [...]int{6, 7, 8}
	fmt.Printf("%T\n", qa)
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basenamestring(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
