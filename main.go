package main

import (
	"fmt"
	"time"

	"github.com/zees-dev/go-dom/models"
)

func main() {
	dom := models.GetExampleDom()

	element, err := dom.GetElementByIDViaGoRoutines("copyright")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(element)

	element, err = dom.GetElementByIDViaWaitGroup("copyright")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(element)

	element, err = dom.GetElementByIDBFS("copyright")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(element)

	element, err = dom.GetElementByIDDFS("copyright")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(element)

	time.Sleep(200 * time.Millisecond)
}
