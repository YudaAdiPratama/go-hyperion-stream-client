// interfaces.go

package client

import "time"

// SavedRequest represents a saved request
type SavedRequest struct {
	Type string      `json:"type"`
	Req  interface{} `json:"req"`
}

// HyperionClientOptions represents options for configuring the Hyperion client
type HyperionClientOptions struct {
	Endpoint  string `json:"endpoint"`
	ChainApi  string `json:"chainApi,omitempty"`
	Debug     bool   `json:"debug,omitempty"`
	LibStream bool   `json:"libStream,omitempty"`
}

// StreamDeltasRequest represents a request for streaming deltas
type StreamDeltasRequest struct {
	Code      string `json:"code"`
	Table     string `json:"table"`
	Scope     string `json:"scope"`
	Payer     string `json:"payer"`
	StartFrom string `json:"start_from"`
	ReadUntil string `json:"read_until"`
}

// RequestFilter represents a filter for requests
type RequestFilter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// StreamActionsRequest represents a request for streaming actions
type StreamActionsRequest struct {
	Contract  string          `json:"contract"`
	Account   string          `json:"account"`
	Action    string          `json:"action"`
	Filters   []RequestFilter `json:"filters"`
	StartFrom string          `json:"start_from"`
	ReadUntil string          `json:"read_until"`
}

// ActionContent represents content for action data
type ActionContent struct {
	Timestamp            time.Time              `json:"@timestamp"`
	GlobalSequence       int                    `json:"global_sequence"`
	AccountRamDeltas     map[string]interface{} `json:"account_ram_deltas"`
	Act                  ActContent             `json:"act"`
	BlockNum             int                    `json:"block_num"`
	ActionOrdinal        int                    `json:"action_ordinal"`
	CreatorActionOrdinal int                    `json:"creator_action_ordinal"`
	CpuUsageUs           int                    `json:"cpu_usage_us"`
	NetUsageWords        int                    `json:"net_usage_words"`
	CodeSequence         int                    `json:"code_sequence"`
	AbiSequence          int                    `json:"abi_sequence"`
	TrxId                string                 `json:"trx_id"`
	Producer             string                 `json:"producer"`
	Notified             string                 `json:"notified"`
	AdditionalFields     map[string]interface{} `json:"@"`
}

// ActContent represents the action content
type ActContent struct {
	Authorization struct {
		Permission string `json:"permission"`
		Actor      string `json:"actor"`
	} `json:"authorization"`
	Account string      `json:"account"`
	Name    string      `json:"name"`
	Data    interface{} `json:"data"`
}

// DeltaContent represents content for delta data
type DeltaContent struct {
	Code             string                 `json:"code"`
	Table            string                 `json:"table"`
	Scope            string                 `json:"scope"`
	Payer            string                 `json:"payer"`
	BlockNum         int                    `json:"block_num"`
	Data             map[string]interface{} `json:"data"`
	AdditionalFields map[string]interface{} `json:"@"`
}

// IncomingData represents incoming data
type IncomingData struct {
	Type         string      `json:"type"`
	Mode         string      `json:"mode"`
	Content      interface{} `json:"content"`
	Irreversible bool        `json:"irreversible"`
}

// LIBData represents LIB data
type LIBData struct {
	ChainID  string `json:"chain_id"`
	BlockNum int    `json:"block_num"`
	BlockID  string `json:"block_id"`
}

// ForkData represents fork data
type ForkData struct {
	ChainID       string `json:"chain_id"`
	StartingBlock int    `json:"starting_block"`
	EndingBlock   int    `json:"ending_block"`
	NewID         string `json:"new_id"`
}

// AsyncHandlerFunction represents an async handler function
type AsyncHandlerFunction func(data IncomingData) error

// EventData represents event data
type EventData interface{}

// EventListener represents an event listener
type EventListener func(data EventData)
