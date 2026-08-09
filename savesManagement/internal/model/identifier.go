package model

import (
	"errors"
	"strings"

	fy "didockerf/savesManagement/internal/model/fileType"
)

const (
	defaultTag = "1.0"
)

type Identifier struct {
	Name     string
	Tag      string
	FileType fy.FileType
}

func CreateIdentifier(name string, tag string, fileType fy.FileType) (Identifier, error) {
	if strings.Trim(name, " ") == "" {
		return Identifier{}, errors.New("name has a invalid value")
	}
	if strings.Trim(tag, " ") == "" {
		tag = defaultTag
	}

	return Identifier{
		Name:     strings.Trim(name, " "),
		Tag:      strings.Trim(tag, " "),
		FileType: fileType,
	}, nil
}

func CreateIdentifierFromStr(identifierStr string, fileType fy.FileType) (Identifier, error) {
	if !IsIdentifierStrValid(identifierStr) {
		return Identifier{}, errors.New("The passed identifier is invalid")
	}

	identifierStrSplitted := strings.Split(identifierStr, identifierStrSeparetor)
	if isOnlyName(identifierStrSplitted) {
		identifierStrSplitted = append(identifierStrSplitted, defaultTag)
	}

	name := identifierStrSplitted[0]
	tag := identifierStrSplitted[1]

	return Identifier{
		Name:     name,
		Tag:      tag,
		FileType: fileType,
	}, nil
}

func isOnlyName(identifierStrSplitted []string) bool {
	return len(identifierStrSplitted) == 1
}
