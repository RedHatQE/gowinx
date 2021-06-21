// +build windows
package main

import (
	"github.com/adrianriobo/gowinx/pkg/crc"
)

func main() {
	// Simple Click item
	crc.Click([]string{crc.ACTION_STOP})
	// Double click submenu item
	crc.Click([]string{crc.ACTION_COPY_OC_COMMAND, crc.ACTION_COPY_OC_COMMAND_KUBEADMIN})
}
