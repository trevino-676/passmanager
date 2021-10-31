package encrypt_test

import (
	"testing"

	"github.com/trevino-676/passmanager/pkg/encrypt"
)

func TestEncryptRegistry(t *testing.T) {
	registry := "facebook|trevino|!twitter\n"
	if hash, _ := encrypt.EncryptRegistry(registry, "secret_key_secre"); hash == "" {
		t.Fatalf("Can't encrypt the registry. want: %v", hash)
	}
}

func TestDecryptRegistry(t *testing.T) {
	registry := "facebook|trevino|!twitter\n"
	hash, _ := encrypt.EncryptRegistry(registry, "secret_key_secre")
	if decrypt, _ := encrypt.DecryptRegistry(hash, "secret_key_secre"); decrypt != registry {
		t.Fatalf("Can't decrypt the registry. recibe: %v", decrypt)
	}
}
