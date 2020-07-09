package cccrypto

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"log"
	"strconv"
	"strings"
)

type Des struct {
	encryptedBytes []byte
	plaintextBytes []byte
	username       int
	IV             []byte
	desCipher      cipher.Block
}

func New(newUsername int) Des {
	var str strings.Builder
	str.Write([]byte("Jsg#kkdf*777"))
	str.Write([]byte(strconv.Itoa(newUsername)))
	key := str.String()

	hasher := md5.New()
	hasher.Write([]byte(key))
	md5 := hasher.Sum(nil)

	finalKey := append(md5, md5[:8]...)

	desBlock, err := des.NewTripleDESCipher(finalKey)
	if err != nil {
		log.Fatal(err)
	}

	return Des{
		encryptedBytes: nil,
		plaintextBytes: nil,
		username:       newUsername,
		IV:             []byte{93, 3, 22, 99, 4, 82, 162, 34},
		desCipher:      desBlock,
	}
}

func (d *Des) Decrypt(newEncryptedBytes []byte) string {
	d.encryptedBytes = newEncryptedBytes
	newPlaintextBytes := make([]byte, len(newEncryptedBytes))
	d.plaintextBytes = newPlaintextBytes

	desDecrypter := cipher.NewCBCDecrypter(d.desCipher, d.IV)
	desDecrypter.CryptBlocks(d.plaintextBytes, d.encryptedBytes)

	return string(pkcs5UnPadding(d.plaintextBytes))
}

func (d *Des) Encrypt(newPlaintextBytes []byte) string {
	bs := d.desCipher.BlockSize()
	d.plaintextBytes = pkcs5Padding(newPlaintextBytes, bs)

	newEncryptedBytes := make([]byte, len(d.plaintextBytes))
	d.encryptedBytes = newEncryptedBytes

	desEncrypter := cipher.NewCBCEncrypter(d.desCipher, d.IV)
	desEncrypter.CryptBlocks(d.encryptedBytes, d.plaintextBytes)

	return base64.StdEncoding.EncodeToString(d.encryptedBytes)
}

func pkcs5Padding(plaintextBytes []byte, blockSize int) []byte {
	padding := blockSize - len(plaintextBytes)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintextBytes = append(plaintextBytes, padtext...)

	if len(plaintextBytes)%blockSize != 0 {
		log.Fatal("Need a multiple of the blocksize")
	}
	return plaintextBytes
}

func pkcs5UnPadding(paddedPlaintext []byte) []byte {
	length := len(paddedPlaintext)
	unpadding := int(paddedPlaintext[length-1])
	return paddedPlaintext[:(length - unpadding)]
}
