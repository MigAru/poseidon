package upload

type Upload struct {
	ID            string
	TotalSize     int
	UploadedBytes int
	ProjectName   string
	Queue         [][]byte
}

func NewUpload(id, projectName string, totalSize int) *Upload {
	return &Upload{
		ID:          id,
		ProjectName: projectName,
		TotalSize:   totalSize,
		Queue:       make([][]byte, 0),
	}
}
