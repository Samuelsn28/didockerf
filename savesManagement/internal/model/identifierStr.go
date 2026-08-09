package model

import (
	"didockerf/util"
)

const (
	identifierStrSeparetor = ":"

	nameFormatRegex       = `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*`
	tagFormatRegex        = `[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*`
	identifierFormatRegex = `^` + nameFormatRegex + identifierStrSeparetor + tagFormatRegex + `$`
)

func IsIdentifierStrValid(identifierStr string) bool {
	preparedIdentifierRegex := util.CreatePreparedRegex(identifierFormatRegex)
	preparedOnlyNameRegex := util.CreatePreparedRegex(nameFormatRegex)

	return preparedIdentifierRegex.MatchString(identifierStr) || preparedOnlyNameRegex.MatchString(identifierStr)
}
