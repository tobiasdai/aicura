package nexus

const (
	ComponentsPath = "/components"
)

type ComponentResponse struct {
	Items             []Component `json:"items,omitempty"`
	ContinuationToken interface{} `json:"continuationToken,omitempty"`
}

// Component is the base structure for Nexus Component
type Component struct {
	ID         string `json:"id,omitempty"`
	Repository string `json:"repository,omitempty"`
	Format     string `json:"format,omitempty"`
	Group      string `json:"group,omitempty"`
	Name       string `json:"name,omitempty"`
	Version    string `json:"version,omitempty"`
}
