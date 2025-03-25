package transactions

type TransactionType string

const (
	TransactionTypeWidthdraw = "W"
	TransactionTypeDeposit   = "D"
	TransactionTypeInterest  = "I"
)
