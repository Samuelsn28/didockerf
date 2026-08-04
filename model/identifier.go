package model

import (
	"errors"
	"strings"
)

const (
	defaultTag = "1.0"
)

type Identifier struct {
	Name string
	Tag  string
}

func (id Identifier) GetIdentifierStr() IdentifierStr {
	return IdentifierStr(id.Name + identifierStrSeparetor + id.Tag)
}

func CreateIdentifier(name string, tag string) (Identifier, error) {
	if strings.Trim(name, " ") == "" {
		return Identifier{}, errors.New("name has a invalid value")
	}
	if strings.Trim(tag, " ") == "" {
		tag = defaultTag
	}

	return Identifier{
		Name: strings.Trim(name, " "),
		Tag:  strings.Trim(tag, " "),
	}, nil
}

func CreateIdentifierFromStr(identifierStr IdentifierStr) (Identifier, error) {
	if !IsIdentifierStrValid(identifierStr) {
		return Identifier{}, errors.New("the passed identifier is invalid")
	}

	identifierStrSplitted := identifierStr.getIdentifierSplitted()
	if isOnlyName(identifierStrSplitted) {
		identifierStrSplitted = append(identifierStrSplitted, defaultTag)
	}

	name := identifierStrSplitted[0]
	tag := identifierStrSplitted[1]

	return Identifier{
		Name: name,
		Tag:  tag,
	}, nil
}

func isOnlyName(identifierStrSplitted []string) bool {
	return len(identifierStrSplitted) == 1
}
