package worker

import (
	"context"
	"log"
	"payment/internal/application/port/out"
	"payment/internal/domain"
	"time"
)

func RunOutboxWorker(ctx context.Context, txManager out.Tx, publisher out.BrokerPublisher, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var messages []domain.Outbox
			err := txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
				var err error
				messages, err = tx.OutboxRepo().GetUnprocessed(ctx, 10)
				return err
			})

			if err != nil {
				log.Println("failed to fetch unprocessed messages")
				continue
			}

			for _, msg := range messages {
				_ = txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
					err = publisher.Publish(ctx, msg.Type, msg.Payload)
					if err != nil {
						log.Printf("failed to publish outbox message %s: %v\n", msg.ID, err)
						return err
					}
					err = tx.OutboxRepo().MarkProcessed(ctx, msg.ID)
					if err != nil {
						log.Printf("failed to mark outbox processed message %s: %v\n", msg.ID, err)
						return err
					}
					return nil
				})
			}
		}
	}
}
