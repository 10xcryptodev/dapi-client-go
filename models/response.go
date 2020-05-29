package models

type AddressSummaryResponse struct {
	AddrStr                 []string
	Balance                 float64
	BalanceSat              int64
	TotalReceived           float64
	TotalReceivedSat        int64
	TotalSent               float64
	TotalSentSat            int64
	UnconfirmedBalance      float64
	UnconfirmedBalanceSat   int64
	UnconfirmedTxApperances int
	UnconfirmedAppearances  int
	TxApperances            int
	TxAppearances           int
	Transactions            []string
}
type BlockHashResponse string
type BestBlockHashResponse string
