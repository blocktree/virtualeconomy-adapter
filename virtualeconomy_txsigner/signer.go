package virtualeconomy_txsigner

import (
	"fmt"
	"github.com/blocktree/go-owcrypt"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {

}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {

	signature, ret := owcrypt.Signature(privateKey, nil, 0, msg, uint16(len(msg)), owcrypt.ECC_CURVE_X25519)
	if ret != owcrypt.SUCCESS {
		return nil, fmt.Errorf("ECC sign hash failed")
	}
	return signature,nil
}
