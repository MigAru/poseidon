package v2_2

type Manifest struct {
	SchemaVersion string
	MediaType     string
	Config        ManifestConfig
}

type ManifestConfig struct {
	MediaType string
	Digest    string
	Size      int
	Layers    []Layer
}

type Layer struct {
	MediaType string
	Size      int
	Digest    string
	Urls      []string
}
