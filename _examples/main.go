// +build windows
package main

import (
	actionCenter "github.com/RedHatQE/gowinx/pkg/app/action-center"
)

func main() {
	// Click action center
	actionCenter.Click()
}
