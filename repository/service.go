package repository

type service struct {
	dataSource DataSource
	jwt        *Jwt
}

func NewService(dataSource DataSource) service {
	return service{
		dataSource: dataSource,
		jwt:        NewJwt("secret"),
	}
}
