package blockchain

import (
	"context"
	// "fmt"
	"io/ioutil"
	"log"
	"math/big"
	// "strconv"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const gasLimit = 21000

func MakeTransaction(account1 string, account2 string, value float64, blClinet ethclient.Client) (string, error){
	ac1Adrs := common.HexToAddress(AccountMap[account1].AddressString)
	ac2Adrs := common.HexToAddress(AccountMap[account2].AddressString)

	ac1Details := AccountMap[account1]

	nonce, err :=  blClinet.PendingNonceAt(context.Background(), ac1Adrs)
	if(err != nil){
		log.Fatal(err)
	}
	//1 ether = 10^18 wei
	amount := big.NewInt((int64)(value*1000000000000000000))

	gasPrice, err := blClinet.SuggestGasPrice(context.Background())
	// gasLimitBigF := new(big.Int)
	// gasLimitBigF.SetString(strconv.FormatInt( gasLimit, 10), 10)
	// fmt.Println("gasPrice: ", gasPrice)
	// fmt.Println("trnx cost: ", (gasPrice).Mul(gasPrice, gasLimitBigF))
	if(err != nil){
		log.Fatal(err)
	}

	trnx := types.NewTransaction(nonce, ac2Adrs, amount, gasLimit, gasPrice, nil)

	chainId, err := blClinet.NetworkID(context.Background())
	if(err != nil){
		log.Fatal(err)
	}

	keystoreBytes, err := ioutil.ReadFile(ac1Details.URLpath)
	if(err != nil){
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(keystoreBytes, ac1Details.Password)

	signedTx, err := types.SignTx(trnx, types.NewEIP155Signer(chainId), key.PrivateKey)
	if(err != nil){
		log.Fatal(err)
	}

	err = blClinet.SendTransaction(context.Background(), signedTx)
	if(err != nil){
		log.Fatal(err)
	}

	return signedTx.Hash().Hex(), nil
}