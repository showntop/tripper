package models

type Spot struct {
	BaseModel
	Title       string `jsonapi:"attr,title"`
	Description string `jsonapi:"attr,description"`
}

func NewSpot(title, description string) *Spot {
	return &Spot{
		Title:       title,
		Description: description,
	}
}
