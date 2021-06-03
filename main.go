package main

import (
	"fmt"

	"github.com/adrianriobo/gowinx/pkg/ux"
)

const (
	NIOW_CLASS string = "NotifyIconOverflowWindow"
)

func main() {
	fmt.Print("hello world")
	ux.GetWindowHandlerByClass(NIOW_CLASS)
}
