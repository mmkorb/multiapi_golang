package domain

type Limit struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func NewLimit(id int, name string, value float64) *Limit {
	return &Limit{
		ID:    id,
		Name:  name,
		Value: value,
	}
}
