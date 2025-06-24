package request

type CreateUserRequest struct {
	FullName string  `json:"fullName"`
	Document string  `json:"document"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	Balance  float64 `json:"balance"`
}
