package models_test

import (
	"testing"

	"github.com/trevino-676/passmanager/pkg/models"
)

func TestCreatePassword(t *testing.T) {
	password := models.NewPassword("twitter", "trevino", "!facebook")
	expectedString := "twitter|trevino|!facebook"
	if password.ToString() != expectedString {
		t.Fatalf("The password struct doesn't create correctly. Obtain: %v, Expected: %v",
			password.ToString(), expectedString)
	}
}

func TestValidateString(t *testing.T) {
	password := models.NewPassword("Twitter", "trevino", "")
	if password.Validate() {
		t.Fatal("The password is correct")
	}
}

/* func TestEncryptPasswordStruct(t *testing.T) {
	secret_key := "secret_123456789"
	expectedEncryptString, _ := encrypt.EncryptRegistry("twitter|trevino|!facebook", secret_key)
	password := models.NewPassword("twitter", "trevino", "!facebook")
	if encryptPassword, _ := password.Encrypt(secret_key); encryptPassword != expectedEncryptString {
		t.Fatalf("The encrypt password doesn't match, Obtain: %v, Expected: %v",
			encryptPassword, expectedEncryptString)
	}
}*/
