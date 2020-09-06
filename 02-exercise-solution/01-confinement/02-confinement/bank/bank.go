package bank

// TODO: modify the implementation to share balance amount by memory sharing (mutex).

var balance int

func Deposit(amount int) { balance += amount }

func Withdrawal(amount int) { balance -= amount }

func Balance() int { return balance }
