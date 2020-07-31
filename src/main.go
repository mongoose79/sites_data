package main

import (
	"fmt"
	"internal/sites_data_service"
	"io"
	"log"
	"os"
	"sync"
)

const logFile = "sites_data_task.log"

func main() {
	initLog()
	filename := "src\\internal\\config\\tsconfig.json"
	conf, err := sites_data_service.ReadConfiguration(filename)
	if err != nil {
		log.Fatal("Failed to read configuration file")
		return
	}

	log.Println("Starting sites data task...")
	var wg sync.WaitGroup
	for _, siteId := range conf.SiteIds {
		for i := 1; i < conf.TotalCatalogsCount+1; i++ {
			wg.Add(1)
			go sites_data_service.GetSitesData(conf.BaseServerURL, i, siteId, &wg)
		}
	}
	wg.Wait()
	sites_data_service.PrintResults()
	fmt.Println("Report Done")
}

func initLog() {
	fmt.Println("Start initializing the log")
	logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Failed to create log file")
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
