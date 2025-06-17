package worker

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"order/internal/application/port/out"
	"order/internal/domain"
	"time"
)

type inboxPayload struct {
	Id uuid.UUID `json:"id"`
}

func RunInboxWorker(ctx context.Context, txManager out.Tx, subscriber out.BrokerSubscriber) {
	for {
		select {
		case <-ctx.Done():
			log.Println("inbox worker stopped")
			return
		case msg := <-subscriber.Messages():
			var payload inboxPayload
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("failed to parse message body: %v\n", err)
				continue
			}
			log.Printf("received type: %s\n", msg.Type)
			log.Printf("msg: %+v", msg)
			err := txManager.Exec(ctx, func(ctx context.Context, tx out.TxRepo) error {
				inbox := &domain.Inbox{
					ID:        payload.Id,
					Type:      msg.Type,
					Payload:   string(msg.Body),
					CreatedAt: time.Now(),
				}
				return tx.InboxRepo().NewInbox(ctx, inbox)
			})

			if err != nil {
				log.Printf("failed to save inbox message: %v\n", err)
				continue
			}
		}
	}
}
