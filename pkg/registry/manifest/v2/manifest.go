package v2

type Manifest struct {
	SchemaVersion string
	MadiaType     string
	Config        ManifestConfig
}

type ManifestConfig struct {
	MadiaType string
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
