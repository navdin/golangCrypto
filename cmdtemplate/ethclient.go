package cmdtemplate

import (
	"github.com/go/crypto/blockchain"
)



var EthBlchain *blockchain.EthBlockChain

func CreateClientTestNet() *blockchain.EthBlockChain {
	if(EthBlchain != nil){
		return EthBlchain
	}
	EthBlchain = blockchain.GetEthBlockchainCustom(CustomBlockchainUrl)
	return EthBlchain
}

