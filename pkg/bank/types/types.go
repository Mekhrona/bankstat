package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет собой категорию, в которой был совершён платёж
type Category string


//Payment представляет информацию о платеже.
type Payment struct{
	ID int
	Amount Money
	Category Category
}