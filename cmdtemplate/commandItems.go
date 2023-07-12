package cmdtemplate

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/go/crypto/blockchain"
)

var ListItems = []string{
	"0.  CREATE WALLET",
	"1.  WALLET BALANCE",
	"2.  SHOW ALL WALLETS",
	"3.  LAST " + strconv.Itoa(NTxSize) + " TRANSACTIONS OF A WALLET",
	"4.  LAST " + strconv.Itoa(NTxSize) + " TRANSACTIONS OVERALL",
	"5.  ADD ETH TO WALLET FROM MAIN ACCOUNT",
	"6.  TRANSFER ETH BETWEEN TWO WALLETS",
	"7.  CHECK MAIN ACCOUNT BALANCE",
	"8.  DELETE WALLET",
	"9.  EXIT",
	"    ",
	"10. SMART-CONTRACTS (coming soon)",
}

func F0() {
	//create wallet
	fmt.Println("Are you sure you want to create new wallet? Enter y/n: ")
	var input string
	fmt.Scanln(&input)
	if input == "n" {
		return
	}
	CreateWallet()
}

func F1() {
	//wallet balance
	if !IsAnyWalletCreated() {
		return
	}
	got, validInput := takeValidWalletInput()
	if !got {
		return
	}

	bal := GetCurrentBalanceInAccount(validInput)
	fmt.Println("Balance in wallet ", validInput, " is: ", bal, " ETH")

}

func F2() {
	GetAllAccounts()
}

func F3() {
	if !IsAnyWalletCreated() {
		return
	}
	got, validInput := takeValidWalletInput()
	if !got {
		return
	}
	acDetails := blockchain.AccountMap[validInput]
	lastNTxs := mapAccLastNTranxs[acDetails.Id]
	if lastNTxs == nil {
		fmt.Println("No transactions made on wallet-id ", validInput)
		return
	}
	fmt.Println("Last upto ", NTxSize, " transactions made below:")
	for _, tx := range lastNTxs {
		TnxPrettyPrint(tx)
	}
}

func F4() {
	if !IsAnyWalletCreated() {
		return
	}
	fmt.Println("Last upto ", NTxSize, " transactions overall are below:")
	for _, tx := range lastNTransactions {
		TnxPrettyPrint(tx)
	}
}

func F5() {
	if !IsAnyWalletCreated() {
		return
	}
	goAhead := CheckBalaceInWallet(ACCOUNT_MAIN)
	if !goAhead {
		fmt.Println("Insufficient funds in main-account. \nClose and restart this program. \nIf problem persists then contact administrator to add more funds in main-account.")
	}
	got, validInput := takeValidWalletInput()
	if !got {
		return
	}
	fmt.Println("Enter only ", MinTransferValue, " amount. Maintain more balance in main-account.")
	ethValue, got := takeEthValueInput()
	if !got {
		return
	}
	TransferEthBetweenAccounts(ACCOUNT_MAIN, validInput, ethValue, true)
}

func F6() {
	if !IsAnyWalletCreated() {
		return
	}
	if blockchain.GetAccountSize() < 2 {
		fmt.Println("Only one wallet created. Create atleast 2 wallets.")
		return
	}
	fmt.Println("From-Wallet ")
	got, fromWallet := takeValidWalletInput()
	if !got {
		return
	}
	goAhead := CheckBalaceInWallet(fromWallet)
	if !goAhead {
		return
	}
	fmt.Println("To-Wallet ")
	got, toWallet := takeValidWalletInput()
	if !got {
		return
	}

	ethValue, got := takeEthValueInput()
	if !got {
		return
	}
	TransferEthBetweenAccounts(fromWallet, toWallet, ethValue, true)
}

func F7() {
	//main account balance
	bal := GetCurrentBalanceInAccount(ACCOUNT_MAIN)
	fmt.Println("Balance in MAIN_ACCOUNT is: ", bal, " ETH")
}

func F8() {
	// delete wallet
	if !IsAnyWalletCreated() {
		return
	}
	got, validInput := takeValidWalletInput()
	if !got {
		return
	}
	ethBal := GetCurrentBalanceInAccount(validInput)
	walletDetails := blockchain.AccountMap[validInput]
	ethValueFloat64, goAhead := MaintainMinBalance(ethBal)
	if goAhead {
		TransferEthBetweenAccounts(validInput, ACCOUNT_MAIN, ethValueFloat64, true)
	}
	RemoveFileFromCurrentDir(walletDetails.URLpath)
	delete(blockchain.AccountMap, walletDetails.Id)
	fmt.Println("Wallet ", validInput, " is deleted.")
}

func F9() {
	//exit
	transferBalancesToMainAccount(true)
	os.Exit(0)

}

func transferBalancesToMainAccount(enablePrints bool) {
	mainAcDetails := blockchain.AccountMap[ACCOUNT_MAIN]

	var waitgrp = sync.WaitGroup{}
	if enablePrints {
		fmt.Println("Processing... transfering balances in all wallets created in this session into main-account")
	}
	for key, val := range blockchain.AccountMap {
		if key == ACCOUNT_MAIN {
			continue
		}
		waitgrp.Add(1)
		go func(val blockchain.AccountDetails) {
			fromAcDetails := blockchain.AccountMap[val.Id]
			ethBal := GetCurrentBalanceInAccount(val.Id)
			ethValueFloat64, goAhead := MaintainMinBalance(ethBal)
			if goAhead {
				TransferEthBetweenAccounts(fromAcDetails.Id, mainAcDetails.Id, ethValueFloat64, false)
			}
			RemoveFileFromCurrentDir(fromAcDetails.URLpath)
			delete(blockchain.AccountMap, key)
			waitgrp.Done()
		}(val)
	}
	waitgrp.Wait()
	fmt.Println("Transactions Done.")
}

func MaintainMinBalance(ethBal *big.Float) (float64, bool) {
	ethValueF64, _ := ethBal.Float64()
	ethValueF64 = ethValueF64 - WalletMinBalance
	if ethValueF64 < MinTransferValue {
		return 0, false
	}
	return ethValueF64, true
}

func RoundOffToFloat64(ethBal *big.Float) float64 {

	ethValueF64, _ := ethBal.Float64()
	copyVal := new(big.Float)
	copyVal.SetString(fmt.Sprintf("%v", ethValueF64))

	if ethBal.Cmp(copyVal) != 0 {
		balStr := fmt.Sprintf("%v", ethBal)
		balF64str := fmt.Sprintf("%v", ethValueF64)
		lastCharBigF := balStr[len(balF64str)-1 : len(balF64str)]
		fmt.Println("lastCharBigF: ", lastCharBigF)
		lastCharF64 := balF64str[len(balF64str)-1:]
		fmt.Println("lastCharF64: ", lastCharF64)
		if lastCharBigF == lastCharF64 {
			return ethValueF64
		}
		balF64str = balF64str[0 : len(balF64str)-1]
		ethValueF64, _ = strconv.ParseFloat(balF64str, 64)
		fmt.Println("ethBal:      ", ethBal)
		fmt.Println("ethValueF64: ", ethValueF64)
	}
	return ethValueF64
}

func takeValidWalletInput() (bool, string) {
	count := 0
	for {
		count++
		fmt.Println("Enter wallet-id: ")
		fmt.Println("Created wallet-ids ", blockchain.AccountsList)

		var input string
		fmt.Scanln(&input)
		if input == ACCOUNT_MAIN {
			fmt.Println("Cannot edit main-account wallet.")
			if count < 3 {
				continue
			}
			return false, ""
		}
		_, got := blockchain.AccountMap[input]
		if !got {
			fmt.Println("Wallet not found. Try again")
			if count < 3 {
				continue
			}
			return false, ""
		}

		return true, input
	}
}

func takeEthValueInput() (float64, bool) {
	count := 0
	for {
		count++
		if count == 2 {
			fmt.Println("Max retries exceeded.")
			return -1, false
		}
		fmt.Println("Enter eth value: (enter value between 0.01 and 0.1)")

		var input float64
		fmt.Scanln(&input)

		if input == 0 || input == 0.0 {
			fmt.Println("Zero value entered. Enter a value greater than 0.01")
			continue
		}
		return input, true
	}
}

func IsAnyWalletCreated() bool {
	if blockchain.GetAccountSize() > 0 {
		return true
	}
	fmt.Println(NO_WALLET_CREATED)
	return false
}

func TnxPrettyPrint(tnx EthTransaction) {

	fmt.Println("From: ", tnx.fromAcc)
	fmt.Println("To: ", tnx.toAcc)
	fmt.Println("Amount: ", tnx.ethValue, " eth")
	fmt.Println("Time: ", tnx.time)
	fmt.Println("Hex: ", tnx.txHex)

}

func WishToContinue() {
	fmt.Println()
	fmt.Println("Do you wish to continue: (y/n)")
	var input string
	fmt.Scanln(&input)
	if input == "n" {
		F9()
	}
}

func LoadMainAccount() {

	URLpath, err := filepath.Abs("./wallet/main")
	if err != nil {
		panic(err)
	}
	mainAc := blockchain.AccountDetails{
		Id:               ACCOUNT_MAIN,
		AddressString:    "0x280FE7cF3849dc013F21b358A2C54D71b2a0CbC6",
		KeystoreLocation: "./wallet/main",
		Password:         "pass",
		URLpath:          URLpath + "/UTC--2023-06-29T19-28-33.790683000Z--280fe7cf3849dc013f21b358a2c54d71b2a0cbc6",
	}
	blockchain.AccountMap[ACCOUNT_MAIN] = mainAc
}

func RemoveFileFromCurrentDir(URLpath string) {
	path := filepath.Join(URLpath)
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteWalletsFromPrevSession() {
	fmt.Println("Transfering balances in wallets created in previous session into main account.")
	files, err := ioutil.ReadDir("./wallet/others/")
	if len(files) == 0 {
		fmt.Println("Done.")
		return
	}
	path := filepath.Join("./wallet/others/")
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		URLpath := path + "/" + file.Name()
		acDetails := blockchain.CreateAccountDetailsFromKeyFilepath(URLpath, "AC"+strconv.Itoa(i))
		blockchain.AccountMap[acDetails.Id] = acDetails
		// fmt.Println(URLpath, file.IsDir())
	}
	transferBalancesToMainAccount(false)
	fmt.Println("Transactions done.")
}

func MoreThanMinTransferValue(val big.Float) bool {
	bigMinTrans := new(big.Float)
	bigMinTrans.SetString(fmt.Sprintf("%v", MinTransferValue))
	if val.Cmp(bigMinTrans) == -1 {
		return false
	}
	return true
}

func CheckBalaceInWallet(walletId string) bool {
	bal := GetCurrentBalanceInAccount(walletId)
	fmt.Println("Balance in wallet ", walletId, ":", bal, " ETH.")
	balF64, _ := bal.Float64()
	transferableBal := balF64 - WalletMinBalance
	if transferableBal >= MinTransferValue {
		fmt.Println("Transferable Balance: ", transferableBal)
		return true
	} else {
		fmt.Println("Transferable balance is less than minimum transfer amount. Please add ", MinTransferValue, " ETH to from-wallet ", walletId)
	}
	return false
}
