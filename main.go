package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Respuesta struct {
	Saludo string `json:Saludo`
}

type AllRespuesta []Respuesta

var leer = AllRespuesta{
	{
		Saludo: "Hola todos",
	},
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(leer)
}

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")
	http.ListenAndServe(":8083", nil)
}
