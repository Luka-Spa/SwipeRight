package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"reflect"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var _secret string

type Cryptor struct {
}

func NewCryptor(secret string) *Cryptor {
	_secret = secret
	return &Cryptor{}
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (*Cryptor) Encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(_secret))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(text))
	cfb.XORKeyStream(cipherText, text)
	return encode(cipherText), nil
}

func (*Cryptor) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(_secret))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func (cryptor *Cryptor) EncryptProps(o interface{}) error {
	ptrValue := reflect.ValueOf(o)
	v := ptrValue.Elem()
	vType := v.Type()
	if vType.Kind() != reflect.Struct {
		return errors.New("type of object is not a struct")
	}
	for i := 0; i < vType.NumField(); i++ {
		sourceTypeField := vType.Field(i)
		sourceField := v.Field(i)
		if sourceTypeField.Type.Kind() == reflect.Struct {
			l := sourceField.Addr().Interface()
			cryptor.EncryptProps(l)
			continue
		}
		_, ok := sourceTypeField.Tag.Lookup("encrypt")
		if !ok {
			continue
		}

		result, err := cryptor.Encrypt([]byte(sourceField.String()))
		if err != nil {
			return err
		}
		sourceField.SetString(result)
	}
	return nil
}

func (cryptor *Cryptor) DecryptProps(o interface{}) error {
	ptrValue := reflect.ValueOf(o)
	v := ptrValue.Elem()
	vType := v.Type()
	if vType.Kind() != reflect.Struct {
		return errors.New("type of object is not a struct")
	}
	for i := 0; i < vType.Elem().NumField(); i++ {
		sourceTypeField := vType.Elem().Field(i)
		sourceField := v.Elem().Field(i)
		_, ok := sourceTypeField.Tag.Lookup("encrypt")
		if !ok {
			continue
		}
		result, err := cryptor.Decrypt(sourceField.String())
		if err != nil {
			return err
		}
		sourceField.SetString(result)
	}
	return nil
}
