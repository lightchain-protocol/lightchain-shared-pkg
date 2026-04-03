package chain

import (
	"reflect"
	"testing"
)

func TestJobRegistryInterfaceIncludesNonce(t *testing.T) {
	t.Parallel()

	if _, ok := reflect.TypeOf((*JobRegistry)(nil)).Elem().MethodByName("Nonce"); !ok {
		t.Fatal("JobRegistry interface must expose Nonce")
	}
}
