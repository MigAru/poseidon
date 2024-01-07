package v2_2

import "encoding/json"

type Manifest struct {
	SchemaVersion int            `json:"schemaVersion"`
	MediaType     string         `json:"mediaType"`
	Config        ManifestConfig `json:"config"`
	Layers        []Layer        `json:"layers"`
}

func (m *Manifest) GetLayersDigests() []string {
	var layers []string
	for _, layer := range m.Layers {
		layers = append(layers, layer.Digest)
	}

	layers = append(layers, m.Config.Digest)

	return layers
}

func (m *Manifest) GetLength() int {
	b, _ := json.Marshal(m)
	return len(b)
}

type ManifestConfig struct {
	MediaType string `json:"mediaType"`
	Digest    string `json:"digest"`
	Size      int    `json:"size,omitempty"`
}

type Layer struct {
	MediaType string   `json:"mediaType"`
	Size      int      `json:"size"`
	Digest    string   `json:"digest"`
	Urls      []string `json:"urls,omitempty"` /*Кейс с urls - очень редкий. TODO: Добавить возможность скачивать с других адресов
	  С возможностью указаний где хранится другой кусок и сколько байтов в этом куске*/
}
