package sorting

import "internal/models"

type SiteDataSorter []models.SiteDataOutput

func (sds SiteDataSorter) Len() int      { return len(sds) }
func (sds SiteDataSorter) Swap(i, j int) { sds[i], sds[j] = sds[j], sds[i] }
func (sds SiteDataSorter) Less(i, j int) bool {
	if sds[i].SiteId == sds[j].SiteId {
		if sds[i].CatalogId == sds[j].CatalogId {
			return sds[i].CategoryName < sds[j].CategoryName
		}
		return sds[i].CatalogId < sds[j].CatalogId
	}
	return sds[i].SiteId < sds[j].SiteId
}
