package core

// Stub is an example domain type
type Stub struct {
	ID string `json:"id,omitempty"`
}

// NewStubs is an example factory
func NewStubs() []Stub {
	return []Stub{
		{ID: "a"},
		{ID: "b"},
		{ID: "c"},
	}
}
