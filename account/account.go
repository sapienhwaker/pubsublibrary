package account

import "go/types"

type Account struct {
	AccountNo string
	AccountType types.BankAccountType
} 