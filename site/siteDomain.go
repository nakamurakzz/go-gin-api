package site

type SiteEntity struct {
	ID        int
	Name      string
	IsEnabled bool
}

func NewSiteEntity(param SiteEntity) *SiteEntity {
	return &SiteEntity{
		ID:        param.ID,
		Name:      param.Name,
		IsEnabled: param.IsEnabled,
	}
}
