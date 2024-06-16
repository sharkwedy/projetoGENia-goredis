package service

type ObjectService interface {
	FetchObjects() ([]Object, error)
}

// Object is a representation of the JSON object from the external API.
type Object struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data *Data  `json:"data"`
}

type Data struct {
	Color    string `json:"color"`
	Capacity string `json:"capacity"`
}
