// +build windows

package main

import (
	"github.com/adrianriobo/gowinx/pkg/crc"
)

func main() {
	crc.Click(crc.ACTION_STOP)
}
