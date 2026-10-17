package chat

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// The legacy-aud golden vectors are no longer what MintEmbedToken emits, but the
// HS256 primitive must still reproduce them byte-for-byte.
func TestLegacyAudienceVectorsReplay(t *testing.T) {
	for _, name := range []string{"jwt_vector.json", "jwt_vector_credentials.json"} {
		raw, err := os.ReadFile("../internal/contract/testdata/chat/" + name)
		if err != nil {
			t.Fatal(err)
		}
		var v struct {
			Secret       string `json:"secret"`
			HeaderJSON   string `json:"header_json"`
			ClaimsJSON   string `json:"claims_json"`
			SigningInput string `json:"signing_input"`
			Token        string `json:"token"`
		}
		if err := json.Unmarshal(raw, &v); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(v.ClaimsJSON, `"aud":"rootcause:chat:acme"`) {
			t.Fatalf("%s: not the legacy-aud vector", name)
		}
		if got := b64([]byte(v.HeaderJSON)) + "." + b64([]byte(v.ClaimsJSON)); got != v.SigningInput {
			t.Fatalf("%s: signing input differs", name)
		}
		if got := v.SigningInput + "." + signHS256(v.Secret, v.SigningInput); got != v.Token {
			t.Fatalf("%s: token\n got: %s\nwant: %s", name, got, v.Token)
		}
	}
}
