package iso4217

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Removes Diacritics from submitted string
func convertDia(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)
	return s
}

// Check map for Alpha Match and return ISO
// example, "euro" will return "EUR"
func AlphaMatch(s string) *string {
	var (
		returnStr string
	)
	s = strings.ToLower(strings.TrimSpace(s))

	for key, val := range isocodes {
		if s == strings.ToLower(key) {
			returnStr = fmt.Sprintf("%v", key)
			return &returnStr
		}
		if convertDia(s) == strings.ToLower(strings.TrimSpace(val.currencyName)) {
			returnStr = fmt.Sprintf("%v", key)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(val.numCode)) {
			returnStr = fmt.Sprintf("%v", key)
			return &returnStr
		}
	}
	return nil
}

// Check map for Alpha3 Match and return ISO
// example, "gb" will return GBR Struct
func StructMatch(s string) *ISOEntry {
	var (
		entry ISOEntry
	)
	s = strings.ToLower(strings.TrimSpace(s))
	for key, val := range isocodes {
		if s == strings.ToLower(key) {
			entry = ISOEntry(val)
			return &entry
		}
		if convertDia(s) == strings.ToLower(strings.TrimSpace(val.currencyName)) {
			entry = ISOEntry(val)
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(val.numCode)) {
			entry = ISOEntry(val)
			return &entry
		}
	}
	return nil
}
