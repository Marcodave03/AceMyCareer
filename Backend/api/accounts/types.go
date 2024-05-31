package accounts

type ValidateAccountRequest struct {
	Username          string `json:"username"`
	PasswordToken     []byte `json:"password"`
	PasswordSalt      string `json:"salt"`
}

type MakeAccountRequest struct {
	Username          string `json:"username"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	PasswordToken     []byte `json:"password"`
	PasswordSalt      string `json:"salt"`
	ProfilePictureUrl string `json:"imageurl"`
	CartId            int    `json:"-"`
}

type Account struct {
	Username          string `json:"username"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	PasswordToken     []byte `json:"-"`
	PasswordSalt      string `json:"salt"`
	ProfilePictureUrl string `json:"imageurl"`
	CartId            int    `json:"-"`
}

