package zengincode

// Bank is Japanes Bank information structure
type Bank struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Kana     string `json:"kana"`
	Hira     string `json:"hira"`
	Roma     string `json:"roma"`
	Branches *BranchDB
}
