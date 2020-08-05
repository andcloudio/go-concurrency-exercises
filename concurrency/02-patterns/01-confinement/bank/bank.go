package bank

//TODO: modify the implementation to share balance amount by communication

var balance int

func Deposit(amount int) { balance += amount }

func Withdrawal(amount int) { balance -= amount }

func Balance() int { return balance }
