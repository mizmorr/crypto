package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
)

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func signMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func verifySignature(publicKey *ecdsa.PublicKey, message, signature []byte) bool {
	hash := sha256.Sum256(message)
	return ecdsa.VerifyASN1(publicKey, hash[:], signature)
}

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		fmt.Println("Failed to generate key pair:", err)
		os.Exit(1)
	}

	filePath := "file.txt"
	message, err := readFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		os.Exit(1)
	}

	signature, err := signMessage(privateKey, message)
	if err != nil {
		fmt.Println("Failed to sign message:", err)
		os.Exit(1)
	}

	signatureFilePath := "signature.sig"
	err = writeFile(signatureFilePath, signature)
	if err != nil {
		fmt.Println("Failed to write signature file:", err)
		os.Exit(1)
	}

	signatureToVerify, err := readFile(signatureFilePath)
	if err != nil {
		fmt.Println("Failed to read signature file:", err)
		os.Exit(1)
	}

	isValidSignature := verifySignature(publicKey, message, signatureToVerify)
	if isValidSignature {
		fmt.Println("Подпись верна.")
	} else {
		fmt.Println("Подпись неверна.")
	}
}
