package worker

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"order/internal/application/port/out"
	"order/internal/domain"
	"time"
)

type inboxFullPayload struct {
	ID      uuid.UUID     `json:"id"`
	OrderID uuid.UUID     `json:"order_id"`
	Status  domain.Status `json:"status"`
}

func RunProcessingOrderWorker(ctx context.Context, txManager out.Tx, interval time.Duration, msgTypeIn string) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var messages []domain.Inbox
			err := txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
				var err error
				messages, err = tx.InboxRepo().GetUnprocessed(ctx, 10, msgTypeIn)
				return err
			})

			if err != nil {
				log.Println("failed to fetch unprocessed messages")
				continue
			}

			for _, msg := range messages {
				_ = txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
					var payload inboxFullPayload
					if err = json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
						log.Printf("failed to parse inbox body: %v\n", err)
						return err
					}

					if !domain.IsValidStatus(payload.Status) {
						log.Printf("invalid status: %s in order: %s\n", payload.Status, payload.OrderID)
						return errors.New("invalid status")
					}

					order, err := tx.OrderRepo().GetOrder(ctx, payload.OrderID)
					if err != nil {
						log.Printf("failed to fetch order %s: %v\n", payload.OrderID, err)
						return err
					}

					order.Status = payload.Status

					err = tx.OrderRepo().SaveOrder(ctx, order)
					if err != nil {
						log.Printf("failed to save order: %s\n", payload.OrderID)
						return err
					}

					err = tx.InboxRepo().MarkProcessed(ctx, msg.ID)
					if err != nil {
						log.Printf("failed to mark processed inbox for order: %s\n", payload.OrderID)
						return err
					}

					return nil
				})
			}
		}
	}
}
