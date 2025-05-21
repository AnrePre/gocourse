package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	//Hashing
	hash := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println(password)
	fmt.Println(hash)
	fmt.Printf("SHA-256 Hash hex val: %x \n", hash)
	fmt.Printf("SHA-512 Hash hex val: %x \n", hash512)

	//Salting
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", salt)
		return
	}

	//Hashing the password with salt
	signupHash := hashPassword(password, salt)

	//Store the salt and password in database (print to console for now)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", signupHash)

	//verify password
	//retrieve the saltStr and decode
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt", err)
	}
	loginHash := hashPassword(password, decodedSalt)

	//Compare the stored signupHash with the loginHash
	if signupHash == loginHash {
		fmt.Println("Password is correct. You are logged in!")
	} else {
		fmt.Println("Password is incorrect. Please check user credentials.")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])

}
