package api

type CreateAccountRequest struct {
	Username      string `json:"username"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
}

type Account struct {
	Username      string `json:"username"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	PasswordToken []byte `json:"-"`
	PasswordSalt  []byte `json:"salt"`
}

func makeAccount(values CreateAccountRequest) *Account {
	return &Account{
		Username:      values.Username,
		FirstName:     values.FirstName,
		LastName:      values.LastName,
		PasswordToken: []byte("testingtoken"),
		PasswordSalt:  []byte("testingsalt"),
	}
}
