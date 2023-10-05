package types

type BankAccountType int

const (
        unknown BankAccountType = iota
        saving
        current
)