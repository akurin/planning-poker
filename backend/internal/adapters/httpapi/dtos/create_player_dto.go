package dtos

type CreatePlayerDto struct {
	Name string `json:"name" validate:"required"`
}

func (o CreatePlayerDto) Validate() string {
	if o.Name == "" {
		return "The field name is required."
	}
	return ""
}
