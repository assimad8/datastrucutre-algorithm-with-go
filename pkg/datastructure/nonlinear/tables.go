package nonlinear

import "fmt"

// Column class
type Column struct{
	Id int 
	Value string
}

// Row class
type Row struct {
	Columns []Column
	Id int
}
// Talbe class
type Table struct{
	Rows []Row
	Name string
	ColumnNames []string
}

// printTable
func (table Table) PrintTibale() {
	rows := table.Rows
	fmt.Println(table.Name)
	for _,row := range rows {
		columns := row.Columns
		for i,column := range columns {
			fmt.Println(table.ColumnNames[i],column.Id,column.Value)
		}
	}
}