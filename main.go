package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/isidroevc/monserver/configuration"
	"github.com/isidroevc/monserver/database"
	"github.com/isidroevc/monserver/models"
	"github.com/isidroevc/monserver/services"
)

var upgrader = websocket.Upgrader{}

func Receiver(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	config, getConfigError := configuration.GetConfiguration()
	if getConfigError != nil {
		panic(getConfigError)
	}
	if err != nil {
		log.Print("Could not upgrade connection")
	}
	defer connection.Close()
	for {
		message := new(models.StatMessage)

		err := connection.ReadJSON(message)
		if message.CommunityChain != config.CommunityChain {
			log.Print(config.CommunityChain)
			log.Println("Failed auth with CC: ", message.CommunityChain)
			break
		}
		if err != nil {
			log.Println("read:", err)
			break
		}
		storeService := services.NewStoreService(database.GetConnection())
		err = storeService.UpdateStats(message)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	log.SetFlags(0)
	http.HandleFunc("/", Receiver)
	log.Fatal(http.ListenAndServe(":3000", nil))
	fmt.Println("benever")
}
