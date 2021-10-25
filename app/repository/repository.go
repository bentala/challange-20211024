package repository

//go:generate mockgen -destination=repository/mock_data.go -package=repository . Data
type Data interface {
	Get(words string, page int) (interface{}, error)
}

type Records interface {
	Add(event string, data interface{}) error
}
