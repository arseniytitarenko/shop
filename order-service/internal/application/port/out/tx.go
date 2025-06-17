package out

import "context"

type Tx interface {
	Exec(ctx context.Context, fn func(ctx context.Context, tx TxRepo) error) error
}

type TxRepo interface {
	OrderRepo() OrderRepo
	OutboxRepo() OutboxRepo
	InboxRepo() InboxRepo
}
