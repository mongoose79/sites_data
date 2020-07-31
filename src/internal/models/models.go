package models

type Configuration struct {
	BaseServerURL      string `json:"BaseServerURL"`
	SiteIds            []int  `json:"SiteIds"`
	TotalCatalogsCount int    `json:"TotalCatalogsCount"`
}

type SitesRaw struct {
	InvocationId   string      `json:"invocationId"`
	ResponseStatus string      `json:"responseStatus"`
	Errors         interface{} `json:"errors"`
	MappingList    []Mapping   `json:"mappingList"`
}

type Mapping struct {
	SiteID           int    `json:"siteID"`
	CategoryId       int    `json:"categoryId"`
	CatalogId        int    `json:"catalogId"`
	CategoryName     string `json:"categoryName"`
	CatalogName      string `json:"catalogName"`
	IsCatalogEnabled bool   `json:"isCatalogEnabled"`
	VcsId            int    `json:"vcsId"`
}

type SiteDataOutput struct {
	SiteId       int
	CatalogId    int
	CategoryName string
}
