package worker

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"payment/internal/application/errs"
	"payment/internal/application/port/out"
	"payment/internal/domain"
	"time"
)

type inboxFullPayload struct {
	ID      uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id"`
	UserID  uuid.UUID `json:"user_id"`
	Amount  uint      `json:"amount"`
}

func RunProcessingOrderWorker(ctx context.Context, txManager out.Tx, interval time.Duration,
	msgTypeIn, msgTypeOut string) {
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
					inboxStatus := domain.StatusFinished

					account, err := tx.AccountRepo().GetAccount(ctx, payload.UserID)

					if err != nil {
						if errors.Is(err, errs.ErrAccountNotFound) {
							inboxStatus = domain.StatusCancelled
						} else {
							log.Printf("failed to fetch account for user %s: %v\n", payload.UserID, err)
							return err
						}
					} else {
						if account.Balance < payload.Amount {
							inboxStatus = domain.StatusCancelled
						} else {
							account.Balance -= payload.Amount
						}
					}

					err = tx.AccountRepo().SaveAccount(ctx, account)
					if err != nil {
						log.Printf("failed to save account for user %s: %v\n", payload.UserID, err)
						return err
					}

					err = tx.InboxRepo().MarkProcessed(ctx, msg.ID)
					if err != nil {
						log.Printf("failed to mark processed inbox for user %s: %v\n", payload.UserID, err)
						return err
					}

					payloadBytes, _ := json.Marshal(map[string]interface{}{
						"id":       msg.ID,
						"order_id": payload.OrderID,
						"status":   inboxStatus,
					})

					outbox := domain.Outbox{
						ID:        msg.ID,
						Type:      msgTypeOut,
						Payload:   string(payloadBytes),
						CreatedAt: time.Now(),
					}
					return tx.OutboxRepo().NewOutbox(ctx, &outbox)
				})
			}
		}
	}
}
