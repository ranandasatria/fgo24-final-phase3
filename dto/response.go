package dto

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Results any `json:"results,omitempty"`
	Errors any `json:"errors,omitempty"`
}