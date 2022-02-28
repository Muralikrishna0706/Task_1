package model

type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Mobile_Number string `json:"mobile_number"`
	Email     string    `json:"email"`
	Address      string `json:"address"`
	Education  string `json:"education"`
}
type Response struct {
	Status int         `json:"status"`
	Error  bool        `json:"error"`
	Data   interface{} `json:"data"`
}
