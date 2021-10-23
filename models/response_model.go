package models

type ResponseModel struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

type RequestBody struct {
	Mensaje  string
	Numero   int
	Booleano bool
}

type ResponseModelJson struct {
	ResponseModel
	Data RequestBody `json:"data"`
}
