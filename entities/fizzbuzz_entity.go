package entities

type FizzBuzzRequest struct {
	Count uint `json:"count,omitempty" validate:"required"`
}

type FizzBuzzResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
