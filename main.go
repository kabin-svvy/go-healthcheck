package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kabin-svvy/go-healthcheck/report"
	"github.com/spf13/viper"
)

var (
	fullPath = ""
)

func main() {

	v, err := getConfig("config", "./")
	if err != nil {
		log.Fatalf("load config failed %v", err)
	}

	h := report.Handler{
		URL:         v.GetString("report.uri"),
		AccessToken: v.GetString("report.access_token"),
	}

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

	if req.TotalWebsites > 0 {
		res, status, err := h.SendReport(req)
		log.Printf("%+v", res)
		if err != nil {
			log.Fatalf("send report error %v", err)
		}
		if status != http.StatusOK {
			log.Printf("send report failure with status %v", status)
		}
		log.Printf("send report success")
	}
}

func getConfig(fileName string, filePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(fileName)
	v.AddConfigPath(filePath)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
