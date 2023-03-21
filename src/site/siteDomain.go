package site

type SiteEntity struct {
	ID        int
	Name      string
	IsEnabled bool
}

type SiteCreateParam struct {
	Name string
}

func NewSiteEntity(param SiteEntity) *SiteEntity {
	return &SiteEntity{
		ID:        param.ID,
		Name:      param.Name,
		IsEnabled: param.IsEnabled,
	}
}

func createSiteEntity(param SiteCreateParam) SiteEntity {
	return SiteEntity{
		Name:      param.Name,
		IsEnabled: true,
	}
}
