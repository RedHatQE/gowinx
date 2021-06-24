// +build windows
package main

import (
	actionCenter "github.com/adrianriobo/gowinx/pkg/app/action-center"
	"github.com/adrianriobo/gowinx/pkg/app/crc"
)

func main() {
	// Simple Click item
	crc.Click([]string{crc.ACTION_STOP})
	// Double click submenu item
	crc.Click([]string{crc.ACTION_COPY_OC_COMMAND, crc.ACTION_COPY_OC_COMMAND_KUBEADMIN})
	// Click action center
	actionCenter.Click()
}
