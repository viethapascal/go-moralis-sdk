package core

import (
	"fmt"
)

type NFTAPI struct {
	*Moralis
}

func newNFTApi(m *Moralis) *NFTAPI {
	return &NFTAPI{m}
}

func (n *NFTAPI) GetNFTByWallet(wallet string, opts ...RequestOption) (*WalletNFT, error) {
	//urlPath := fmt.Sprintf("%s/%s/nft?chain=%s", n.APIUrl, wallet, n.ChainID)
	urlPath := n.Uri.Encode("getNFTs", map[string]string{"address": wallet})
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
