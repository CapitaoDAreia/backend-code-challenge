package configuration

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func GenerateSecretKey() {
	key := "my_temp_secret_key"

	base64Result := base64.StdEncoding.EncodeToString([]byte(key))

	fmt.Println("SecretKey to userToken generated: " + base64Result)

	os.Setenv("SecretKey", base64Result)
}

func GetSecretKey() []byte {
	base64Key := os.Getenv("SecretKey")

	decodedKey, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatal("Error on decode base64key...")
	}

	return decodedKey
}
