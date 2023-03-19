package site

import "github.com/nakamurakzz/go-gin-api/db"

type Site struct {
	ID        int
	Name      string
	IsEnabled bool
}

type SiteRepository interface {
	FindAll() ([]*Site, error)
	FindByID(id int) (*Site, error)
}

type SiteRepositoryImpl struct {
}

func (s *SiteRepositoryImpl) FindAll() ([]*Site, error) {
	conn := db.GetDBConn().DB
	var sites []*Site
	if err := conn.Where("isEnabled = ?", true).Find(&sites).Error; err != nil {
		return nil, err
	}
	return sites, nil
}

func (s *SiteRepositoryImpl) FindByID(id int) (*Site, error) {
	conn := db.GetDBConn().DB
	var site Site
	if err := conn.Where("id = ?", id).Find(&site).Error; err != nil {
		return nil, err
	}
	return &site, nil
}
