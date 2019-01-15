package zengincode

type Bank struct {
	Name     string `json:"name"`
	Kana     string `json:"kana"`
	Hira     string `json:"hira"`
	Roma     string `json:"roma"`
	Branches map[string]*Branch
}
