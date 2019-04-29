package profiling

import "fmt"

type Printer struct{ I interface{} }

func (p Printer) String() string {
	return "Got to go "
}

type FmtPrinter struct{ I interface{} }

func (p FmtPrinter) String() string {
	return "Got to go " + fmt.Sprint(p.I)
}
