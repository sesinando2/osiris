package kubernetes

// PatchOperation represents a discreet change to be applied to a Kubernetes
// resource
type PatchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
