package file_test

import (
	"fmt"
	"testing"

	passfile "github.com/trevino-676/passmanager/pkg/file"
)

const (
	plataform = "facebook"
	user      = "trevino"
	password  = "!twitter"
)

func TestStorePassword(t *testing.T) {
	expected := true
	if isStored, _ := passfile.StorePassword(plataform, user, password); !isStored {
		t.Fatalf("There no stored the password in the file, expected: %v, obtain: %v",
			expected, isStored)
	}
}

func TestRetrivePassword(t *testing.T) {
	expected := fmt.Sprintf("%v|%v", user, password)
	if retrive, _ := passfile.RetrivePassword(plataform); retrive != expected {
		t.Fatalf("The password of the %v plataform doesn't exist: expected %v, obtain: %v",
			plataform, expected, retrive)
	}
}
