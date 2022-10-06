package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StockTaskModel = (*customStockTaskModel)(nil)

type (
	// StockTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStockTaskModel.
	StockTaskModel interface {
		stockTaskModel
		InsertWithTx(ctx context.Context, tx *sql.Tx, data *StockTask) (sql.Result, error)
	}

	customStockTaskModel struct {
		*defaultStockTaskModel
	}
)

// NewStockTaskModel returns a model for the database table.
func NewStockTaskModel(conn sqlx.SqlConn) StockTaskModel {
	return &customStockTaskModel{
		defaultStockTaskModel: newStockTaskModel(conn),
	}
}

func (m *defaultStockTaskModel) InsertWithTx(ctx context.Context, tx *sql.Tx, data *StockTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, stockTaskRowsExpectAutoSet)
	stmt, err := m.conn.PrepareCtx(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			logx.Error(err)
		}
	}()
	return stmt.ExecCtx(ctx, data.StockId, data.Amount)

	//ret, err := m.conn.ExecCtx(ctx, query, data.StockId, data.Amount)
	//return ret, err
}
