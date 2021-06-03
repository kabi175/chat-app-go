package repository

type service struct {
	datasource Datasource
}

func NewService(datasource Datasource) service {
	return service{
		datasource: datasource,
	}
}
