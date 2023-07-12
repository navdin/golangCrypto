package cmdtemplate

import (
	"fmt"
	"log"
	//"math"
	"math/big"
	"time"

	"github.com/go/crypto/blockchain"
)

type EthTransaction struct{
	fromAcc string
	toAcc string
	ethValue float64
	tnxCost string
	time string
	txHex string
}

// last N transactions to be maintained, change N here to change size of lastNTransactions
const NTxSize = 5

var lastNTransactions []EthTransaction = []EthTransaction{}
var mapAccLastNTranxs map[string][]EthTransaction = map[string][]EthTransaction{}

func TransferEthBetweenAccounts(fromAcc string, toAcc string, ethVal float64, enablePrints bool){
	if(ethVal == 0 || ethVal == 0.0){
		return
	}
	if(ethVal < MinTransferValue){
		fmt.Println("Minimum transfer value should be ", MinTransferValue," to include gas charges.")
		return 
	}
	res := CheckSufficientBalance(fromAcc, ethVal)
	if(res == false){
		if enablePrints{
      		fmt.Println("Insufficient balance in account ", fromAcc, ". ")
		}
		return
	}
	txHex, err, tnxCost := blockchain.MakeTransaction(fromAcc, toAcc, ethVal, *CreateClientTestNet().GetEthClient())
	if(err != nil){
		log.Fatal(err)
	}
	newTx := EthTransaction{
		fromAcc: fromAcc,
		toAcc: toAcc,
		ethValue: ethVal,
		tnxCost: tnxCost.String() + " wei",
		time: time.Now().Local().String(),
		txHex: txHex,
	}

	lastNTransactions = UpdateIfTxnSizeIsN(lastNTransactions, newTx)
	updateAccountLastNTranxs(newTx)
	if(enablePrints){
	    fmt.Println("Transaction completed.")
	}
}

func CheckSufficientBalance(fromAcc string, transferValue float64) (bool) {
	
	accBal := EthBlchain.GetBalanceAtAddress(blockchain.AccountMap[fromAcc].AddressString)
	transferValF := new(big.Float)
	transferValF.SetString(fmt.Sprintf("%v", transferValue))
	// transferValEth :=  new(big.Float).Quo(transferInt, big.NewFloat(math.Pow10(18)))
	res := accBal.Cmp(transferValF)
	fmt.Println("res: ", res)
	if res == -1 {
		return false
	}
	return true
}

func updateAccountLastNTranxs(newTnx EthTransaction) {

	lastN := mapAccLastNTranxs[newTnx.fromAcc]
	mapAccLastNTranxs[newTnx.fromAcc] = UpdateIfTxnSizeIsN(lastN, newTnx)

	lastN = mapAccLastNTranxs[newTnx.toAcc]
	mapAccLastNTranxs[newTnx.toAcc] = UpdateIfTxnSizeIsN(lastN, newTnx)

}

func UpdateIfTxnSizeIsN(txList []EthTransaction, newTx EthTransaction) []EthTransaction {

	if len(txList) == NTxSize {
		txList = txList[1:]
	}
	txList = append(txList, newTx)	
	return txList
}



