package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"reflect"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

type crypt struct {
	secret string
}

type ICryptor interface {
	Encrypt(text []byte) (string, error)
	Decrypt(text string) (string, error)
	EncryptProps(o interface{}) error
	DecryptProps(o interface{}) error
}

func NewCryptor(secret string) ICryptor {
	return &crypt{secret: secret}
}

func (*crypt) encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (*crypt) decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (c *crypt) Encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(c.secret))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(text))
	cfb.XORKeyStream(cipherText, text)
	return c.encode(cipherText), nil
}

func (c *crypt) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(c.secret))
	if err != nil {
		return "", err
	}
	cipherText := c.decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func (c *crypt) EncryptProps(o interface{}) error {
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
			c.EncryptProps(l)
			continue
		}
		_, ok := sourceTypeField.Tag.Lookup("encrypt")
		if !ok {
			continue
		}

		result, err := c.Encrypt([]byte(sourceField.String()))
		if err != nil {
			return err
		}
		sourceField.SetString(result)
	}
	return nil
}

func (c *crypt) DecryptProps(o interface{}) error {
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
		result, err := c.Decrypt(sourceField.String())
		if err != nil {
			return err
		}
		sourceField.SetString(result)
	}
	return nil
}
