package salts_repo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type project struct {
	ProjectId     string `json:"project_id"`
	ShorthandName string `json:"shorthand_name"`
	EncryptedSalt string `json:"encrypted_salt"`
}

type saltsRepo struct {
	StorageKeySalt string    `json:"storage_key_salt"`
	Salts          []project `json:"salts"`
}

func ReadSaltRepo(fileName string) saltsRepo {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	var salts_repo saltsRepo

	err = json.Unmarshal(bytes, &salts_repo)
	if err != nil {
		fmt.Println(err)
	}
	return (salts_repo)
}
