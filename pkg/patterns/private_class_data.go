package patterns

import (
	"encoding/json"
	"fmt"
)

// AccountDetail struct
type AccountDetail struct{
	id string
	accountType string
}
// Account struct
type PAccount struct{
// Private attribute becouse it's name start with small letter  
	details *AccountDetail
	CustomerName string
}
// PAcount class method setDetails
func (account *PAccount) setDetails(id ,accountType string) {
	account.details = &AccountDetail{id,accountType}
}
func (account *PAccount) getId() string {
	return account.details.id
}
func (account *PAccount) getAccountType() string {
	return account.details.accountType
}

func PrivateFunc() {
	account := &PAccount{CustomerName:"Emad Lakhbizi"}
	account.setDetails("21","current")
	jsonAccount,_ := json.Marshal(account)
	fmt.Println("Private Class hidden",string(jsonAccount))
	fmt.Println("Account ID::",account.getId())
	fmt.Println("Account TYPE::",account.getAccountType())
}