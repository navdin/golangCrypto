package blockchain

import "math/big"

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// )

func IsZero(val big.Float) bool{
	big0 := new(big.Float)
	big0.SetString("0")
	if(val.Cmp(big0) == 0){
		return true;
	}
	return false
}










