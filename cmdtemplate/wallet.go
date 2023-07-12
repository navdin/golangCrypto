package cmdtemplate

import (
	"fmt"
	"math/big"

	"github.com/go/crypto/blockchain"
)


func CreateWallet(){
	
	fmt.Println("Creating wallet...")
	acDetails := blockchain.NewAccount("./wallet/others/", "pass")
	fmt.Println("Account or wallet created. ID is: ", acDetails.Id)
	fmt.Println("You can use this ID for further transactions.")
}

func GetAllAccounts(){
	fmt.Println("Wallets created in this session are: ", blockchain.AccountsList)
	for _,val := range blockchain.AccountMap{
		fmt.Println()
		fmt.Println("id: ",val.Id)
		fmt.Println("adrString: ",val.AddressString)
		fmt.Println("KeystoreLocation: ", val.KeystoreLocation)
		fmt.Println("Password: ", val.Password)
		fmt.Println("URLpath: ", val.URLpath)
		// fmt.Println("priKey: ",val.GetPrivateKeyString())
	}
}


func GetCurrentBalanceInAccount(account string) *big.Float {
	accBal := EthBlchain.GetBalanceAtAddress(blockchain.AccountMap[account].AddressString)
	return accBal
}

