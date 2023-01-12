# go-moralis-sdk

Golang sdk to interact with Moralis API

## Installation
```shell
go get module github.com/viethapascal/go-moralis-sdk
```

## Usage
Get Moralis API Key from this [link](https://admin.moralis.io/web3apis) and set environment variable __MORALIS_API_KEY__
- Init MoralisAPI Instance:
```go
moralis, err := MoralisAPI()
if err != nil {
    log.Fatal(err)
}
```
- Get Multiple NFTs:
```go
tokens := MultipleNFTsRequest{[]Token{{
    TokenAddress: "0x06012c8cf97bead5deae237070f9587f8e7a266d",
    TokenID:      100,
}}}
nft, err := moralis.WithChainID("eth").NFT.GetMultipleNFTs(&tokens)
if err != nil {
    log.Fatal(err)
}
```

- Get NFT Collection:

```go
wallet := "0xd8da6bf26964af9d7eed9e03e53415d37aa96045"
ethNft := moralis.WithChainID("eth").NFT
collections, err := ethNft.GetNftCollection(wallet, UseDefaultQuery())
if err != nil {
    log.Fatal(err)
}
```

### Chain ID
| **Chain**                   | **String ID**                                                   |
|-----------------------------|-----------------------------------------------------------------|
| Ethereum Mainnet            | eth, mainnet, 0x1                                               |
| Goerli (Ethereum Testnet)   | goerli, 0x5                                                     |
| Binance Smart Chain Mainnet | bsc, binance, binance smart chain, 0x38                         |
| Binance Smart Chain Testnet | bsc testnet, binance testnet, binance smart chain testnet, 0x61 |
| Polygon (Matic) Mainnet     | matic, polygon, 0x89                                            |
| Mumbai (Matic Testnet)      | mumbai, matic testnet, polygon testnet, 0x13881                 |
| Avalanche Mainnet           | avalanche, 0xa86a                                               |
| Avalanche Testnet           | avalanche testnet, 0xa869                                       |
| Fantom Mainnet              | fantom, 0xfa                                                    |
| Cronos Mainnet              | cronos, 0x19                                                    |
