package main

import (
	// "fmt"
	"fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/nonlinear"
)

func main() {
	table := nonlinear.Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id","Name","SSN"}
	rows := make([]nonlinear.Row,2)
	rows[0] = nonlinear.Row{}
	columns1 := make([]nonlinear.Column,3)
	columns1[0] = nonlinear.Column{1,"123"}
	columns1[0] = nonlinear.Column{1,"Emad lakhbizi"}
	columns1[0] = nonlinear.Column{1,"234234"}
	rows[0].Columns = columns1
	rows[1] = nonlinear.Row{}
	columns2 := make([]nonlinear.Column,3)
	columns2[0] = nonlinear.Column{1,"456"}
	columns2[0] = nonlinear.Column{1,"lakhbizi emad"}
	columns2[0] = nonlinear.Column{1,"667677"}
	rows[1].Columns = columns2
	table.Rows = rows
	fmt.Println(table)
	table.PrintTibale()
}
