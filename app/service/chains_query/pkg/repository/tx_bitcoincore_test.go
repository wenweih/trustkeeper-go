package repository

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
	common "trustkeeper-go/library/util"

	"github.com/btcsuite/btcd/wire"
)

func TestExtractVoutScriptPubKeyt(t *testing.T) {
	txHex := "0100000001511609d82b6052dc7d3101aeabcd033e793f580a8779c033690cd4f1c28f972901000000020000ffffffff0300000000000000001a6a046f6d6e6902000002000004800000030800000000068e77802a2a0f00000000001976a914d7eee4fab2f5193aa3e0175e611e087d7ef5960f88ac00000000000000001976a914c9627232c0aecd6d59531434fdf1da683cafd24e88ac00000000"
	txByte, err := hex.DecodeString(txHex)
	if err != nil {
		t.Errorf("DecodeString: %v", err)
	}
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(bytes.NewReader(txByte)); err != nil {
		t.Errorf("DecodeString: %v", err)
	}
	for _, vout := range msgTx.TxOut {
		pkScriptHex := hex.EncodeToString(vout.PkScript)
		omniPropertyID := common.Hex2int(pkScriptHex[26:34])
		transferAmount := common.Hex2int(pkScriptHex[36:])
		fmt.Println("pkScriptHex", pkScriptHex, omniPropertyID, transferAmount)
	}
}
