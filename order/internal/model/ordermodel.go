package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		InsertTx(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error)
		UpdateTx(ctx context.Context, tx *sql.Tx, data *Order) error
		//FindLastOrder(ctx context.Context, activityId, goodsId, uid int64) (*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn),
	}
}

//func (m *defaultOrderModel) FindLastOrder(ctx context.Context, activityId, goodsId, uid int64) (*Order, error) {
//	query := fmt.Sprintf("select * from %s where activity_id=? and uid=? and goods_id=? and stock_id=?", m.table)
//	o := &Order{}
//	err := m.conn.QueryRowCtx(ctx, o, query, activityId, uid, goodsId, stockId)
//	if err != nil {
//		return nil, err
//	}
//	return o, nil
//}

func (m *defaultOrderModel) UpdateTx(ctx context.Context, tx *sql.Tx, data *Order) error {
	//query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
	// NOTE: 需要一个更合理的方式找到要取消的订单，这里仅做演示
	query := fmt.Sprintf("update `order` set status=-1 where uid=? and activity_id=? and goods_id=? and stockId=?")
	_, err := tx.ExecContext(ctx, query, data.Uid, data.ActivityId, data.GoodsId, data.StockId)
	return err
}

func (m *defaultOrderModel) InsertTx(ctx context.Context, tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	return tx.ExecContext(ctx, query, data.Uid, data.ActivityId, data.GoodsId, data.StockId, data.Count, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
}
