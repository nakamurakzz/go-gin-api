package site

import "github.com/nakamurakzz/go-gin-api/db"

type Site struct {
	ID        int
	Name      string
	IsEnabled bool
}

type SiteRepository interface {
	FindAll() ([]*SiteEntity, error)
	FindByID(id int) (*SiteEntity, error)
}

type SiteRepositoryImpl struct {
}

type SiteQuery struct {
	isEnabled bool
}

func (s *SiteRepositoryImpl) FindAll(siteQuery SiteQuery) ([]*SiteEntity, error) {
	conn := db.GetDBConn().DB
	var sites []*Site
	if err := conn.Where("isEnabled = ?", siteQuery.isEnabled).Find(&sites).Error; err != nil {
		return nil, err
	}
	var returnSites []*SiteEntity
	for _, site := range sites {
		returnSites = append(returnSites, NewSiteEntity(
			SiteEntity{
				ID:        site.ID,
				Name:      site.Name,
				IsEnabled: site.IsEnabled,
			},
		))
	}
	return returnSites, nil
}

func (s *SiteRepositoryImpl) FindByID(id int) (*SiteEntity, error) {
	conn := db.GetDBConn().DB
	var site Site
	if err := conn.Where("id = ?", id).Find(&site).Error; err != nil {
		return nil, err
	}
	var returnSite = NewSiteEntity(
		SiteEntity{
			ID:        site.ID,
			Name:      site.Name,
			IsEnabled: site.IsEnabled,
		},
	)
	return returnSite, nil
}
