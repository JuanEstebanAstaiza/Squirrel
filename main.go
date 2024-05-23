package main

import (
	"fmt"
	"github.com/JuanEstebanAstaiza/Squirrel/Controllers"
	"github.com/JuanEstebanAstaiza/Squirrel/Utils"
	"github.com/rs/cors"
	"log"
	"net/http"

	_ "github.com/JuanEstebanAstaiza/Squirrel/Controllers"
	"github.com/gorilla/mux"
)

func main() {

	err := Utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/user/register", Controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", Controllers.LoginUser).Methods("POST")
	router.HandleFunc("/api/squire/insert-squirrel/AddSquire", Controllers.AddSquire).Methods("POST")
	router.HandleFunc("/api/squire/view/view-squire/{userID}", Controllers.GetSquiresByUser).Methods("GET")
	router.HandleFunc("/api/squire/edit-squirrel/EditSquire", Controllers.EditSquire).Methods("PUT")
	router.HandleFunc("/api/squire/delete-squirrel/DeleteSquire/{squireID}", Controllers.DeleteSquire).Methods("DELETE")

	// Opciones CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permitir solicitudes desde cualquier origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)
	// Iniciar el servidor
	port := ":7250"
	fmt.Println("Servidor escuchando en el puerto", port)
	log.Fatal(http.ListenAndServe(port, handler))

}
