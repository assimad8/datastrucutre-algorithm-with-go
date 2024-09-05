package patterns

import "fmt"

// DataTransferObjectFactory class 
type DataTransferObjectFactory struct {
	pool map[string] DataTransferObject
}

func(factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	dto := factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: "+dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = FCustomer{id:"1"}
		case "employee":
			factory.pool[dtoType] = FEmployee{id:"2"}
		case "manager":
			factory.pool[dtoType] = FManager{id:"3"}
		case "address":
			factory.pool[dtoType] = FAddress{id:"4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// DataTransferObject interface
type DataTransferObject interface{
	getId() string
}
// Customer struct
type FCustomer struct{
	id string
	name string
	ssn string 
}
func (customer FCustomer) getId() string {
	return customer.id
}

// FEmployee struct
type FEmployee struct{
	id string
	name string
}
func (employee FEmployee) getId() string {
	return employee.id
}
// FManager struct
type FManager struct{
	id string
	name string
	dept string
}
func (manager FManager) getId() string {
	return manager.id
}
type FAddress struct {
	id string
	streetLine1 string
	streetLine2 string
	state string
	city string
}
func (address FAddress) getId() string {
	return address.id
}

func FliweightFunc() {
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	customer := factory.getDataTransferObject("customer")
	fmt.Println("Customer: ",customer.getId())
	employee := factory.getDataTransferObject("employee")
	fmt.Println("Employee: ",employee.getId())
	manager := factory.getDataTransferObject("manager")
	fmt.Println("Manager: ",manager.getId())
	address := factory.getDataTransferObject("address")
	fmt.Println("Address: ",address.getId())
}