package zengincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/derekparker/trie"
	"github.com/rakyll/statik/fs"
)

type BankDB struct {
	Bank *trie.Trie
}

type BranchDB struct {
	Branch *trie.Trie
}

// New is making Japanese ZenginCode Information mapping data
func New() (*BankDB, error) {
	var path, dataDir string
	path = os.Getenv("ZENGIN_SOURCE_ROOT")
	if path == "" {
		return nil, errors.New("should set ZENGIN_SOURCE_ROOT environment variable")
	}
	dataDir = filepath.Join(path, "data")

	yaml := os.Getenv("ZENGIN_SOURCE_YAML")

	banksFile, err := getBanksFile(yaml, dataDir)
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

		branchFile, err := getBranchFile(yaml, code, branchDir)

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

func NewWithEmbed() (*BankDB, error) {
	yaml := os.Getenv("ZENGIN_SOURCE_YAML")

	fmt.Println("banks file from embed")
	banksFile, err := getBanksFileFromEmbed(yaml)
	if err != nil {
		return nil, err
	}

	bank := map[string]*Bank{}
	err = json.Unmarshal(banksFile, &bank)
	if err != nil {
		return nil, err
	}

	bankdb := trie.New()
	fmt.Println("bank loop")
	for code := range bank {
		bk := bank[code]
		bk.Code = code

		branchFile, err := getBranchFileFromEmbed(yaml, code, "/branches")

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

func getBanksFile(yaml string, dataDir string) ([]byte, error) {
	var banksFile []byte
	var err error

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

	return banksFile, nil
}

func getBranchFile(yaml string, code string, branchDir string) ([]byte, error) {
	var branchFile []byte
	var err error

	if yaml == "TRUE" {
		branchSource := filepath.Join(branchDir, code+".yaml")
		branchFile, err = ioutil.ReadFile(branchSource)
		if err != nil {
			return nil, err
		}
	} else {
		branchSource := filepath.Join(branchDir, code+".json")
		branchFile, err = ioutil.ReadFile(branchSource)
		if err != nil {
			return nil, err
		}
	}

	return branchFile, nil
}

func getBanksFileFromEmbed(yaml string) ([]byte, error) {
	var banksFile []byte
	var err error
	var f http.File

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	if yaml == "TRUE" {
		f, err = statikFS.Open("/banks.yml")
		if err != nil {
			return nil, err
		}
	} else {
		f, err = statikFS.Open("/banks.json")
		if err != nil {
			return nil, err
		}
	}
	banksFile, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return banksFile, nil
}

func getBranchFileFromEmbed(yaml string, code string, branchDir string) ([]byte, error) {
	var branchFile []byte
	var err error
	var f http.File

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	if yaml == "TRUE" {
		branchSource := filepath.Join(branchDir, code+".yaml")
		f, err = statikFS.Open(branchSource)
	} else {
		branchSource := filepath.Join(branchDir, code+".json")
		f, err = statikFS.Open(branchSource)
	}
	branchFile, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return branchFile, nil
}
