package sites_data_service

import (
	"encoding/json"
	"fmt"
	"internal/models"
	"internal/sorting"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
)

var sitesData []models.SiteDataOutput

func GetSitesData(baseUrl string, catalogId, siteId int, wg *sync.WaitGroup) {
	//msg := fmt.Sprintf("GetSitesData. Start for catalogId=%d, siteId=%d. Start", catalogId, siteId)
	//log.Println(msg)

	url := fmt.Sprintf("%s/%d/%d", baseUrl, catalogId, siteId)
	response, err := http.Get(url)
	if err == nil {
		var sites models.SitesRaw
		if response.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err == nil {
				err = json.Unmarshal(bodyBytes, &sites)
				if err == nil {
					fillDataForResults(sites)
				} else {
					log.Fatal("GetSitesData. Failed to unmarshal the JSON")
				}
			} else {
				log.Fatal("GetSitesData. Failed to read response body")
			}
		}
	} else {
		errMsg := fmt.Sprintf("GetSitesData. Failed to invoke URL: %s", url)
		log.Fatal(errMsg)
	}

	//msg = fmt.Sprintf("GetSitesData for catalogId=%d, siteId=%d. End", catalogId, siteId)
	//log.Println(msg)

	wg.Done()
}

func fillDataForResults(sites models.SitesRaw) {
	for _, mapping := range sites.MappingList {
		sitesData = append(sitesData, models.SiteDataOutput{SiteId: mapping.SiteID, CatalogId: mapping.CatalogId, CategoryName: mapping.CategoryName})
	}
}

func PrintResults() {
	log.Println("Start printing results...")
	sort.Sort(sorting.SiteDataSorter(sitesData))
	for _, siteData := range sitesData {
		output := fmt.Sprintf("Site %d - Catalog %d is mapped to category %s", siteData.SiteId, siteData.CatalogId, siteData.CategoryName)
		log.Println(output)
	}
	log.Println("Printing results was completed successfully")
}

func ReadConfiguration(filename string) (models.Configuration, error) {
	log.Println("Start init configuration file")
	file, err := os.Open(filename)
	if err != nil {
		return models.Configuration{}, err
	}
	decoder := json.NewDecoder(file)
	var configuration models.Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		return models.Configuration{}, err
	}
	log.Println("Init of the configuration file was completed successfully")
	return configuration, nil
}
