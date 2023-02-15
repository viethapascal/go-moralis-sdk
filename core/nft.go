package core

import (
	"fmt"
	"log"
)

type NFTAPI struct {
	*Moralis
}

func newNFTApi(m *Moralis) *NFTAPI {
	return &NFTAPI{m}
}

func (n *NFTAPI) GetNFTByWallet(wallet string, opts ...RequestOption) (*WalletNFT, error) {
	//urlPath := fmt.Sprintf("%s/%s/nft?chain=%s", n.APIUrl, wallet, n.ChainID)
	urlPath := n.Uri.Encode("getNFTs", Params{"address": wallet})
	result := &WalletNFT{}
	err := n.Get(urlPath, result, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *NFTAPI) GetMultipleNFTs(tokens *MultipleNFTsRequest, opts ...RequestOption) ([]NFTDetail, error) {
	//urlPath := fmt.Sprintf("%s/nft/getMultipleNFTs?chain=%s", n.APIUrl, n.ChainID)
	urlPath := n.Uri.Encode("getMultipleNFTs", nil)
	result := make([]NFTDetail, 0)
	err := n.Post(urlPath, tokens, &result, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (n *NFTAPI) GetNFTTransfer(wallet string, opts ...RequestOption) (*NFTTransactions, error) {
	urlPath := fmt.Sprintf("%s/%s/nft/transfers?chain=%s", n.APIUrl, wallet, n.ChainID)
	result := &NFTTransactions{}
	err := n.Get(urlPath, result, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Collection APIs
func (n *NFTAPI) GetNftCollection(wallet string, includeNft bool, opts ...RequestOption) (*NFTCollections, error) {
	//urlPath := n.Uri.Encode("getWalletNFTCollections", map[string]string{"address": wallet})
	urlPath := fmt.Sprintf("%s/%s/nft/collections?chain=%s", n.APIUrl, wallet, n.ChainID)
	result := &NFTCollections{}
	err := n.Get(urlPath, result, opts...)
	if err != nil {
		return nil, err
	}
	if includeNft && len(result.Result) > 0 {
		tokens := make([]string, 0)
		for _, col := range result.Result {
			tokens = append(tokens, col.TokenAddress)
		}
		log.Println("tokens:", tokens)
		log.Println(opts)
		nftByWallet, err := n.GetNFTByWallet(wallet, TokenAddresses(tokens...), Normalize())
		if err != nil {
			return nil, err
		}
		for i := range result.Result {
			result.Result[i].Total = len(nftByWallet.Result)
			result.Result[i].NFTs = nftByWallet.Result
		}
	}
	return result, nil
}

func (n *NFTAPI) GetNFTMetadata(contract string, tokenId string, opts ...RequestOption) (*NFTDetail, error) {
	urlPath := n.Uri.Encode("getTokenIdMetadata", Params{"address": contract, "token_id": tokenId})
	result := &NFTDetail{}
	err := n.Get(urlPath, result, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (n *NFTAPI) GetCollectionMetadata(contract string, includeNft bool, opts ...RequestOption) (*Collection, error) {
	urlPath := n.Uri.Encode("getNFTMetadata", Params{"address": contract})
	result := &Collection{}
	err := n.Get(urlPath, result, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
