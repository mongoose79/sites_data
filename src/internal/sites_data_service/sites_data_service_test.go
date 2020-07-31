package sites_data_service

import (
	"github.com/stretchr/testify/assert"
	"internal/models"
	"internal/sorting"
	"sort"
	"sync"
	"testing"
)

var conf models.Configuration

func init() {
	fileName := "C:\\workspace\\sites_data\\src\\internal\\config\\tsconfig.json"
	conf, _ = ReadConfiguration(fileName)
}

func TestInitConfig(t *testing.T) {
	assert.NotNil(t, conf)
	assert.Equal(t, "http://localhost:8080/catalogmapping/getcatalogmapping/json", conf.BaseServerURL)
	assert.Equal(t, []int{0, 2, 3, 77}, conf.SiteIds)
	assert.Equal(t, 200, conf.TotalCatalogsCount)
}

func TestGetSitesData_SingleSiteAndCatalog(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	GetSitesData(conf.BaseServerURL, 1, conf.SiteIds[0], &wg)
	sort.Sort(sorting.SiteDataSorter(sitesData))

	assert.Equal(t, 4, len(sitesData))
	assert.Equal(t, 0, sitesData[0].SiteId)
	assert.Equal(t, 1, sitesData[0].CatalogId)
	assert.Equal(t, "DVDs & Blu-ray Discs", sitesData[0].CategoryName)
}

func TestGetSitesData_AllSites(t *testing.T) {
	var wg sync.WaitGroup
	for _, siteId := range conf.SiteIds {
		for i := 1; i < conf.TotalCatalogsCount-1; i++ {
			wg.Add(1)
			go GetSitesData(conf.BaseServerURL, i, siteId, &wg)
		}
	}
	wg.Wait()
	sort.Sort(sorting.SiteDataSorter(sitesData))

	assert.Equal(t, 1923, len(sitesData))
}
