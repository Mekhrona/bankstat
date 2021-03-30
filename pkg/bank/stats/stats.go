package stats

import "github.com/Mekhrona/bankstat/pkg/bank/types"

//Avg расчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money{
	numberOfpayments:=0
	sum:=0
	
	for _, payment := range payments {
		numberOfpayments++
		sum+=int(payment.Amount)	
	}

	avg:=sum/numberOfpayments

	return types.Money(avg)
}