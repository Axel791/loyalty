package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

type UnitOfWorkHandler struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWorkHandler {
	return &UnitOfWorkHandler{db: db}
}

// Do запускает транзакцию (BEGIN), вызывает переданную функцию fn,
// при ошибке откатывает (ROLLBACK), при успехе — коммитит (COMMIT).
func (u *UnitOfWorkHandler) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	// Начинаем транзакцию
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("can't begin tx: %w", err)
	}

	// Создаём "транзакционный контекст", в который кладём саму транзакцию
	txCtx := context.WithValue(ctx, txCtxKey, tx)

	if err := fn(txCtx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx rollback error (%v), original error: %w", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("can't commit tx: %w", err)
	}
	return nil
}

type txCtxKeyType string

const txCtxKey txCtxKeyType = "txCtxKey"
