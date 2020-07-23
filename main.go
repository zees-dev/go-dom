package main

import (
	"fmt"
	"time"

	"github.com/zees-dev/go-dom/models"
)

func main() {
	dom := models.GetExampleDom()

	element, err := dom.GetElementByID("copyright")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(element)

	element, err = dom.GetElementByIDBFS("copyright")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(element)

	element, err = dom.GetElementByIDDFS("copyright")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(element)

	time.Sleep(200 * time.Millisecond)
}
