package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func savePrivateKey(privateKey *rsa.PrivateKey, filename string) error {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	return os.WriteFile(filename, privateKeyPEM, 0600)
}

func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func encryptMessage(publicKey *rsa.PublicKey, message string) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func decryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) (string, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		fmt.Println("Failed to generate key pair:", err)
		return
	}

	err = savePrivateKey(privateKey, "private_key.pem")
	if err != nil {
		fmt.Println("Failed to save private key:", err)
		return
	}

	loadedPrivateKey, err := loadPrivateKey("private_key.pem")
	if err != nil {
		fmt.Println("Failed to load private key:", err)
		return
	}

	var message string
	fmt.Print("Введите сообщение для шифрования: ")
	fmt.Scanln(&message)

	encryptedMessage, err := encryptMessage(publicKey, message)
	if err != nil {
		fmt.Println("Failed to encrypt message:", err)
		return
	}
	fmt.Println("Зашифрованное сообщение:", encryptedMessage)

	decryptedMessage, err := decryptMessage(loadedPrivateKey, encryptedMessage)
	if err != nil {
		fmt.Println("Failed to decrypt message:", err)
		return
	}
	fmt.Println("Расшифрованное сообщение:", decryptedMessage)
}
