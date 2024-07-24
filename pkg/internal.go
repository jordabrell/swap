package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func CheckAndReturnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1) 
	}
}

func GetHomeDirectory() string {
	homeDir, err := os.UserHomeDir()
	CheckAndReturnError(err)
	return homeDir
}
// Check if the file "~/.aws/credentials" exist on the user system
func FileHomeExist() bool {
	filePath := GetHomeDirectory() + "/.aws/credentials" 
	if _, err := os.Stat(filePath); err == nil {
		return false
	}

	return true
}

func ReadFile() {
	// Obtener el directorio de inicio del usuario
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Fail to get home directory: %v ", err)
		os.Exit(1)
	}

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Obtener la sección "Default" del archivo INI
	section := inidata.Section("default")

	// Imprimir las claves solicitadas
	fmt.Println(section.Key("aws_access_key_id").String())
	fmt.Println(section.Key("aws_secret_access_key").String())
	fmt.Println(section.Key("aws_session_token").String())
}
