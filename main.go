package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type getRecordsRequest struct {
	Id int64 `json:"id"`
}

type read struct {
	rd *csv.Reader
}

func New(reader *csv.Reader) *read {
	return &read{rd: reader}
}

func (rd *read) records(id int64) ([][]string, error) {
	recordId := strconv.FormatInt(id, 10)
	rd.rd.Comment = '#'
	sliceRecordsOut := make([][]string, 0)
	for {
		record, e := rd.rd.Read()
		if e != nil {
			fmt.Println(e)
			return nil, e
		}
		if (record[1]) == recordId {
			sliceRecordsOut = append(sliceRecordsOut, record)
			fmt.Println(record)
		}
	}
	return sliceRecordsOut, nil
}

func (rd *read) getRecords(w http.ResponseWriter, r *http.Request) {
	var req getRecordsRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordToResponse, err := rd.records(req.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	file, err := os.Open("ueba.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	rd := New(reader)
	fmt.Println(rd.records(1534))
	http.HandleFunc("get_records", rd.getRecords)
}
