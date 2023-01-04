package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var names = []string{"Gab", "Mike", "jin", "Ren", "Nim"}
	fmt.Println(names)

	fmt.Println("Select the item you wish to remove")
	itemReader := bufio.NewReader(os.Stdin)
	itemInput, _ := itemReader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(itemInput), 8, 8)

	removedItem := names[index]
	fmt.Printf("You removed: %v\n", removedItem)

	names = append(names[:index], names[index+1:]...)

	fmt.Println("New Array is now: ", names)

	var mike User= User{"Mike", "mike@gmail.com", 23, 54}

	fmt.Println(mike)
}

type User struct {
	Name string
	Email string
	Age int
	Num int32
}