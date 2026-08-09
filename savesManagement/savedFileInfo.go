package savesManagement

import (
	"didockerf/savesManagement/internal/model"
)

type SavedFileInfo struct {
	name string
	tag  string
}

func (info SavedFileInfo) GetName() string {
	return info.name
}

func (info SavedFileInfo) GetTag() string {
	return info.tag
}

func createSavedFileInfoOf(savedFile model.SavedFile) SavedFileInfo {
	return SavedFileInfo{
		name: savedFile.Identifier.Name,
		tag:  savedFile.Identifier.Tag,
	}
}
