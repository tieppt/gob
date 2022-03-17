package dto

type GenericErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
