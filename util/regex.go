package util

import (
	"regexp"
)

func CreatePreparedRegex(regex string) *regexp.Regexp {
	preparedRegex, errOnCompile := regexp.Compile(regex)

	if errOnCompile != nil {
		// Falha na aplicação, para tudo!
		return nil
	}
	return preparedRegex
}
