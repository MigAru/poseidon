package upload

//TODO: вынести отдельный контроллер на upload

type Manager struct {
	uploads *Uploads
	buses   []*Bus
}

type Uploads struct {
	unsafe map[string]*Upload
}

type Upload struct {
}

type Metadata struct {
}

//Bus - for partitioned bus
type Bus struct {
	id      string
	channel chan Metadata
}
