package util

import (
	"testing"

	"github.com/Luka-Spa/SwipeRight/packages/profile/model"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	// 32-byte secret
	var secret = "rsPzbMGz07CavdHwcAOy00aXWvXcnoG2"

	// String to be encrypted
	var str = "This string will be encrypted"
	var encoded, err = Encrypt([]byte(str), secret)
	assert.Nil(t, err)
	assert.NotEqual(t, str, encoded)
}

func TestDecrypt(t *testing.T) {
	// 32-byte secret
	var secret = "xMmz2AWsH4Vmzcidgt2a043394qgnUrI"

	// String to be decrypted
	var str = "This string will be decrypted"

	var encoded, _ = Encrypt([]byte(str), secret)

	var decoded, err = Decrypt(encoded, secret)

	assert.Nil(t, err)
	assert.NotEqual(t, str, encoded)
	assert.Equal(t, str, decoded)
}

func TestEncryptInvalidSecret(t *testing.T) {
	var secret = "invalidsecret"

	var str = "This string will be encrypted"

	var encoded, err = Encrypt([]byte(str), secret)

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
	EncryptProps(user, "xMmz2AWsH4Vmzcidgt2a043394qgnUrI")
	assert.NotEqual(t, school, user.School)
	assert.NotEqual(t, email, user.Email)
	assert.Equal(t, anthem, user.Anthem)
}

func TestDecryptProps(t *testing.T) {
	var secret = "xMmz2AWsH4Vmzcidgt2a043394qgnUrI"
	var email = "test@example.com"
	var anthem = "test anthem"
	var school = "test school"
	var user = &model.UserProfile{
		Email:  email,
		Anthem: anthem,
		School: school,
	}
	EncryptProps(&user, secret)
	DecryptProps(&user, secret)
	assert.Equal(t, school, user.School)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, anthem, user.Anthem)
}
