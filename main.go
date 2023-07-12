package main

import (
	"fmt"
	"strings"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go/crypto/blockchain"
	"github.com/go/crypto/cmdtemplate"
)


func main() {
	// cmdtemplate.DeleteWalletsFromPrevSession()
	// test()
	// return
	cmdtemplate.CreateClientTestNet()
	cmdtemplate.LoadMainAccount()
	cmdtemplate.DeleteWalletsFromPrevSession()

	var choice int
	for{
		fmt.Println()
		for _, item := range cmdtemplate.ListItems {
			fmt.Println(item)
		}
		fmt.Println()
		fmt.Println("Select an option from 0 to 9.")
		fmt.Scanln(&choice)
		cmdtemplate.CreateClientTestNet()
		switch choice {
		case 0 :
			cmdtemplate.F0()
		case 1 :
			cmdtemplate.F1()			
		case 2 :
			cmdtemplate.F2()
		case 3 :
			cmdtemplate.F3()
		case 4:
			cmdtemplate.F4()
		case 5:
			cmdtemplate.F5()
		case 6:
			cmdtemplate.F6()
		case 7:
			cmdtemplate.F7()
		case 8:
			cmdtemplate.F8()
		case 9:
			cmdtemplate.F9()
		default:
			fmt.Println("\nNo valid option selected.")
		}
		cmdtemplate.WishToContinue()
	}


}



func test(){
	//  amountF := new(big.Float)
	//  amountF.SetString(".099999998430674964")
	//  amountF64 := cmdtemplate.RoundOffToFloat64(amountF)
	//  fmt.Println("\n\namountF:   ", amountF)
	//  fmt.Println("amountF64: ", amountF64)
	 
	//  os := fmt.Sprintf("%v", amountF)
	//  fmt.Println("os:        ", os)
	//  amountF_64, _ := amountF.Float64()
	//  fmt.Println("amountF:   ",amountF)
	//  fmt.Println("amountF64: ",amountF_64)
	//  s := fmt.Sprintf("%v", amountF64)
	//  fmt.Println("last char: ",s[len(s)-1:])
	//  dotIndex := strings.Index(s, ".")
	//  factor := len(s)-1-dotIndex
	//  fmt.Println("factor: ",factor)
	//  fmt.Println()
	//  fmt.Println("float64(1/factor*10): ", float64(1/factor*10))
	// fmt.Println("sub: ", (amountF64*20))
	//  amountF64 = (amountF64*float64(factor)*10- 1)/float64(factor)*10
	 
	//  fmt.Println("amountF64 after roundoff: ", amountF64)
	// amount := big.NewInt((int64)(amountF64*1000000000000000000)-1)
	// fmt.Println("amount: ",amount)
	// fmt.Println("back amount in F64: ", amount)

	// amount := big.NewInt(int64((0.0999999984306749444 * 1000000000000000000)))
	// fmt.Println(amount)
	url := "/Users/navin/Documents/goCrypto/wallet/others/UTC--2023-07-12T20-02-02.451108000Z--52857bcdb2c6cb1c60131bc41ad67749863b1773"
	urlStartIndex := strings.Index(url, "/wallet")
	fmt.Println("urlStartIndex: ",urlStartIndex)
	urlPath := "."+url[urlStartIndex:]
	fmt.Println(urlPath)
	return

	// adr1 := common.HexToAddress("b2dd978ef739f2a8baaa995b66ee44d1c2a9255a")
	// fmt.Println("adr1: ", adr1)

	// "0xF27f127F8D636f73B1395266fCDFbee9cF599885"
//    ethClient := blockchain.GetEthBlockchainCustom(infuriaSepoliaUrl)

// fmt.Println("balance at ac1 infuraMainnetUrl: ",ethClient.GetBalanceAtAddress("f27f127f8d636f73b1395266fcdfbee9cf599885"))

fmt.Println("main ac hex to adr", common.HexToAddress("0x280FE7cF3849dc013F21b358A2C54D71b2a0CbC6"))
return
map1 := map[string]string{}
map1["a"]="b"
fmt.Println("before: ",map1)
for key,_ := range map1{
	delete(map1, key)
}
fmt.Println("after: ",map1)
cmdtemplate.LoadMainAccount()

// ethClient := blockchain.GetEthBlockchainLocalPanache()
// ethClient.GetEthClient()
// addr := "0x4E7B922B545986d7467492762dba55934385da7B"
// ethClient.GetBalanceAtAddress(addr)


ethClient := blockchain.GetEthBlockchainCustom(cmdtemplate.InfuraMainnetUrl)
addr := "0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7"
ethClient.GetEthClient()
ethClient.GetBalanceAtAddress(addr)

ac1 := blockchain.NewAccount("./wallet/others", "pass")
fmt.Println("ac1 addr string: ", ac1.AddressString)
ac2 := blockchain.NewAccount("./wallet/others", "pass")
fmt.Println("ac2 addr string: ", ac2.AddressString)

ethClient = blockchain.GetEthBlockchainCustom(cmdtemplate.InfuraMainnetUrl)
fmt.Println("balance at ac1 infuraMainnetUrl: ",ethClient.GetBalanceAtAddress(ac1.AddressString))
ethClient = blockchain.GetEthBlockchainCustom(cmdtemplate.InfuraGuerliUrl)
fmt.Println("balance at ac1 infuraGuerliUrl: ",ethClient.GetBalanceAtAddress(ac1.AddressString))
for _,val := range blockchain.AccountMap{
	fmt.Println()
	fmt.Println()
	fmt.Println("adrString: ",val.AddressString)
	fmt.Println("HexToAdr: ",common.HexToAddress(val.AddressString))
	fmt.Println("id: ",val.Id)
	fmt.Println("KeystoreLocation: ", val.KeystoreLocation)
	fmt.Println("Password: ", val.Password)
	fmt.Println("URLpath: ", val.URLpath)
	fmt.Println("priKey: ",val.GetPrivateKeyString())
	// if(key != "AC_MAIN"){
	// 	cmdtemplate.RemoveFileFromCurrentDir(val.URLpath)
	// }
}
}




