package models

type UTXORequestParameter struct {
	From       int      `json:"from,omitempty"`
	To         int      `json:"to,omitempty"`
	FromHeight int      `json:"fromHeight,omitempty"`
	ToHeight   int      `json:"toHeight,omitempty"`
	Addresses  []string `json:"address"`
}

type GetBlockParameter struct {
	Hash   string `json:"hash,omitempty"`
	Height uint32 `json:"height,omitempty"`
}

type SendTransactionParameter struct {
	Transaction   []byte `json:"transaction"`
	AllowHighFees bool   `json:"allow_high_fees,omitempty"`
	BypassLimits  bool   `json:"bypass_limits,omitempty"`
}

type GetTransactionParameter struct {
	Id string `json:"id"`
}

type ApplyStateTransactionParameter struct {
	StateTransition []byte `json:"state_transition"`
}

type GetIdentityParameter struct {
	Id string `json:"id"`
}
