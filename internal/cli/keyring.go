package cli

import (
	"github.com/zalando/go-keyring"
)

const MAIN_KEYRING_SERVICE = "awasm-cli"

func SetValueInKeyring(key, value string) error {
	return keyring.Set(MAIN_KEYRING_SERVICE, key, value)
}

func GetValueInKeyring(key string) (string, error) {
	return keyring.Get(MAIN_KEYRING_SERVICE, key)
}

func DeleteValueInKeyring(key string) error {
	return keyring.Delete(MAIN_KEYRING_SERVICE, key)
}