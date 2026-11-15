package main

import "fmt"

type ErrorObj struct {
	Line int
	Ref string
	Msg string
}

var allErrors []ErrorObj

func AppendError(line int, ref, msg string) {
	allErrors = append(allErrors, ErrorObj{
		Line: line,
		Ref: ref,
		Msg: msg,
	})
}

func ReportErrors() {
	for _,e:= range allErrors {
		fmt.Printf("Error at line : %v \n near here: ..... %v \n\n | \n%v\n", e.Line, e.Ref, e.Msg)
	}
}