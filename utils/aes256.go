package utils

import (
	"bytes"
	"consumerskripsi/config"
	"consumerskripsi/models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

// Decrypt from base64 to decrypted string
func DecryptMerchantKey(cryptoText string) (models.MerchantKey, error) {
	var result models.MerchantKey
	var merchKey models.MerchantKey
	keyText := config.SALT_KEY
	key := []byte(keyText)
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return result, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return result, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	unMarshall := json.Unmarshal(ciphertext, &merchKey)
	fmt.Println(unMarshall)
	result = merchKey
	return result, nil
}

// EncryptMerchantKey  to base64 crypto using AES
func EncryptMerchantKey(req models.MerchantKey) (models.ResultEncryptMerchKey, error) {
	var result models.ResultEncryptMerchKey
	keyText := config.SALT_KEY
	key := []byte(keyText)
	var DataEncrypt interface{}
	DataEncrypt = req
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(DataEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(reqBodyBytes.Bytes()))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return result, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], reqBodyBytes.Bytes())

	// convert to base64
	result.Result = base64.URLEncoding.EncodeToString(ciphertext)
	return result, nil
}
