package fileType

type FileType struct {
	filePrefix   string
	saveLocation string
}

var allFileTypes = []FileType{
	GetDockerfileType(),
	GetComposeFileType(),
}

func (fileType FileType) GetFilePrefix() string {
	return fileType.filePrefix
}

func (fileType FileType) GetSaveLocation() string {
	return fileType.saveLocation
}

func GetDockerfileType() FileType {
	return FileType{
		filePrefix:   "dockerfile",
		saveLocation: "dockerfiles",
	}
}

func GetComposeFileType() FileType {
	return FileType{
		filePrefix:   "composefile",
		saveLocation: "compose-files",
	}
}

func GetAllFileTypes() []FileType {
	return allFileTypes
}

func GetFileTypeOfFilePrefix(filePrefix string) *FileType {
	for _, fileType := range allFileTypes {
		if filePrefix == fileType.GetFilePrefix() {
			return &fileType
		}
	}
	return nil
}
