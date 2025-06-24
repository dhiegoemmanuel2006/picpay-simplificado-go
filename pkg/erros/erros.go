package erros

import "errors"

var (
	ErrUserNotFound                   = errors.New("User not found")
	ErrLojistaNotCanSendMoney         = errors.New("Lojista not can send money")
	ErrPayerDontHaveSufficientBalance = errors.New("Payer Dont Have Sufficient Balance")
	ErrTranferUnauthorized            = errors.New("Transfer unauthorized")
	ErrThisRoleIsNotAllowed           = errors.New("This role is not allowed")
)
