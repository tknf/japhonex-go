package japhonex

import (
	"regexp"
)

const (
	REQUIRE_JA_TEL_REGEX = `(^0([0-9]-[0-9]{4}|[0-9]{2}-[0-9]{3}|[0-9]{3}-[0-9]{2}|[0-9]{4}-[0-9]{1})-[0-9]{4}$)|(^[0-9]{3}-[0-9]{4}-[0-9]{4}$)|(^0120-([0-9]-[0-9]{5}|[0-9]{2}-[0-9]{4}|[0-9]{3}-[0-9]{3}|[0-9]{4}-[0-9]{2}|[0-9]{5}-[0-9]|[0-9]{6})$)`
	OPTIONAL_HYPHEN_JA_TEL_REGEX  = `(^0([0-9]-?[0-9]{4}|[0-9]{2}-?[0-9]{3}|[0-9]{3}-?[0-9]{2}|[0-9]{4}-?[0-9]{1})-?[0-9]{4}$)|(^[0-9]{3}-?[0-9]{4}-?[0-9]{4}$)|(^0120-?([0-9]-?[0-9]{5}|[0-9]{2}-?[0-9]{4}|[0-9]{3}-?[0-9]{3}|[0-9]{4}-?[0-9]{2}|[0-9]{5}-?[0-9]|[0-9]{6})$)`
	NO_HYPHEN_JA_TEL_REGEX = `^[0-9]{10,11}$`
)

func RequireHyphen(s string) (bool, error) {
	return regexp.Match(REQUIRE_JA_TEL_REGEX, []byte(s))
}

func OptionalHyphen() (bool, error) {
	return regexp.Match(OPTIONAL_HYPHEN_JA_TEL_REGEX, []byte("fjiod"))
}

func NoHyphen(s string) (bool, error) {
  return	regexp.Match(NO_HYPHEN_JA_TEL_REGEX, []byte(s))
}