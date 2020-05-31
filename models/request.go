package models

type UTXORequestParameter struct {
	From       int      `json:"from,omitempty"`
	To         int      `json:"to,omitempty"`
	FromHeight int      `json:"fromHeight,omitempty"`
	ToHeight   int      `json:"toHeight,omitempty"`
	Addresses  []string `json:"address"`
}
