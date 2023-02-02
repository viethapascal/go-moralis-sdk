package core

import (
	"encoding/json"
	"time"
)

type NFTDetail struct {
	TokenAddress      string          `json:"token_address"`
	TokenID           string          `json:"token_id"`
	Amount            string          `json:"amount"`
	OwnerOf           string          `json:"owner_of"`
	TokenHash         string          `json:"token_hash"`
	BlockNumberMinted string          `json:"block_number_minted"`
	BlockNumber       string          `json:"block_number"`
	ContractType      string          `json:"contract_type"`
	Name              string          `json:"name"`
	Symbol            string          `json:"symbol"`
	TokenURI          string          `json:"token_uri"`
	Metadata          string          `json:"metadata"`
	LastTokenURISync  time.Time       `json:"last_token_uri_sync"`
	LastMetadataSync  time.Time       `json:"last_metadata_sync"`
	MinterAddress     string          `json:"minter_address"`
	NormalizedData    *NormalizedData `json:"normalized_metadata,omitempty"`
}

type PaginationResult struct {
	Status   string `json:"status,omitempty"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
type WalletNFT struct {
	PaginationResult
	Result []NFTDetail `json:"result"`
}

type NormalizedData struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	ExternalLink string `json:"external_link"`
	AnimationURL string `json:"animation_url"`
	Attributes   []struct {
		TraitType   string `json:"trait_type"`
		DisplayType string `json:"display_type"`
		MaxValue    int    `json:"max_value"`
		TraitCount  int    `json:"trait_count"`
		Order       int    `json:"order"`
	} `json:"attributes"`
}

type Transaction struct {
	BlockNumber      string      `json:"block_number"`
	BlockTimestamp   time.Time   `json:"block_timestamp"`
	BlockHash        string      `json:"block_hash"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex int         `json:"transaction_index"`
	LogIndex         int         `json:"log_index"`
	Value            string      `json:"value"`
	ContractType     string      `json:"contract_type"`
	TransactionType  string      `json:"transaction_type"`
	TokenAddress     string      `json:"token_address"`
	TokenID          string      `json:"token_id"`
	FromAddress      string      `json:"from_address"`
	ToAddress        string      `json:"to_address"`
	Amount           string      `json:"amount"`
	Verified         int         `json:"verified"`
	Operator         interface{} `json:"operator"`
}

type NFTTransactions struct {
	PaginationResult
	Result []Transaction `json:"result"`
}

type Token struct {
	TokenAddress string `json:"token_address"`
	TokenID      int    `json:"token_id"`
}
type MultipleNFTsRequest struct {
	Tokens []Token `json:"tokens"`
}

func (d WalletNFT) String() string {
	bytes, _ := json.MarshalIndent(d, "", "\t")
	return string(bytes)
}

func (d NFTDetail) String() string {
	bytes, _ := json.MarshalIndent(d, "", "\t")
	return string(bytes)
}

func (d NFTTransactions) String() string {
	bytes, _ := json.MarshalIndent(d, "", "\t")
	return string(bytes)
}

type Collection struct {
	TokenAddress string      `json:"token_address"`
	ContractType string      `json:"contract_type"`
	Name         string      `json:"name"`
	Symbol       string      `json:"symbol"`
	NFTs         []NFTDetail `json:"nft_list"`
}

type NFTCollections struct {
	PaginationResult
	Result []Collection `json:"result"`
}
