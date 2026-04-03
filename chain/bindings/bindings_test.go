package bindings

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobRegistryABIIncludesUpgradeableAdminSurface(t *testing.T) {
	t.Parallel()

	var entries []struct {
		Type string `json:"type"`
		Name string `json:"name"`
	}
	require.NoError(t, json.Unmarshal([]byte(JobRegistryABI), &entries))

	names := make(map[string]struct{}, len(entries))
	for _, entry := range entries {
		if entry.Type != "function" {
			continue
		}
		names[entry.Name] = struct{}{}
	}

	for _, name := range []string{
		"owner",
		"UPGRADE_INTERFACE_VERSION",
		"proxiableUUID",
		"upgradeToAndCall",
		"initialize",
		"withdraw",
	} {
		_, ok := names[name]
		require.Truef(t, ok, "JobRegistryABI missing %s", name)
	}
}
