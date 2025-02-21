package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// Check and return the error.
func CheckAndReturnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Get the home directory
func GetHomeDirectory() string {
	homeDir, err := os.UserHomeDir()
	CheckAndReturnError(err)
	return homeDir
}

// Check if "~/.aws/credentials" file exists on the user system
func FileHomeExist() bool {
	filePath := GetHomeDirectory() + "/.aws/credentials"
	if _, err := os.Stat(filePath); err == nil {
		return false
	}

	return true
}

// Check if "~/.swap/saved-configuration" file exists on the user system
func ConfigFileExist() bool {
	filePath := GetHomeDirectory() + "/.swap/saved-configuration"
	if _, err := os.Stat(filePath); err == nil {
		return false
	}

	return true
}

// List all the aws profiles
func ReadFile() {
	
	filePath := GetHomeDirectory() + "/.aws/credentials"

	// Load de ini file
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

// Read all the profiles sections
func ReadArray() []string {
	
	filePath := GetHomeDirectory() + "/.aws/credentials"

	// Cargar el archivo INI
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	var profiles []string
	for _, section := range inidata.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		profiles = append(profiles, section.Name())
	}

	return profiles
}

// Check if the profile exists
func CheckArray(profile string) {

	array := ReadArray()
	count := 0
	for _, value := range array {
		if value == profile {
			count++
		}
	}
	if count < 1 {
		fmt.Println("\nswap: oh! It seems that you don't have this profile!\nswap: please check your credentials file.")
	}
}

//Change the default profile
func ChangeCredentials(profileName string) *ini.File {
	
	filePath := GetHomeDirectory() + "/.aws/credentials"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	defaultSection := inidata.Section("default")

	// Create a new bridge section
	bridgeSection, err := inidata.NewSection("bridge")
	CheckAndReturnError(err)

	// Change the data between the default and the bridge section
	for _, key := range defaultSection.Keys() {
		keyName := key.Name()
		value := defaultSection.Key(keyName).String()
		bridgeSection.NewKey(keyName, value)
	}

	targetSection := inidata.Section(profileName)

	// Change the data between the target and the default section
	for _, key := range targetSection.Keys() {
		keyName := key.Name()
		value := targetSection.Key(keyName).String()
		defaultSection.NewKey(keyName, value)
	}

	// Change the data between the bridge and the default section
	for _, key := range bridgeSection.Keys() {
		keyName := key.Name()
		value := bridgeSection.Key(keyName).String()
		targetSection.NewKey(keyName, value)
	}

	err = inidata.SaveTo(filePath)
	CheckAndReturnError(err)

	return inidata
}

//Change the default profile
func ChangeConfig(profileName string) *ini.File {
	
	filePath := GetHomeDirectory() + "/.aws/config"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	defaultSection := inidata.Section("default")

	// Create a new bridge section
	bridgeSection, err := inidata.NewSection("bridge")
	CheckAndReturnError(err)

	// Change the data between the default and the bridge section
	for _, key := range defaultSection.Keys() {
		keyName := key.Name()
		value := defaultSection.Key(keyName).String()
		bridgeSection.NewKey(keyName, value)
	}

	targetSection := inidata.Section(profileName)

	// Change the data between the target and the default section
	for _, key := range targetSection.Keys() {
		keyName := key.Name()
		value := targetSection.Key(keyName).String()
		defaultSection.NewKey(keyName, value)
	}

	// Change the data between the bridge and the default section
	for _, key := range bridgeSection.Keys() {
		keyName := key.Name()
		value := bridgeSection.Key(keyName).String()
		targetSection.NewKey(keyName, value)
	}

	err = inidata.SaveTo(filePath)
	CheckAndReturnError(err)

	return inidata
}

func DeleteBridge() *ini.File {
	
	filePath := GetHomeDirectory() + "/.aws/credentials"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	// Check if the profile exists
	_, err = inidata.GetSection("bridge")
	CheckAndReturnError(err)

	// Delete the bridge profile
	inidata.DeleteSection("bridge")

	err = inidata.SaveTo(filePath)
	if err != nil {
		fmt.Printf("Error guardando el archivo INI: %v", err)
	}

	return inidata
}

func DeleteBridgeConfig() *ini.File {
	
	filePath := GetHomeDirectory() + "/.aws/config"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	// Check if the profile exists
	_, err = inidata.GetSection("bridge")
	CheckAndReturnError(err)

	// Delete the bridge profile
	inidata.DeleteSection("bridge")

	err = inidata.SaveTo(filePath)
	if err != nil {
		fmt.Printf("Error guardando el archivo INI: %v", err)
	}

	return inidata
}

func SaveCredentials() *ini.File {

	filePath := GetHomeDirectory() + "/.aws/credentials"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	copyPath := GetHomeDirectory() + "/.swap/saved-aws-credentials"

	// Create the dir
	dirPath := filepath.Dir(copyPath)
	err = os.MkdirAll(dirPath, 0755) // Create the dir and assing permissions
	CheckAndReturnError(err)

	err = inidata.SaveTo(copyPath)
	CheckAndReturnError(err)

	//Return de ini file
	return inidata
}

func SaveConfig() *ini.File {

	filePath := GetHomeDirectory() + "/.aws/config"

	// Load the ini file
	inidata, err := ini.Load(filePath)
	CheckAndReturnError(err)

	copyPath := GetHomeDirectory() + "/.swap/saved-aws-config"

	// Create the dir
	dirPath := filepath.Dir(copyPath)
	err = os.MkdirAll(dirPath, 0755) // Create the dir and assing permissions
	CheckAndReturnError(err)

	err = inidata.SaveTo(copyPath)
	CheckAndReturnError(err)

	//Return de ini file
	return inidata
}

func RestoreCredentials() *ini.File {

	filePath := GetHomeDirectory() + "/.aws/credentials"

	copyPath := GetHomeDirectory() + "/.swap/saved-aws-credentials"

	copydata, err := ini.Load(copyPath)
	CheckAndReturnError(err)

	// Save the ini file in the configuration file
	err = copydata.SaveTo(filePath)
	CheckAndReturnError(err)

	// Return the ini file
	return copydata
}

func RestoreConfig() *ini.File {

	filePath := GetHomeDirectory() + "/.aws/config"

	copyPath := GetHomeDirectory() + "/.swap/saved-aws-config"

	copydata, err := ini.Load(copyPath)
	CheckAndReturnError(err)

	// Save the ini file in the configuration file
	err = copydata.SaveTo(filePath)
	CheckAndReturnError(err)

	// Return the ini file
	return copydata
}