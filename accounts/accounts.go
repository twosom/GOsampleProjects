package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	//== &를 붙여 주소값을 반환하고 최종 반환값은 *를 붙여 그 주소값으로 만들어진 객체를 반환한다.
	//== 객체를 "복사" 하는것이 아니다. "주소값" 을 반환하여 그것으로 객체를 만드는 것이다. ==//
	return &account
}
