package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform/helper/schema"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func loadJSONFromFile(filename string, target interface{}) error {
	if !fileExists(filename) {
		return errors.New("file does not exist")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	return nil
}

func saveJSONToFile(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getModulePath() (string, error) {
	modulePath := os.Getenv("INFRA_TERRAFORM_MODULE_PATH")
	if modulePath == "" {
		// default to current working directory
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return wd, nil
	}
	return modulePath, nil
}

func getProviderConfig(d *schema.ResourceData, provider string) (map[string]interface{}, error) {
	providerConfig := make(map[string]interface{})

	if v, ok := d.GetOk("provider_config"); ok {
		config := v.(map[string]interface{})
		if config[provider] != nil {
			providerConfig = config[provider].(map[string]interface{})
		}
	}

	return providerConfig, nil
}

func logError(err error) {
	log.Printf("ERROR: %s\n", err)
}

func logInfo(message string) {
	log.Printf("INFO: %s\n", message)
}

func logDebug(message string) {
	if os.Getenv("INFRA_TERRAFORM_DEBUG") != "" {
		log.Printf("DEBUG: %s\n", message)
	}
}