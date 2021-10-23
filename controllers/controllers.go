package controllers

import (
	"encoding/json"
	"fmt"
	"miRest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func InitController(resW http.ResponseWriter, req *http.Request) {

	jsonEnconder := json.NewEncoder(resW)
	reponseModel := models.ResponseModel{}

	switch req.Method {
	case "GET":
		reponseModel.Title = "REALIZADO METODO -GET-"
		reponseModel.Message = "Has realizado un metodo GET correctamente"
	case "POST":
		reponseModel.Title = "REALIZADO METODO -POST-"
		reponseModel.Message = "Has realizado un metodo POST correctamente"
	default:
		reponseModel.Title = "Metodo Invalido -" + req.Method
		reponseModel.Message = "Este endpoint solo acepta metodos GET y POST"

	}

	jsonEnconder.Encode(reponseModel)

	fmt.Println(reponseModel.Title)

}

func PathParameterController(resW http.ResponseWriter, req *http.Request) {

	pathParam := mux.Vars(req)["data"]

	responseModel := models.ResponseModel{
		Title:   "Metodo GET con PATH PARAMETER",
		Message: "Has enviado - " + pathParam,
	}

	json.NewEncoder(resW).Encode(responseModel)

	fmt.Println("Metodo get con Path Parameters")

}

func QueryParameters(resW http.ResponseWriter, req *http.Request) {

	queryParams := req.URL.Query()

	jsonEncoder := json.NewEncoder(resW)
	responseModel := models.ResponseModel{}

	data, present := queryParams["miQuery"]

	soloData := queryParams.Get("miQuery2")
	//Get devuelve el valor del query en String si acaso no esta devuelve un string vacio

	if present == true {

		responseModel.Title = "Metodo GET con Query Parameters"
		responseModel.Message = "enviado miQuery - valor: " + data[0] +
			"- enviado miQuery2: " + soloData

		jsonEncoder.Encode(responseModel)

		return

	} else if soloData != "" {
		responseModel.Title = "Metodo GET con Query Parameters"
		responseModel.Message = "enviado miQuery2 - valor: " + soloData

		jsonEncoder.Encode(responseModel)
		return
	}

	responseModel.Title = "Metodo GET sin Query Parameters"
	responseModel.Message = "- miQuery o miQuery2- no fueron enviados"

	jsonEncoder.Encode(responseModel)

	fmt.Println("Metodo Get - Query Parameters")

}

func PostBodyJson(resW http.ResponseWriter, req *http.Request) {

	jsonDecoder := json.NewDecoder(req.Body)
	jsonEncoder := json.NewEncoder(resW)

	requestBody := models.RequestBody{}

	responseModel := models.ResponseModel{}

	err := jsonDecoder.Decode(&requestBody)

	if err != nil {
		responseModel.Title = "Solictud Post invalida"
		responseModel.Message = "Debes enviar un Json Valido"

		jsonEncoder.Encode(responseModel)

		fmt.Println(responseModel.Title)

		return
	}

	responseJsonModel := models.ResponseModelJson{}

	responseJsonModel.Title = "Solicitud POST enviada correctamente"
	responseJsonModel.Message = "has enviado un json valido"
	responseJsonModel.Data = requestBody

	jsonEncoder.Encode(responseJsonModel)

	fmt.Println(responseJsonModel.Title + "-" + responseJsonModel.Message)
}

func PostFile(resW http.ResponseWriter, req *http.Request) {

	responseModel := models.ResponseModel{}

	jsonEnconder := json.NewEncoder(resW)

	req.ParseMultipartForm(2 << 20)

	file, header, err := req.FormFile("miFile")

	if err != nil {

		responseModel.Title = "Solicitud Post Incorrecta"
		responseModel.Message = "Archivo no enviado correctamente -miFile-"
		jsonEnconder.Encode(responseModel)

		fmt.Println(responseModel.Title)

		return

	}
	defer file.Close()

	responseModel.Title = "Has subido un archivo correctamente"
	responseModel.Message = "Has subido el archivo : -" + header.Filename

	jsonEnconder.Encode(responseModel)

	fmt.Println(responseModel.Message)

}
