package main

import (
	"log"
	"main/AddressBookServer/controllers/stdhttp"
	"main/AddressBookServer/gate/psg"
	"net/http"
)

func main() {

	db, err := psg.NewPsg("localhost:5432", "postgres", "123")
	if err != nil {
		log.Fatal("Error creating database connection:", err)
	}

	controller := stdhttp.NewController("localhost:8080", db)

	http.HandleFunc("/record/add", controller.RecordAdd)
	http.HandleFunc("/records/get", controller.RecordsGet)
	http.HandleFunc("/record/update", controller.RecordUpdate)
	http.HandleFunc("/record/delete", controller.RecordDeleteByPhone)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
