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
	// element := dom.GetElementByIDBFS("copyright")

	fmt.Println(element)
	time.Sleep(200 * time.Millisecond)
}
