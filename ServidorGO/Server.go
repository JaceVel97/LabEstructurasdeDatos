package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//Struct con una variable dentro
type Response struct {
	//el json:"respuesta", establece la manera en la cual se construira el json.
	Respuesta string `json:"respuesta"`
}

//Struct con dos variables dentro
type Informacion struct {
	//El omitempty ignora el valor si viene vacio.
	Nombre string `json:"nombre,omitempty"`
	NumeroFavorito int32 `json:"favorito,omitempty"`
}

//Creación de variable para uso dentro de los gets y posts
var Datos Informacion

//Función que devuelve un texto en formato JSON
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	//Variable de tipo Response, el cual es un struct
	var res Response
	//Asignacion del valor
	res.Respuesta="Hello World"

	json.NewEncoder(w).Encode(res)
}

//Función que devuelve una variable en formato JSON
func GetData(w http.ResponseWriter, req *http.Request) {
	//Obtencion de parametros del POST
	json.NewEncoder(w).Encode(Datos)
}

//Función que establece el valor de una variable
func SetData(w http.ResponseWriter, req *http.Request) {
	_ = json.NewDecoder(req.Body).Decode(&Datos)
	json.NewEncoder(w).Encode("Recibido")
}

func main() {
	//Creacion y asignación corta de un nuevo enrutador denominado router
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/getHello", HelloWorld).Methods("GET")

	router.HandleFunc("/getData", GetData).Methods("GET")

	router.HandleFunc("/setData", SetData).Methods("POST")

	//El listenandserve, crea el servidor en el puerto que se establece, siendo en este caso el puerto 7700
	//El log.Fatal se utiliza para visualizar si ocurre un error al iniciar el servidor
	log.Fatal(http.ListenAndServe(":7000", router))
}