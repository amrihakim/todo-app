package response

type GlobalResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`	
}

type TitleResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Title   string `json:"title"`
}
