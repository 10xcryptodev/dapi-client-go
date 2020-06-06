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

type GetDataContractParameter struct {
	Id string `json:"id"`
}

type GetDocumentsParameter struct {
	DataContractId string `json:"data_contract_id"`
	DocumentType   string `json:"document_type"`
	Where          []byte `json:"where,omitempty"`
	OrderBy        []byte `json:"order_by,omitempty"`
	Limit          uint32 `json:"limit,omitempty"`
	StartAt        uint32 `json:"start_at,omitempty"`
	StartAfter     uint32 `json:"start_after,omitempty"`
}

type BloomFilterRequest struct {
	Data     []byte `json:"v_data"`
	HashFunc uint32 `json:"n_hash_funcs"`
	Tweak    uint32 `json:"n_tweak"`
	Flags    uint32 `json:"n_flags"`
}

type SubscribeToTransactionsWithProofsRequest struct {
	BloomFilter           BloomFilterRequest `json:"bloom_filter"`
	Count                 *int               `json:"count,omitempty"`
	FromBlockHash         *[]byte            `json:"from_block_hash,omitempty"`
	FromBlockHeight       *int               `json:"from_block_height,omitempty"`
	SendTransactionHashes *bool              `json:"send_transaction_hashes,omitempty"`
}
