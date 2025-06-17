package out

import "context"

type Tx interface {
	Exec(ctx context.Context, fn func(ctx context.Context, tx TxRepo) error) error
}

type TxRepo interface {
	AccountRepo() AccountRepo
	InboxRepo() InboxRepo
	OutboxRepo() OutboxRepo
}
