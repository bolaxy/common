package common

import (
	"os"
	"path"
	"path/filepath"
)

func Home(base string) string {
	usrHome, _ := os.UserHomeDir()
	if usrHome == "" {
		return ""
	}
	return path.Join(usrHome, base)
}

func Env(key string) string {
	envPath, _ := os.LookupEnv(key)
	return envPath
}

func WorkDir() string {
	pwd, err := os.Getwd()
	if err == nil {
		return pwd
	}

	return ""
}

func ExeDir() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	return ""
}
