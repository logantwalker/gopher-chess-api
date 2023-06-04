package model

type UciCommand struct {
	UciString string `json:"uci_string"`
	Moves []string `json:"moves"`
}