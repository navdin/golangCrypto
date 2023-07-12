package blockchain

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generatePrivateKey() string {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("error occured while creating private key.", err)
	}

	priKeyData := crypto.FromECDSA(priKey)
	priKeyString := hexutil.Encode(priKeyData)
	fmt.Println("private key as string: ", priKeyString)
	return priKeyString
}

func generatePublicKey(priKey *ecdsa.PrivateKey) string {

	pubKey := crypto.FromECDSAPub(&priKey.PublicKey)
	pubKeyString := hexutil.Encode(pubKey)
	fmt.Println("public key as string: ", pubKeyString)
	return pubKeyString
}

func generateAddressFromPriKey(priKey *ecdsa.PrivateKey) string{
	address := crypto.PubkeyToAddress(priKey.PublicKey).Hex()
	return address
}

func generatePriPubAdressKeys() (string, string, string) {
	priKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("error occured while creating private key.", err)
	}

	priKeyData := crypto.FromECDSA(priKey)
	priKeyString := hexutil.Encode(priKeyData)
	fmt.Println("private key as string: ", priKeyString)

	pubKey := crypto.FromECDSAPub(&priKey.PublicKey)
	pubKeyString := hexutil.Encode(pubKey)
	fmt.Println("public key as string: ", pubKeyString)

	addressString := crypto.PubkeyToAddress(priKey.PublicKey).Hex()

	return priKeyString, pubKeyString, addressString
}
