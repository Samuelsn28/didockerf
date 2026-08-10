package savesManagement

import fy "didockerf/savesManagement/internal/model/fileType"

func GetAllSavedComposeFilesInfos() []SavedFileInfo {
	return getAllSavedFilesOf(fy.GetComposeFileType())
}
