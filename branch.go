package zengincode

// Branch is Japanese Bank Branch inofrmation structure
type Branch struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}
