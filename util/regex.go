package util

import (
	"regexp"
	"strings"

	"didockerf/out"
)

func CreatePreparedRegex(regex string) *regexp.Regexp {
	preparedRegex, errOnCompile := regexp.Compile(regex)

	if errOnCompile != nil {
		out.FatalError(errOnCompile)
	}
	return preparedRegex
}

func CreatePreparedRegexWithReplace(regex string, old string, new string) *regexp.Regexp {
	newRegex := strings.ReplaceAll(regex, old, new)

	return CreatePreparedRegex(newRegex)
}
