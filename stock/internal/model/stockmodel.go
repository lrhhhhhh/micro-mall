package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StockModel = (*customStockModel)(nil)

type (
	// StockModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStockModel.
	StockModel interface {
		stockModel
		DecrOne(ctx context.Context, Id int) (sql.Result, error)
	}

	customStockModel struct {
		*defaultStockModel
	}
)

// NewStockModel returns a model for the database table.
func NewStockModel(conn sqlx.SqlConn) StockModel {
	return &customStockModel{
		defaultStockModel: newStockModel(conn),
	}
}

func (m *customStockModel) DecrOne(ctx context.Context, Id int) (sql.Result, error) {
	query := fmt.Sprintf("update %s set `count` = `count`-1 where id=?", m.tableName())
	return m.conn.ExecCtx(ctx, query, Id)
}
