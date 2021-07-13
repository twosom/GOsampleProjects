package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}


var errNoMoney = errors.New("Not Enough Balance.")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	//== &를 붙여 주소값을 반환하고 최종 반환값은 *를 붙여 그 주소값으로 만들어진 객체를 반환한다.
	//== 객체를 "복사" 하는것이 아니다. "주소값" 을 반환하여 그것으로 객체를 만드는 것이다. ==//
	return &account
}

// Deposit x amount on your account
//== receiver 를 통해 객체의 값을 변경하고자 할 때는,
//== receiver 의 type 앞에 *을 붙여준다.
//== 그렇지 않으면 receiver 를 받는 변수에 객체가 "복사" 되어 전달된다.==//
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

