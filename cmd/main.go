package main

import (
	"log"
	"net/http"

	"github.com/giusepperoro/MaxTz/internal/get_data"
)

func main() {
	dataProvider, err := get_data.NewDataProvider("ueba.csv")
	if err != nil {
		log.Fatalf("unable to cretae data provider: %v", err)

	}
	http.HandleFunc("/get_data", get_data.HandlerGetData(dataProvider))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
