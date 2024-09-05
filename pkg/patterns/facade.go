package patterns

import "fmt"

// Account struct
type Account struct{
	id string
	accountType string
}

// Account class method
func (account *Account) create(accountType string) *Account {
	fmt.Println(":: Creating account with type ",accountType)
	account.accountType = accountType
	return account
}
func (account *Account) getById(id string) *Account {
	fmt.Println(":: Getting account by id :: ",id)
	return account
}
func (account *Account) deleteById(id string) {
	fmt.Println(":: Deleting account by id :: ",id)
}

// Customer struct
type Customer struct{
	name string
	id int
}
func (customer *Customer) create(name string) *Customer {
	fmt.Println(":: Creating new customer")
	customer.name = name
	customer.id = 1
	return customer
}
// Transaction struct
type Transaction struct{
	id int
	amount float32
	srcAccountId string
	destAccountId string
}

func (transaction *Transaction) create(amount float32,srcAccountId,destAccountId string)  *Transaction {
	fmt.Println(":: Creating new transaction")
	transaction.id = 1
	transaction.amount = amount
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	return transaction
}

// BranchManagerFacad struct
type BranchManagerFacad struct{
	account *Account
	customer *Customer
	transaction *Transaction
}

// method NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacad {
	return &BranchManagerFacad{new(Account),new(Customer),new(Transaction)}
}

// BranchManagerFacade method
func (facade *BranchManagerFacad) createCustomerAccount(customerName, accountType string) (*Customer,*Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType) 
	return customer,account
}
func (facade *BranchManagerFacad) createTransaction(srcAccountId string,
	destAccountId string, amount float32) *Transaction {
		transaction := facade.transaction.create(amount,srcAccountId,destAccountId)
		return transaction
	}

func FacadeFunc() {
	facade := NewBranchManagerFacade()
	customer := new(Customer)
	account := new(Account)
	customer,account = facade.createCustomerAccount("Emad Lakhbizi","Saving")
	fmt.Println(customer.name)
    fmt.Println(account.accountType)
	transaction := facade.createTransaction("123","321",1122)
	fmt.Println(transaction.id)
	fmt.Println(transaction.amount)

}