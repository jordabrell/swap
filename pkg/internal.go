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
	CheckAndReturnError(err)

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	fmt.Printf("PROFILE:\n-----\n")
	var count int32 = 1
	for _, section := range inidata.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		fmt.Printf("%d) %s", count, section.Name())
		count ++
		fmt.Println()
	}
}
