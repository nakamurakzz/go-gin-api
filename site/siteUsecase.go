package site

type SiteUsecase interface {
	FindAll() ([]*Site, error)
	FindByID(id int) (*Site, error)
}

type SiteUsecaseImpl struct {
}

var siteRepository = SiteRepositoryImpl{}

func (s *SiteUsecaseImpl) FindAll() ([]*SiteEntity, error) {
	sites, err := siteRepository.FindAll(SiteQuery{
		isEnabled: true,
	})
	if err != nil {
		return nil, err
	}
	return sites, nil
}

func (s *SiteUsecaseImpl) FindByID(id int) (*SiteEntity, error) {
	site, err := siteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return site, nil
}
