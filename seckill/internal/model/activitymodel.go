package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActivityModel = (*customActivityModel)(nil)

type (
	// ActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityModel.
	ActivityModel interface {
		activityModel
		FindLimitOffset(ctx context.Context, limit, offset int) ([]*Activity, error)
	}

	customActivityModel struct {
		*defaultActivityModel
	}
)

// NewActivityModel returns a model for the database table.
func NewActivityModel(conn sqlx.SqlConn) ActivityModel {
	return &customActivityModel{
		defaultActivityModel: newActivityModel(conn),
	}
}

func (m *defaultActivityModel) FindLimitOffset(ctx context.Context, limit, offset int) ([]*Activity, error) {
	query := fmt.Sprintf("select %s from %s limit ? offset ?", activityRows, m.table)
	var resp []*Activity
	err := m.conn.QueryRowsCtx(ctx, &resp, query, limit, offset)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
