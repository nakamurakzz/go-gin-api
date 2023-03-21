package site

type SiteUsecase struct {
}

var siteRepository = SiteRepositoryImpl{}

func (s *SiteUsecase) FindAll() ([]*SiteEntity, error) {
	sites, err := siteRepository.FindAll(SiteQuery{
		isEnabled: true,
	})
	if err != nil {
		return nil, err
	}
	return sites, nil
}

func (s *SiteUsecase) FindByID(id int) (*SiteEntity, error) {
	site, err := siteRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return site, nil
}

func (s *SiteUsecase) Create(siteReq SiteCreateRequestBody) (*SiteEntity, error) {
	site := createSiteEntity(SiteCreateParam{
		Name: siteReq.Name,
	})
	newSite, err := siteRepository.Create(site)
	if err != nil {
		return nil, err
	}
	return newSite, nil
}

func (s *SiteUsecase) Delete(id int) (int, error) {
	targetSite, err := siteRepository.FindByID(id)
	if err != nil {
		return 0, err
	}
	if targetSite == nil {
		return 0, nil
	}
	err = siteRepository.Delete(id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
