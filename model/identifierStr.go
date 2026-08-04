package model

import (
	"strings"

	"didockerf/util"
)

const (
	identifierStrSeparetor = ":"

	nameFormatRegex       = `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*`
	tagFormatRegex        = `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*`
	identifierFormatRegex = `^` + nameFormatRegex + identifierStrSeparetor + tagFormatRegex + `$`
)

type IdentifierStr string

func (idstr IdentifierStr) GetIdentifier() Identifier {
	identifier, errToGet := CreateIdentifierFromStr(idstr)

	if errToGet != nil {
		return Identifier{}
	}

	return identifier
}

func (idstr IdentifierStr) GetName() string {
	if !IsIdentifierStrValid(idstr) {
		return ""
	}

	identifierSplitted := idstr.getIdentifierSplitted()
	return identifierSplitted[0]
}

func (idstr IdentifierStr) GetTag() string {
	if !IsIdentifierStrValid(idstr) {
		return ""
	}

	identifierSplitted := idstr.getIdentifierSplitted()
	return identifierSplitted[1]
}

func (idstr IdentifierStr) getIdentifierSplitted() []string {
	return strings.Split(string(idstr), identifierStrSeparetor)
}

func IsIdentifierStrValid(identifierStr IdentifierStr) bool {
	preparedIdentifierRegex := util.CreatePreparedRegex(identifierFormatRegex)
	preparedOnlyNameRegex := util.CreatePreparedRegex(nameFormatRegex)

	return preparedIdentifierRegex.MatchString(string(identifierStr)) || preparedOnlyNameRegex.MatchString(string(identifierStr))
}
