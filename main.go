package main

import (
	"backend/handlers"
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	// Inicializa la configuración de Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(".env")
	config := &firebase.Config{ProjectID: "test-5eebf"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// Inicializa el cliente de Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Cerrar el cliente de Firestore cuando ya no se necesite
	defer client.Close()

	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "UP and running...") // imprime la respuesta del cliente
	})

	router.HandleFunc("/reset-password", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		handlers.ResetPassword(email)
	}).Methods("POST")

	log.Println("Server listening on port", port) // imprime en el servidor
	log.Fatal(http.ListenAndServe(port, router))
}
