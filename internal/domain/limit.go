package domain

// Limit represents a limit value for a resource or configuration
type Limit struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// NewLimit creates a new instance of Limit
func NewLimit(id int, name string, value float64) *Limit {
	return &Limit{
		ID:    id,
		Name:  name,
		Value: value,
	}
}
