package core

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"log"
	"regexp"
	"testing"
)

func TestNFT(t *testing.T) {
	moralis, err := MoralisAPI()
	log.Println("APIKEY:", moralis.apiKey)
	//moralis = moralis.WithChainID("eth")
	if err != nil {
		log.Fatal(err)
	}
	wallet := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	ethNft := moralis.WithChainID("eth").NFT
	nftByWallet, err := ethNft.GetNFTByWallet(
		wallet,
		DisableTotal(true),
		Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nftByWallet)

}

func TestWithDefaultParams(t *testing.T) {
	opt := RequestQuery{
		Page:           1,
		Limit:          100,
		DisableTotal:   false,
		TokenAddresses: nil,
	}
	v, _ := query.Values(opt)
	log.Println(v.Encode())
}

func TestNFTAPI_GetMultipleNFTs(t *testing.T) {
	tokens := MultipleNFTsRequest{[]Token{{
		TokenAddress: "0x06012c8cf97bead5deae237070f9587f8e7a266d",
		TokenID:      100,
	}}}
	bytes, err := json.Marshal(tokens)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))
	moralis, err := MoralisAPI()
	if err != nil {
		log.Fatal(err)
	}
	nft, err := moralis.WithChainID("eth").NFT.GetMultipleNFTs(&tokens)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nft)
}

func TestNFTAPI_GetNFTTransfer(t *testing.T) {
	moralis, err := MoralisAPI()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("APIKEY:", moralis.apiKey)

	wallet := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	ethNft := moralis.WithChainID("eth").NFT
	nftByWallet, err := ethNft.GetNFTTransfer(
		wallet,
		DisableTotal(true),
		Page(2),
		Query(map[string]interface{}{"direction": "both", "format": "decimal"}),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nftByWallet)
}

func TestDefaultQuery(t *testing.T) {
	path := "/{address}/events/{id}"
	re, _ := regexp.Compile(`\{([a-z_]+)}`)
	log.Println(re.FindString(path))
	groups := re.FindAllStringSubmatch(path, -1)
	//assert.Len(t, groups, 2)
	log.Println("groups:", groups)
	//assert.Equal(t, "address", groups[1])
}

//func TestUrlBuilder_UrlOf(t *testing.T) {
//	apis := EndpointList{
//		EndpointData{
//			Endpoint:      "abc",
//			Path:          "/{address}/nft/transfers",
//			Price:         0,
//			RateLimitCost: 0,
//		},
//	}
//	ep := apis.GetEndPoint("abc")
//	log.Println(ep.Encode(map[string]string{"address": "0x0000"}))
//	assert.Equal(t, "/0x0000/nft/transfers", ep.Encode(map[string]string{"address": "0x0000"}))
//}

func TestNFTAPI_GetNftCollection(t *testing.T) {
	moralis, err := MoralisAPI()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("APIKEY:", moralis.apiKey)

	//wallet := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	wallet := "0xeBC6d76Fa16545e5eF99d78423a8108de1932bd5"
	ethNft := moralis.WithChainID("mumbai").NFT
	collections, err := ethNft.GetNftCollection(wallet, true, Normalize())
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := json.MarshalIndent(collections, "", "\t")
	log.Println(string(bytes))
	// Get metadata
	//metadata, err := ethNft.GetNFTMetadata(
	//	"0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB", "1",
	//	UseDefaultQuery(),
	//	Normalize())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bytes, err = json.MarshalIndent(metadata, "", "\t")
	//log.Println(string(bytes))
	//collectionMetadata, err := moralis.WithChainID("polygon").NFT.GetCollectionMetadata("0xa9a6a3626993d487d2dbda3173cf58ca1a9d9e9f")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bytes, err = json.MarshalIndent(collectionMetadata, "", "\t")
	//log.Println(string(bytes))
}

func TestMoralis_WithChainID(t *testing.T) {
	chainID := "avalanche testnet"
	wallet := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
	moralis, err := MoralisAPI()
	if err != nil {
		log.Fatal(err)
	}
	_, err = moralis.WithChainID(chainID).NFT.GetNftCollection(wallet, false, Limit(1), Normalize())

	if err != nil {
		log.Fatal(err)
	}
}
