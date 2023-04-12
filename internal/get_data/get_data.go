package get_data

import (
	"io"
	"log"
	"os"

	gocsv "github.com/trimmer-io/go-csv"

	"github.com/giusepperoro/MaxTz/internal/entity"
)

type DataProvider struct {
	records []entity.Record
}

func (d *DataProvider) GetRecords() []entity.Record {
	return d.records
}

func NewDataProvider(filename string) (DataProvider, error) {
	var d DataProvider
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("unable to open file with data: %v", err)
	}

	rawData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("unable to read from file: %v", err)
	}

	d.records = make([]entity.Record, 0)
	err = gocsv.Unmarshal(rawData, &d.records)

	return d, nil
}
