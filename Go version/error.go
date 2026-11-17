package main

import "fmt"

type Diagnostics struct {
	Line int
	Ref  string
	Msg  string
}


func AppendDiagnostics(allErrors []Diagnostics, line int, ref, msg string) []Diagnostics {
	allErrors = append(allErrors, Diagnostics{
		Line: line,
		Ref:  ref,
		Msg:  msg,
	})

	return allErrors
}

func ReportDiagnostics(allErrors []Diagnostics) {
	for _, e := range allErrors {
		fmt.Printf("Error at line : %v \n near here: ..... %v \n\n | \n%v\n", e.Line, e.Ref, e.Msg)
	}
}
