package dana

type PayMethodEnum int

const (
	Balance PayMethodEnum = iota
	Coupon
	NetBanking
	CreditCard
	DebitCard
	VirtualAccount
	Otc
	DirectDebitCreditCard
	DirectDebitDebitCard
)

func (p PayMethodEnum) String() string {
	return [...]string{
		"BALANCE",
		"COUPON",
		"NET_BANKING",
		"CREDIT_CARD",
		"DEBIT_CARD",
		"VIRTUAL_ACCOUNT",
		"OTC",
		"DIRECT_DEBIT_CREDIT_CARD",
		"DIRECT_DEBIT_DEBIT_CARD"}[p]
}

type ActorTypeEnum int

const (
	User ActorTypeEnum = iota
	Merchant
	MerchantOperator
	BackOffice
	System
)

func (a ActorTypeEnum) String() string {
	return [...]string{
		"USER",
		"MERCHANT",
		"MERCHANT_OPERATOR",
		"BACK_OFFICE",
		"SYSTEM"}[a]
}

type RefundDestinationEnum int

const (
	ToBalance RefundDestinationEnum = iota
	ToSource
)

func (r RefundDestinationEnum) String() string {
	return [...]string{
		"TO_BALANCE",
		"TO_SOURCE"}[r]
}
