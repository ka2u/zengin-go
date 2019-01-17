package zengincode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ka2u/zengin-code-go/embed"
)

// New is making Japanese ZenginCode Information mapping data
func New() (map[string]*Bank, error) {
	path := os.Getenv("ZENGIN_SOURCE_ROOT")
	if path == "" {
		return nil, errors.New("should set ZENGIN_SOURCE environment variable")
	}
	dataDir := filepath.Join(path, "data")
	yaml := os.Getenv("ZENGIN_SOURCE_YAML")
	include := os.Getenv("ZENGIN_SOURCE_INCLUDE")

	banksFile, err := getBanksFile(include, yaml, dataDir)
	if err != nil {
		return nil, err
	}
	bank := map[string]*Bank{}

	err = json.Unmarshal(banksFile, &bank)
	if err != nil {
		return nil, err
	}
	branchDir := filepath.Join(dataDir, "branches")
	for code := range bank {
		branchFile, err := getBranchFile(include, yaml, code, branchDir)

		branch := map[string]*Branch{}
		err = json.Unmarshal(branchFile, &branch)
		if err != nil {
			return nil, err
		}
		bank[code].Branches = branch
	}

	return bank, nil
}

func getBanksFile(include string, yaml string, dataDir string) ([]byte, error) {
	var banksFile []byte
	var err error
	if include == "TRUE" {
		if yaml == "TRUE" {
			banksFile, err = embed.ReadFile("source-data/data/banks.yaml")
			if err != nil {
				return nil, err
			}
		} else {
			banksFile, err = embed.ReadFile("source-data/data/banks.json")
			if err != nil {
				return nil, err
			}
		}
	} else {
		if yaml == "TRUE" {
			yamlSource := filepath.Join(dataDir, "banks.yml")
			banksFile, err = ioutil.ReadFile(yamlSource)
			if err != nil {
				return nil, err
			}
		} else {
			jsonSource := filepath.Join(dataDir, "banks.json")
			banksFile, err = ioutil.ReadFile(jsonSource)
			if err != nil {
				return nil, err
			}
		}
	}
	return banksFile, nil
}

func getBranchFile(include string, yaml string, code string, branchDir string) ([]byte, error) {
	var branchFile []byte
	var err error
	if include == "TRUE" {
		if yaml == "TRUE" {
			branchFile, err = embed.ReadFile("source-data/data/branches/" + code + ".yaml")
			if err != nil {
				return nil, err
			}
		} else {
			branchFile, err = embed.ReadFile("source-data/data/branches/" + code + ".json")
			if err != nil {
				return nil, err
			}
		}

	} else {
		if yaml == "TRUE" {
			branchSource := filepath.Join(branchDir, code+".yaml")
			branchFile, err = ioutil.ReadFile(branchSource)
		} else {
			branchSource := filepath.Join(branchDir, code+".json")
			branchFile, err = ioutil.ReadFile(branchSource)
		}

	}
	return branchFile, nil
}
