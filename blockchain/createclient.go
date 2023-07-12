package blockchain

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthBlockChain struct {
	client *ethclient.Client
}

const panacheUrl = "http://localhost:8545"

func GetEthBlockchainLocalPanache() *EthBlockChain {
	EthClient, err := ethclient.DialContext(context.Background(), panacheUrl)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	EthBlchain := EthBlockChain{
		client: EthClient,
	}
	return &EthBlchain
}

func GetEthBlockchainCustom(blockchainUrl string) *EthBlockChain {
	EthClient, err := ethclient.DialContext(context.Background(), blockchainUrl)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	EthBlchain := EthBlockChain{
		client: EthClient,
	}
	return &EthBlchain
}

func (eth *EthBlockChain) GetEthClient() *ethclient.Client {
	id, err := eth.client.NetworkID(context.Background())
	if(err != nil){
		log.Fatal("error occured ", err)
	}
	fmt.Println("NetworkID: ", id)
	fmt.Println("ethClient: ", eth.client)
	return eth.client
}


func (eth *EthBlockChain) GetBalanceAtAddress(addr string) *big.Float {
	adress := common.HexToAddress(addr)
	balance, err := eth.client.BalanceAt(context.Background(), adress, nil)

	if err != nil {
		log.Fatal("error occured ", err)
	}
	fmt.Println("balance: ", balance)
	balancef := new(big.Float)
	balancef.SetString(balance.String())
	balanceEth := new(big.Float).Quo(balancef, big.NewFloat(math.Pow10(18)))
	
	return balanceEth
}

