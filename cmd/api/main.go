package main

import (
	"fmt"
	"log"
	"to-do-go/pkg/server"
)

func main() {
	r := server.SetupRouter()
	port := ":8080"

	fmt.Println("le serveur est en cours d'exécution sur http://localhost" + port)

	if err := r.Run(port); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur :", err)
	}
}
