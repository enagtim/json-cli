package api

type API struct {
	Key string
}

func NewAPI(key string) *API {
	return &API{
		Key: key,
	}
}
