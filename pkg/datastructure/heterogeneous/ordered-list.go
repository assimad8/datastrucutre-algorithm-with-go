package heterogeneous

import (
	"fmt"
	"sort"
)

// class Employee
type Employee struct {
	Name string
	ID string
	SSN int
	Age int
}

// ToSorting method 
func (e *Employee) ToSorting() string {
	return fmt.Sprintf("%s: %d,%s,%d",e.Name,e.ID,e.Age,e.SSN)
}
// SortByAge type
type SortByAge []Employee
// SortByAge interface method
func (sortIntf SortByAge) Len() int {return len(sortIntf)}
func (sortIntf SortByAge) Swap(i,j int) { 
	sortIntf[i],sortIntf[j] = sortIntf[j],sortIntf[i]
}
func (sortIntf SortByAge) Less(i,j int) bool {
	return sortIntf[j].Age > sortIntf[i].Age
} 

func OrderFunc() {
	emplyees := []Employee{
		{"Graham","231",235643,31},
        {"John", "3434",245643,42},
        {"Michael","8934",32432, 17},
        {"Jenny", "24334",32444,26},
	}
	fmt.Println(emplyees)
	sort.Sort(SortByAge(emplyees))
	fmt.Println(emplyees)
    sort.Slice(emplyees, func(i int, j int) bool {
        return emplyees[i].Age > emplyees[j].Age
    })
    fmt.Println(emplyees)
}

