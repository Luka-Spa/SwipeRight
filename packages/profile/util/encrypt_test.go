package util

import (
	"os"
	"testing"

	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/stretchr/testify/assert"
)

var cryptor *Cryptor

func TestMain(m *testing.M) {
	cryptor = NewCryptor("xMmz2AWsH4Vmzcidgt2a043394qgnUrI")
	code := m.Run()
	os.Exit(code)
}
func TestEncrypt(t *testing.T) {
	// String to be encrypted
	var str = "This string will be encrypted"
	var encoded, err = cryptor.Encrypt([]byte(str))
	assert.Nil(t, err)
	assert.NotEqual(t, str, encoded)
}

func TestDecrypt(t *testing.T) {
	// 32-byte secret
	//var secret = "xMmz2AWsH4Vmzcidgt2a043394qgnUrI"

	// String to be decrypted
	var str = "This string will be decrypted"

	var encoded, _ = cryptor.Encrypt([]byte(str))

	var decoded, err = cryptor.Decrypt(encoded)

	assert.Nil(t, err)
	assert.NotEqual(t, str, encoded)
	assert.Equal(t, str, decoded)
}

func TestEncryptInvalidSecret(t *testing.T) {
	var secret = "invalidsecret"
	var invalidCryptor = NewCryptor(secret)
	var str = "This string will be encrypted"

	var encoded, err = invalidCryptor.Encrypt([]byte(str))

	// Returns empty string
	assert.Equal(t, "", encoded)
	assert.NotEqual(t, nil, err)
}

func TestEncryptProps(t *testing.T) {
	var email = "test@example.com"
	var anthem = "test anthem"
	var school = "test school"
	var user = &model.UserProfile{
		Email:  email,
		Anthem: anthem,
		School: school,
	}
	err := cryptor.EncryptProps(user)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, email, user.Email)
	assert.NotEqual(t, school, user.School)
	assert.Equal(t, anthem, user.Anthem)
}

func TestDecryptProps(t *testing.T) {
	//var secret = "xMmz2AWsH4Vmzcidgt2a043394qgnUrI"
	var email = "test@example.com"
	var anthem = "test anthem"
	var school = "test school"
	var user = &model.UserProfile{
		Email:  email,
		Anthem: anthem,
		School: school,
	}
	cryptor.EncryptProps(&user)
	cryptor.DecryptProps(&user)
	assert.Equal(t, school, user.School)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, anthem, user.Anthem)
}
