package v2

type List struct {
	SchemaVersion string
	MediaType     string
	Manifests     []Config
}

type Config struct {
	MediaType string
	Digest    string
	Size      int
	Platform  Platform
}

type Platform struct {
	Architecture string
	OS           string
	OSVersion    string
	OSFeatures   []string
	Variant      string
	Features     string
}
