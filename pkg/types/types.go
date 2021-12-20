package types

// Money представляет собой денежную сумму в минимальных единицах (цнеты, копейки, дирамы и т.д.).
type Money int64

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, ретораны и т.д.).
type Category string

// Status представляет собой статус платежа.
type Status string

// Предопределённые статусы платежей.
const(
	StatusOk Status ="OK"
	StatusFail Status="FAIL"
	StatusInProgress Status="INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct{
	ID 	int
	Amount Money
	Category Category
	Status Status
}