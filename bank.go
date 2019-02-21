package zengincode

// Bank is Japanes Bank information structure
type Bank struct {
	Name     string `json:"name"`
	Kana     string `json:"kana"`
	Hira     string `json:"hira"`
	Roma     string `json:"roma"`
	Branches *BranchDB
}
