package main

import (
	"debug/pe"
	"fmt"
	"io"
	"os"
)

func dumpSymbols(out io.Writer, symbols []*pe.Symbol) {
	for _, symbol1 := range symbols {
		fmt.Fprintf(out, "%s=%d\n", symbol1.Name, symbol1.Value)
	}
}

func dumpSections(out io.Writer, sections []*pe.Section) {
	for _, section1 := range sections {
		fmt.Fprintf(out, "Name=%s\n", section1.Name)
	}
}

func dumpCoffSymbols(out io.Writer, symbols []pe.COFFSymbol) error {
	var st pe.StringTable
	for _, symbol1 := range symbols {
		name, err := symbol1.FullName(st)
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "%s\n", name)
	}
	return nil
}

func doFile(fname string) error {
	file1, err := pe.Open(fname)
	if err != nil {
		return err
	}
	defer file1.Close()

	if oh := file1.OptionalHeader.(*pe.OptionalHeader64); oh != nil {
		fmt.Printf("[optionalHeader64]\n%d.%d.%d.%d\n",
			oh.MajorImageVersion,
			oh.MinorImageVersion,
			oh.MajorSubsystemVersion,
			oh.MinorSubsystemVersion)
	} else if oh := file1.OptionalHeader.(*pe.OptionalHeader32); oh != nil {
		fmt.Printf("[optionalHeader32]\n%d.%d.%d.%d\n",
			oh.MajorImageVersion,
			oh.MinorImageVersion,
			oh.MajorSubsystemVersion,
			oh.MinorSubsystemVersion)
	}
	if false {
		fmt.Println("[symbols]")
		dumpSymbols(os.Stdout, file1.Symbols)
		fmt.Println("[sections]")
		dumpSections(os.Stdout, file1.Sections)
		fmt.Println("[COFFSymbols]")
		if err := dumpCoffSymbols(os.Stdout, file1.COFFSymbols); err != nil {
			return err
		}
	}
	return nil
}

func main1() error {
	for _, arg1 := range os.Args[1:] {
		if err := doFile(arg1); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := main1(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
