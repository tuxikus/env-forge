// Package status prints information about the status of the program
package status

import (
	"fmt"
)

const (
	StatusCreatingFile = iota
	StatusDoneCreatingFile
	StatusWritingFile
	StatusDoneWritingFile
	StatusErrCreatingFile
)

var _ StatusPrinter = (*statusPrinter)(nil)

type StatusPrinterStatus int

type StatusPrinter interface {
	Print(StatusPrinterStatus, string)
}

type statusPrinter struct{}

func NewStatusPrinter() *statusPrinter {
	return &statusPrinter{}
}

// TODO: make variadic
func (sp *statusPrinter) Print(status StatusPrinterStatus, text string) {
	switch status {
	case StatusCreatingFile:
		fmt.Print("Creating file:", text+"...")
	case StatusWritingFile:
		fmt.Print("Writing file:", text+"...")
	case StatusDoneCreatingFile, StatusDoneWritingFile:
		fmt.Print("done ✅\n")
	case StatusErrCreatingFile:
		fmt.Print("Error creating file:", text+"...")
	}
}
