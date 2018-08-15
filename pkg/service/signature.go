package service

import (
	"github.com/hiromaily/go-bitcoin/pkg/file"
	"github.com/pkg/errors"
	"log"
)

// coldwallet側から未署名トランザクションを読み込み、署名を行う

// signatureByHex 署名する
// オフラインで使うことを想定
func (w *Wallet) signatureByHex(hex string) (string, bool, error) {
	//first hex: 未署名トランザクションのhex
	// Hexからトランザクションを取得
	msgTx, err := w.Btc.ToMsgTx(hex)
	if err != nil {
		return "", false, err
	}

	// 署名
	signedTx, isSigned, err := w.Btc.SignRawTransaction(msgTx)
	if err != nil {
		return "", false, err
	}
	log.Printf("[Debug] isSigned01 is %t", isSigned)

	hexTx, err := w.Btc.ToHex(signedTx)
	if err != nil {
		return "", false, errors.Errorf("w.Btc.ToHex(msgTx): error: %v", err)
	}

	return hexTx, isSigned, nil

}

// SignatureByHex Hex文字列から署名を行う
func (w *Wallet) SignatureByHex(hex string) (string, bool, error) {
	//署名
	hexTx, isSigned, err := w.signatureByHex(hex)
	if err != nil {
		return "", isSigned, err
	}
	//log.Println("hex:", hexTx)

	//ファイルに書き込む
	file.WriteFileForSigned(hexTx)

	return hexTx, isSigned, nil
}

// SignatureFromFile 渡されたファイルからtransactionを読み取り、署名を行う
func (w *Wallet) SignatureFromFile(filePath string) (string, bool, error) {
	//ファイルからhexを読み取る
	hex, err := file.ReadFile(filePath)
	if err != nil {
		return "", false, err
	}

	//署名
	hexTx, isSigned, err := w.signatureByHex(hex)
	if err != nil {
		return "", isSigned, err
	}
	//log.Println("hex:", hexTx)

	//ファイルに書き込む
	file.WriteFileForSigned(hexTx)

	return hexTx, isSigned, nil
}
