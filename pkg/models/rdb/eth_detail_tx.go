// Code generated by SQLBoiler 3.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// EthDetailTX is an object representing the database table.
type EthDetailTX struct {
	ID              int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	TXID            int64         `boil:"tx_id" json:"tx_id" toml:"tx_id" yaml:"tx_id"`
	SenderAccount   string        `boil:"sender_account" json:"sender_account" toml:"sender_account" yaml:"sender_account"`
	SenderAddress   string        `boil:"sender_address" json:"sender_address" toml:"sender_address" yaml:"sender_address"`
	SenderAmount    types.Decimal `boil:"sender_amount" json:"sender_amount" toml:"sender_amount" yaml:"sender_amount"`
	ReceiverAccount string        `boil:"receiver_account" json:"receiver_account" toml:"receiver_account" yaml:"receiver_account"`
	ReceiverAddress string        `boil:"receiver_address" json:"receiver_address" toml:"receiver_address" yaml:"receiver_address"`
	ReceiverAmount  types.Decimal `boil:"receiver_amount" json:"receiver_amount" toml:"receiver_amount" yaml:"receiver_amount"`
	Fee             types.Decimal `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	GasLimit        uint32        `boil:"gas_limit" json:"gas_limit" toml:"gas_limit" yaml:"gas_limit"`
	Nonce           int64         `boil:"nonce" json:"nonce" toml:"nonce" yaml:"nonce"`
	UnsignedHexTX   string        `boil:"unsigned_hex_tx" json:"unsigned_hex_tx" toml:"unsigned_hex_tx" yaml:"unsigned_hex_tx"`
	SignedHexTX     string        `boil:"signed_hex_tx" json:"signed_hex_tx" toml:"signed_hex_tx" yaml:"signed_hex_tx"`
	SentHashTX      string        `boil:"sent_hash_tx" json:"sent_hash_tx" toml:"sent_hash_tx" yaml:"sent_hash_tx"`
	UpdatedAt       null.Time     `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *ethDetailTXR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ethDetailTXL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EthDetailTXColumns = struct {
	ID              string
	TXID            string
	SenderAccount   string
	SenderAddress   string
	SenderAmount    string
	ReceiverAccount string
	ReceiverAddress string
	ReceiverAmount  string
	Fee             string
	GasLimit        string
	Nonce           string
	UnsignedHexTX   string
	SignedHexTX     string
	SentHashTX      string
	UpdatedAt       string
}{
	ID:              "id",
	TXID:            "tx_id",
	SenderAccount:   "sender_account",
	SenderAddress:   "sender_address",
	SenderAmount:    "sender_amount",
	ReceiverAccount: "receiver_account",
	ReceiverAddress: "receiver_address",
	ReceiverAmount:  "receiver_amount",
	Fee:             "fee",
	GasLimit:        "gas_limit",
	Nonce:           "nonce",
	UnsignedHexTX:   "unsigned_hex_tx",
	SignedHexTX:     "signed_hex_tx",
	SentHashTX:      "sent_hash_tx",
	UpdatedAt:       "updated_at",
}

// Generated where

var EthDetailTXWhere = struct {
	ID              whereHelperint64
	TXID            whereHelperint64
	SenderAccount   whereHelperstring
	SenderAddress   whereHelperstring
	SenderAmount    whereHelpertypes_Decimal
	ReceiverAccount whereHelperstring
	ReceiverAddress whereHelperstring
	ReceiverAmount  whereHelpertypes_Decimal
	Fee             whereHelpertypes_Decimal
	GasLimit        whereHelperuint32
	Nonce           whereHelperint64
	UnsignedHexTX   whereHelperstring
	SignedHexTX     whereHelperstring
	SentHashTX      whereHelperstring
	UpdatedAt       whereHelpernull_Time
}{
	ID:              whereHelperint64{field: "`eth_detail_tx`.`id`"},
	TXID:            whereHelperint64{field: "`eth_detail_tx`.`tx_id`"},
	SenderAccount:   whereHelperstring{field: "`eth_detail_tx`.`sender_account`"},
	SenderAddress:   whereHelperstring{field: "`eth_detail_tx`.`sender_address`"},
	SenderAmount:    whereHelpertypes_Decimal{field: "`eth_detail_tx`.`sender_amount`"},
	ReceiverAccount: whereHelperstring{field: "`eth_detail_tx`.`receiver_account`"},
	ReceiverAddress: whereHelperstring{field: "`eth_detail_tx`.`receiver_address`"},
	ReceiverAmount:  whereHelpertypes_Decimal{field: "`eth_detail_tx`.`receiver_amount`"},
	Fee:             whereHelpertypes_Decimal{field: "`eth_detail_tx`.`fee`"},
	GasLimit:        whereHelperuint32{field: "`eth_detail_tx`.`gas_limit`"},
	Nonce:           whereHelperint64{field: "`eth_detail_tx`.`nonce`"},
	UnsignedHexTX:   whereHelperstring{field: "`eth_detail_tx`.`unsigned_hex_tx`"},
	SignedHexTX:     whereHelperstring{field: "`eth_detail_tx`.`signed_hex_tx`"},
	SentHashTX:      whereHelperstring{field: "`eth_detail_tx`.`sent_hash_tx`"},
	UpdatedAt:       whereHelpernull_Time{field: "`eth_detail_tx`.`updated_at`"},
}

// EthDetailTXRels is where relationship names are stored.
var EthDetailTXRels = struct {
}{}

// ethDetailTXR is where relationships are stored.
type ethDetailTXR struct {
}

// NewStruct creates a new relationship struct
func (*ethDetailTXR) NewStruct() *ethDetailTXR {
	return &ethDetailTXR{}
}

// ethDetailTXL is where Load methods for each relationship are stored.
type ethDetailTXL struct{}

var (
	ethDetailTXAllColumns            = []string{"id", "tx_id", "sender_account", "sender_address", "sender_amount", "receiver_account", "receiver_address", "receiver_amount", "fee", "gas_limit", "nonce", "unsigned_hex_tx", "signed_hex_tx", "sent_hash_tx", "updated_at"}
	ethDetailTXColumnsWithoutDefault = []string{"tx_id", "sender_account", "sender_address", "sender_amount", "receiver_account", "receiver_address", "receiver_amount", "fee", "gas_limit", "nonce", "unsigned_hex_tx", "signed_hex_tx", "sent_hash_tx"}
	ethDetailTXColumnsWithDefault    = []string{"id", "updated_at"}
	ethDetailTXPrimaryKeyColumns     = []string{"id"}
)

type (
	// EthDetailTXSlice is an alias for a slice of pointers to EthDetailTX.
	// This should generally be used opposed to []EthDetailTX.
	EthDetailTXSlice []*EthDetailTX

	ethDetailTXQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ethDetailTXType                 = reflect.TypeOf(&EthDetailTX{})
	ethDetailTXMapping              = queries.MakeStructMapping(ethDetailTXType)
	ethDetailTXPrimaryKeyMapping, _ = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, ethDetailTXPrimaryKeyColumns)
	ethDetailTXInsertCacheMut       sync.RWMutex
	ethDetailTXInsertCache          = make(map[string]insertCache)
	ethDetailTXUpdateCacheMut       sync.RWMutex
	ethDetailTXUpdateCache          = make(map[string]updateCache)
	ethDetailTXUpsertCacheMut       sync.RWMutex
	ethDetailTXUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single ethDetailTX record from the query.
func (q ethDetailTXQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EthDetailTX, error) {
	o := &EthDetailTX{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for eth_detail_tx")
	}

	return o, nil
}

// All returns all EthDetailTX records from the query.
func (q ethDetailTXQuery) All(ctx context.Context, exec boil.ContextExecutor) (EthDetailTXSlice, error) {
	var o []*EthDetailTX

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EthDetailTX slice")
	}

	return o, nil
}

// Count returns the count of all EthDetailTX records in the query.
func (q ethDetailTXQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count eth_detail_tx rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ethDetailTXQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if eth_detail_tx exists")
	}

	return count > 0, nil
}

// EthDetailTxes retrieves all the records using an executor.
func EthDetailTxes(mods ...qm.QueryMod) ethDetailTXQuery {
	mods = append(mods, qm.From("`eth_detail_tx`"))
	return ethDetailTXQuery{NewQuery(mods...)}
}

// FindEthDetailTX retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEthDetailTX(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*EthDetailTX, error) {
	ethDetailTXObj := &EthDetailTX{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `eth_detail_tx` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ethDetailTXObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from eth_detail_tx")
	}

	return ethDetailTXObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EthDetailTX) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no eth_detail_tx provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(ethDetailTXColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ethDetailTXInsertCacheMut.RLock()
	cache, cached := ethDetailTXInsertCache[key]
	ethDetailTXInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ethDetailTXAllColumns,
			ethDetailTXColumnsWithDefault,
			ethDetailTXColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `eth_detail_tx` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `eth_detail_tx` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `eth_detail_tx` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, ethDetailTXPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into eth_detail_tx")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ethDetailTXMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for eth_detail_tx")
	}

CacheNoHooks:
	if !cached {
		ethDetailTXInsertCacheMut.Lock()
		ethDetailTXInsertCache[key] = cache
		ethDetailTXInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the EthDetailTX.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EthDetailTX) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	ethDetailTXUpdateCacheMut.RLock()
	cache, cached := ethDetailTXUpdateCache[key]
	ethDetailTXUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ethDetailTXAllColumns,
			ethDetailTXPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update eth_detail_tx, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `eth_detail_tx` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, ethDetailTXPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, append(wl, ethDetailTXPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update eth_detail_tx row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for eth_detail_tx")
	}

	if !cached {
		ethDetailTXUpdateCacheMut.Lock()
		ethDetailTXUpdateCache[key] = cache
		ethDetailTXUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q ethDetailTXQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for eth_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for eth_detail_tx")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EthDetailTXSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ethDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `eth_detail_tx` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ethDetailTXPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ethDetailTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ethDetailTX")
	}
	return rowsAff, nil
}

var mySQLEthDetailTXUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EthDetailTX) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no eth_detail_tx provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(ethDetailTXColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEthDetailTXUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ethDetailTXUpsertCacheMut.RLock()
	cache, cached := ethDetailTXUpsertCache[key]
	ethDetailTXUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ethDetailTXAllColumns,
			ethDetailTXColumnsWithDefault,
			ethDetailTXColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ethDetailTXAllColumns,
			ethDetailTXPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert eth_detail_tx, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "eth_detail_tx", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `eth_detail_tx` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for eth_detail_tx")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == ethDetailTXMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(ethDetailTXType, ethDetailTXMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for eth_detail_tx")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for eth_detail_tx")
	}

CacheNoHooks:
	if !cached {
		ethDetailTXUpsertCacheMut.Lock()
		ethDetailTXUpsertCache[key] = cache
		ethDetailTXUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single EthDetailTX record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EthDetailTX) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EthDetailTX provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ethDetailTXPrimaryKeyMapping)
	sql := "DELETE FROM `eth_detail_tx` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from eth_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for eth_detail_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ethDetailTXQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ethDetailTXQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from eth_detail_tx")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for eth_detail_tx")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EthDetailTXSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ethDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `eth_detail_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ethDetailTXPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ethDetailTX slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for eth_detail_tx")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EthDetailTX) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEthDetailTX(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EthDetailTXSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EthDetailTXSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ethDetailTXPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `eth_detail_tx`.* FROM `eth_detail_tx` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ethDetailTXPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EthDetailTXSlice")
	}

	*o = slice

	return nil
}

// EthDetailTXExists checks if the EthDetailTX row exists.
func EthDetailTXExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `eth_detail_tx` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if eth_detail_tx exists")
	}

	return exists, nil
}

// InsertAll inserts all rows with the specified column values, using an executor.
func (o EthDetailTXSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	var sql string
	vals := []interface{}{}
	for i, row := range o {
		if !boil.TimestampsAreSkipped(ctx) {
			currTime := time.Now().In(boil.GetLocation())

			if queries.MustTime(row.UpdatedAt).IsZero() {
				queries.SetScanner(&row.UpdatedAt, currTime)
			}
		}

		nzDefaults := queries.NonZeroDefaultSet(ethDetailTXColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			ethDetailTXAllColumns,
			ethDetailTXColumnsWithDefault,
			ethDetailTXColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `eth_detail_tx` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(ethDetailTXType, ethDetailTXMapping, wl)
		if err != nil {
			return err
		}
		value := reflect.Indirect(reflect.ValueOf(row))
		vals = append(vals, queries.ValuesFromMapping(value, valMapping)...)
	}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, vals...)
	}

	_, err := exec.ExecContext(ctx, sql, vals...)
	if err != nil {
		return errors.Wrap(err, "models: unable to insert into eth_detail_tx")
	}

	return nil
}
