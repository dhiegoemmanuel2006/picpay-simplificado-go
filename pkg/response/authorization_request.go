package response

type Data struct {
	Authorization bool `json:"authorization"`
}

type ExternAuthorizationResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
