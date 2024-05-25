package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EtherreumPrice struct {
	Price float64
	Data  time.Time
}

func LoadDataFrom(path string) ([]EtherreumPrice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error reading question file: %w", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	var pricePairs []EtherreumPrice
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading from CSV: %w", err)
		}
		parsedData, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Data: %w", err)
		}
		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse price: %w", err)
		}
		pricePairs = append(pricePairs, EtherreumPrice{
			Price: parsedPrice,
			Data:  parsedData,
		})
	}
	return pricePairs, nil
}
