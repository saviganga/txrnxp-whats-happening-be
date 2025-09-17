package media

type MediaStorageProvider interface {
	UploadFile(fileName string, fileData []byte) (string, error)
	RetrieveFile(path string) string
}
