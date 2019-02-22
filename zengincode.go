package zengincode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/derekparker/trie"

	"github.com/ka2u/zengin-code-go/embed"
)

type BankDB struct {
	Bank *trie.Trie
}

type BranchDB struct {
	Branch *trie.Trie
}

// New is making Japanese ZenginCode Information mapping data
func New() (*BankDB, error) {
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

	bankdb := trie.New()

	branchDir := filepath.Join(dataDir, "branches")
	for code := range bank {
		bk := bank[code]
		bk.Code = code

		branchFile, err := getBranchFile(include, yaml, code, branchDir)

		branch := map[string]*Branch{}
		err = json.Unmarshal(branchFile, &branch)
		if err != nil {
			return nil, err
		}
		branchdb := trie.New()
		for bcode := range branch {
			br := branch[bcode]
			branchdb.Add(bcode, br)
			branchdb.Add(br.Name, br)
			branchdb.Add(br.Roma, br)
		}

		brd := &BranchDB{
			Branch: branchdb,
		}
		bk.Branches = brd
		bankdb.Add(code, bk)
		bankdb.Add(bk.Name, bk)
		bankdb.Add(bk.Roma, bk)
	}

	bd := &BankDB{
		Bank: bankdb,
	}
	return bd, nil
}

func (b *BankDB) Find(key string) (*Bank, error) {
	node, ok := b.Bank.Find(key)
	if ok == false {
		return nil, errors.New("not found")
	}
	bank := node.Meta().(*Bank)
	return bank, nil
}
func (b *BankDB) PrefixSearch(pre string) []string {
	return b.Bank.PrefixSearch(pre)
}
func (b *BankDB) HasKeysWithPrefix(key string) bool {
	return b.Bank.HasKeysWithPrefix(key)
}
func (b *BankDB) FuzzySearch(pre string) []string {
	return b.Bank.FuzzySearch(pre)
}

func (b *BranchDB) Find(key string) (*Branch, error) {
	node, ok := b.Branch.Find(key)
	if ok == false {
		return nil, errors.New("not found")
	}
	branch := node.Meta().(*Branch)
	return branch, nil
}
func (b *BranchDB) PrefixSearch(pre string) []string {
	return b.Branch.PrefixSearch(pre)
}
func (b *BranchDB) HasKeysWithPrefix(key string) bool {
	return b.Branch.HasKeysWithPrefix(key)
}
func (b *BranchDB) FuzzySearch(pre string) []string {
	return b.Branch.FuzzySearch(pre)
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
