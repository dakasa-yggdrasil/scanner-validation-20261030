// Synthetic CodeQL fixture; no real credentials are present.
package fixture

import (
    "crypto/sha256"
    "encoding/json"
    "strings"
)

type Reference struct { PasswordRef string }
type Inputs struct { Password string }

func extractPublicName(ref string) string {
    ref = strings.TrimSpace(ref)
    if !strings.HasPrefix(ref, "reference://") { return "" }
    name, _, _ := strings.Cut(strings.TrimPrefix(ref, "reference://"), "/")
    return name
}

func renderReference(external *Reference) map[string]any {
    if external == nil { return nil }
    if name := extractPublicName(external.PasswordRef); name != "" {
        return map[string]any{"referenceName": name}
    }
    return nil
}

func baselinePassword(v Inputs) [32]byte { return sha256.Sum256([]byte(v.Password)) }
func publicName(v *Reference) [32]byte {
    public := renderReference(v)
    raw, _ := json.Marshal(public)
    return sha256.Sum256(raw)
}
