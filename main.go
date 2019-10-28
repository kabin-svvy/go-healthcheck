package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kabin-svvy/go-healthcheck/report"
)

var (
	fullPath = "test.csv"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fullPath = args[1]
	}

	csvfile, err := os.Open(fullPath)
	if err != nil {
		log.Fatalln("failure to open csv,", err)
	}

	r := csv.NewReader(csvfile)

	req := report.Request{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		if len(record) > 0 {
			fmt.Printf("uri %v\n", record[0])
			req.HealthCheck(record[0])
		}
	}
	fmt.Printf("report healthcheck %+v\n", req)
}
