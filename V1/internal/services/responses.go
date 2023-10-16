package services

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
