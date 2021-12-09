package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response2 struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
