package blockchain

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type AccountDetails struct {
	AddressString    string
	KeystoreLocation string
	URLpath          string
	Password         string
	Id               string
}

var AccountMap = make(map[string]AccountDetails)
var AccountsList = []string{}
var accountCount = 0

func NewAccount(keystoreLocation string, pass string) *AccountDetails {
	key := keystore.NewKeyStore(keystoreLocation, keystore.StandardScryptN, keystore.StandardScryptP)
	accountNew, err := key.NewAccount(pass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new wallet address: ", accountNew.Address.Hex())
	accountCount = accountCount+1;

	urlStartIndex := strings.Index(accountNew.URL.Path, "/wallet")
	urlPath := "."+accountNew.URL.Path[urlStartIndex:]
	acDets := AccountDetails{
		AddressString:          accountNew.Address.Hex(),
		KeystoreLocation: keystoreLocation,
		URLpath:          urlPath,
		Password:         pass,
		Id : "AC"+strconv.Itoa(accountCount),
	}
	AccountMap[acDets.Id] = acDets
	AccountsList = append(AccountsList, acDets.Id)
	return &acDets
}


func (accountDts *AccountDetails) GetPrivateKeyString() string {
	key := accountDts.getKey()
	priKeyData := crypto.FromECDSA(key.PrivateKey)

	priKeyString := hexutil.Encode(priKeyData)
	// fmt.Println("private key string: ", priKeyString)
	return priKeyString
}

func (accountDts *AccountDetails) getKey() keystore.Key {
	fileBytes, err := ioutil.ReadFile(accountDts.URLpath)
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(fileBytes, accountDts.Password)
	if err != nil {
		log.Fatal(err)
	}
	return *key
}


func (accountDts *AccountDetails) getPublicKeyString() string{
	pubKeyData := crypto.FromECDSAPub(&accountDts.getKey().PrivateKey.PublicKey)
	pubKeyString := hexutil.Encode(pubKeyData)
	fmt.Println("public key as string: ", pubKeyString)
	return pubKeyString
}

func (accountDts *AccountDetails) getAddressString() string{
	key := accountDts.getKey()
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	return address
}

func GetAccountSize() int16 {
	return int16(accountCount)
}

func CreateAccountDetailsFromKeyFilepath(URLpath string, id string) AccountDetails {
	fileBytes, err := ioutil.ReadFile(URLpath)
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(fileBytes, "pass")
	if err != nil {
		log.Fatal(err)
	}
	return AccountDetails{
		AddressString: key.Address.Hex(),
		KeystoreLocation: "./wallet/others",
		URLpath: URLpath,
		Password: "pass",
		Id: id,
	}
}


