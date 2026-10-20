// Package credentials holds the one validator for the `credentials` map shared by
// the chat token claim (hub decision 20) and the analysis trigger (hub decision
// 29). Both the root package and chat/ need it and imports only run one way, so it
// lives here. Callers map a false result onto their own customer-facing code.
package credentials

import (
	"encoding/json"
	"regexp"
	"strings"
)

// The host re-checks every rule; refusing here surfaces the mistake where it was made.
const (
	MaxEntries = 8
	MaxBytes   = 8 << 10
)

var key = regexp.MustCompile(`^[A-Z][A-Z0-9_]{0,63}$`)

// Valid reports whether credentials obey the claim rules. Empty is valid (it is
// omitted on the wire). A NUL value is refused: an env value cannot carry one.
func Valid(credentials map[string]string) bool {
	if len(credentials) == 0 {
		return true
	}
	if len(credentials) > MaxEntries {
		return false
	}
	for k, v := range credentials {
		if !key.MatchString(k) || strings.HasPrefix(k, "RC_") || strings.ContainsRune(v, 0) {
			return false
		}
	}
	encoded, err := json.Marshal(credentials)
	return err == nil && len(encoded) <= MaxBytes
}
