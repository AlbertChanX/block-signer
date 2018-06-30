package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func main() {
	address := "msXzK5c57uo6aP5pHoaBnGzSbKcmmh9rYe"
	var balance int64 = 65000000 // balance
	var fee int64 = 0.001 * 1e8  // fee
	var leftToMe = balance - fee // keep the change

	// 1. outputs
	outputs := []*wire.TxOut{}

	// 1.1 输出1, 给自己转剩下的钱
	addr, _ := btcutil.DecodeAddress(address, &chaincfg.SimNetParams)
	pkScript, _ := txscript.PayToAddrScript(addr)
	outputs = append(outputs, wire.NewTxOut(leftToMe, pkScript))

	// 1.2 输出2, add comment
	comment := "I Love U"
	pkScript, _ = txscript.NullDataScript([]byte(comment))
	outputs = append(outputs, wire.NewTxOut(int64(0), pkScript))

	// 2. inputs
	prevTxHash := "631f147c89d78e56e4de4163a0e1e8b9cb486a9bc75afcb3fd24bb52fe8db5bd"
	prevPkScriptHex := "76a91483d39c122c7c38b74a60d59cd44e9cfe57f122ae88ac"
	prevTxOutputN := uint32(0)

	hash, _ := chainhash.NewHashFromStr(prevTxHash)   // tx hash
	outPoint := wire.NewOutPoint(hash, prevTxOutputN) // 第几个输出
	txIn := wire.NewTxIn(outPoint, nil, nil)
	inputs := []*wire.TxIn{txIn}

	prevPkScript, _ := hex.DecodeString(prevPkScriptHex)
	prevPkScripts := make([][]byte, 1)
	prevPkScripts[0] = prevPkScript

	tx := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}

	// 3. sign
	privKey := "ur pk" // pk
	sign(tx, privKey, prevPkScripts)

	// 4. 输出Hex
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err := tx.Serialize(buf); err != nil {
	}
	txHex := hex.EncodeToString(buf.Bytes())
	fmt.Println("hex", txHex)
}

// sign for tx
func sign(tx *wire.MsgTx, privKeyStr string, prevPkScripts [][]byte) {
	inputs := tx.TxIn
	wif, err := btcutil.DecodeWIF(privKeyStr)
	fmt.Println("wif err", err)
	privKey := wif.PrivKey

	for i := range inputs {
		pkScript := prevPkScripts[i]
		var script []byte
		script, err = txscript.SignatureScript(tx, i, pkScript, txscript.SigHashAll,
			privKey, false)
		inputs[i].SignatureScript = script
	}
}
