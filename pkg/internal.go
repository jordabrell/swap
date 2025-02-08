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
	homeDir := GetHomeDirectory()

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	fmt.Printf("PROFILE:\n-----\n")
	for _, section := range inidata.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		fmt.Printf("%s", section.Name())
		fmt.Println()
	}
}

//Ara farem un altre funció que es digui ReadArray per a que llegeixi la array perfil per perfil
func ReadArray() []string{
	homeDir := GetHomeDirectory()

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)
	
	var profiles[] string
	for _, section := range inidata.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		profiles = append(profiles, section.Name())
	}

	return profiles
}

func CheckArray(profile string) {
	
	array := ReadArray()
	count := 0
	for _, value := range array {
		if value == profile {
			count ++
		}
	}
	if count < 1{
		fmt.Println("\nOh! It seems that you don't have this profile!\nPlease check your credentials file.")
	}
}

func ChangeProfile(profileName string) *ini.File {
	// Obtener el directorio de inicio del usuario
	homeDir := GetHomeDirectory()

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	//Declarem la secció default
	defaultSection := inidata.Section("prova")
	
	//Creem una nova secció pont
    bridgeSection, err := inidata.NewSection("bridge")
    CheckAndReturnError(err)

	//Passem les dades de default al pont
	for _, key := range defaultSection.Keys(){
		keyName := key.Name()
		value := defaultSection.Key(keyName).String()
		bridgeSection.NewKey(keyName, value)
	}

	targetSection := inidata.Section(profileName)

	//Passem les dades de la target a la default
	for _, key := range targetSection.Keys(){
		keyName := key.Name()
		value := targetSection.Key(keyName).String()
		defaultSection.NewKey(keyName, value)
	}

	//Passem les dades del pont a la target
	for _, key := range bridgeSection.Keys(){
		keyName := key.Name()
		value := bridgeSection.Key(keyName).String()
		targetSection.NewKey(keyName, value)
	}

	err = inidata.SaveTo(filePath)
    if err != nil {
        fmt.Printf("Error guardando el archivo INI: %v", err)
    }

	return inidata
}

func DeleteBridge() *ini.File{
	// Obtener el directorio de inicio del usuario
	homeDir := GetHomeDirectory()

	// Construir la ruta completa al archivo ~/.aws/credentials
	filePath := filepath.Join(homeDir, ".aws", "credentials")

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	//Comprobamos si el archivo existe
	_, err = inidata.GetSection("bridge")
	CheckAndReturnError(err)

	inidata.DeleteSection("bridge")

	err = inidata.SaveTo(filePath)
    if err != nil {
        fmt.Printf("Error guardando el archivo INI: %v", err)
    }

	return inidata
}