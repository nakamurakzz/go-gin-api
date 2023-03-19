package site

type SiteEntity struct {
	iD        int
	name      string
	isEnabled bool
}

func NewSiteEntity(param SiteEntity) *SiteEntity {
	return &SiteEntity{
		iD:        param.iD,
		name:      param.name,
		isEnabled: param.isEnabled,
	}
}
