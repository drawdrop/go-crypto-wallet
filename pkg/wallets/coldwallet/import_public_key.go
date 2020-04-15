package coldwallet

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/address"
	"github.com/hiromaily/go-bitcoin/pkg/model/rdb/coldrepo"
	"github.com/hiromaily/go-bitcoin/pkg/wallets/types"
)

// ImportPubKey import pubKey from csv file for sign wallet
//  only multisig account is available
func (w *ColdWallet) ImportPubKey(fileName string, accountType account.AccountType) error {
	if w.wtype != types.WalletTypeSignature {
		return errors.New("it's available on sign wallet")
	}

	//validate account, only multisig account is ok
	if !account.AccountTypeMultisig[accountType] {
		w.logger.Info("only multisig account is allowed")
		return nil
	}

	//validate file name
	if err := w.addrFileRepo.ValidateFilePath(fileName, accountType); err != nil {
		return err
	}

	// read file for full public key
	pubKeys, err := w.addrFileRepo.ImportPubKey(fileName)
	if err != nil {
		return errors.Wrapf(err, "fail to call fileStorager.ImportPubKey() fileName: %s", fileName)
	}

	// insert full pubKey into added_pubkey_history_table
	addedPubkeyHistorys := make([]coldrepo.AddedPubkeyHistoryTable, len(pubKeys))
	for i, key := range pubKeys {
		inner := strings.Split(key, ",")
		//FullPublicKey is required
		addedPubkeyHistorys[i] = coldrepo.AddedPubkeyHistoryTable{
			FullPublicKey:         inner[2],
			AuthAddress1:          "",
			AuthAddress2:          "",
			WalletMultisigAddress: "",
			RedeemScript:          "",
		}
	}
	//TODO:Upsert would be better to prevent error which occur when data is already inserted
	err = w.storager.InsertAddedPubkeyHistoryTable(accountType, addedPubkeyHistorys, nil, true)
	if err != nil {
		return errors.Wrap(err, "fail to call storager.InsertAddedPubkeyHistoryTable()")
	}

	return nil
}

// ImportMultisigAddress import multisig address generated by sign wallet into database
func (w *ColdWallet) ImportMultisigAddress(fileName string, accountType account.AccountType) error {
	if w.wtype != types.WalletTypeKeyGen {
		return errors.New("it's available on keygen wallet")
	}
	// validate
	if !account.AccountTypeMultisig[accountType] {
		w.logger.Info("only multisig account is allowed")
		return nil
	}
	if err := w.addrFileRepo.ValidateFilePath(fileName, accountType); err != nil {
		return err
	}

	// read file for full public key
	pubKeys, err := w.addrFileRepo.ImportPubKey(fileName)
	if err != nil {
		return errors.Errorf("key.ImportPubKey() error: %s", err)
	}

	//added_pubkey_history_receiptテーブルにInsert
	accountKeyTable := make([]coldrepo.AccountKeyTable, len(pubKeys))

	tm := time.Now()
	for i, pubkey := range pubKeys {
		inner := strings.Split(pubkey, ",")
		// csv file structure
		//	record.FullPublicKey,
		//	record.AuthAddress1,
		//	record.AuthAddress2,
		//	record.WalletMultisigAddress,
		//	record.RedeemScript,
		accountKeyTable[i] = coldrepo.AccountKeyTable{
			FullPublicKey:         inner[0],
			WalletMultisigAddress: inner[3],
			RedeemScript:          inner[4],
			AddrStatus:            address.AddrStatusValue[address.AddrStatusMultiAddressImported],
			UpdatedAt:             &tm,
		}
	}
	//TODO: Upsert would be better??
	err = w.storager.UpdateMultisigAddrOnAccountKeyTableByFullPubKey(accountType, accountKeyTable, nil, true)
	if err != nil {
		return errors.Errorf("DB.UpdateMultisigAddrOnAccountKeyTableByFullPubKey() error: %s", err)
	}

	return nil
}
