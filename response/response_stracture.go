package response

type Response struct {
	Status  int         `Json:"status"`
	Message string      `Json:"message"`
	Data    interface{} `json:"data"`
}
