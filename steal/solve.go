package main

import (
	"fmt"
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"io"
)

func GunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer gr.Close()

	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)

	return nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	key, _ := hex.DecodeString("8a40cdd5c4608b251b2c5926270540dc")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("Cipher text too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func decrypt(dd string) {
    data, _ := hex.DecodeString(dd)
	dcData := bytes.Buffer{}
	if err := GunzipWrite(&dcData, data); err != nil {
		fmt.Printf("err gunzip")
	}

	// Decrypt the data
	decryptData, err := Decrypt(dcData.Bytes())
	if err != nil {
		fmt.Printf("err decrypt")
	}


	fmt.Printf("%s\n", decryptData)
}

func main() {
    decrypt("1f8b08000000000002ff9a252e7af98452f8d9a3b2fb7e363d2af7e50808d5f2cbaeee7f04080000ffff01e298b519000000")
	decrypt("1f8b08000000000002ff007d0082ff2e8c8da9989c7b986ca497950bfbe8fc213cd3e5f135206a17346ac8e74e519f27525be6a166c0ba8f97a3c51ad45a2253ed6fdffd9eb0f2e88ba439d9ce61ee6d3c1ae717f05641029d576fefba12ed91bde10b7d31217c5842a4d5d54d5d4dbc8819d9c82a8fd1bd2715ea51509cff4c0fab2e32f6f2700a2dccf5e7010000ffffdc404e107d000000")
	decrypt("1f8b08000000000002ff007d0082ffbb142bfd17b60cbf5365894c6fc3089cdd1025388213a96fe20dccc435bb312e07d58441c806f9ef34395fe760212c26dc6367e8cd220d8a26bf2d34bab4e2adcbaca3010196398c4a79365e1c27340f4f984d84b5ffde137e16df5a8ad9b8f028ee1476c307e71bab37fdb0f98be04c2bbd5443f10e1f91bff1681182010000fffff7f3c9337d000000")
	decrypt("1f8b08000000000002ff00af0050ff044f1a467dab6b9371a8cbd55cce223a2573f4144345f8a1d43d7db33258a794d5da83bfabbe5ac3ab36f9c2b9b034091202a6dc5dd67f0d42cc641c379f89de12a56214b519d468347eea14a7ef621e4f3ff9c9b86185b663b4490379da470226b4ad98b3500ef75d241afae01cdd0e1e0f22760b44f8bb2fb1f9ff308b2565d5d2fb59988723ce11e8bc68186ff4b768244f79dc5503ceafb9b60841f9ce7869c048644a318e9196358851879d4d010000ffff3cd91697af000000")
}
