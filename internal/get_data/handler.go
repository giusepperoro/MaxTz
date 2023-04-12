package get_data

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandlerGetData(provider DataProvider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		records := provider.GetRecords()

		rawData, err := json.Marshal(records)
		if err != nil {
			log.Println("unable to marshall provided data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(rawData)
		w.WriteHeader(http.StatusOK)
	}
}
