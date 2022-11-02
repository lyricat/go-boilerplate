package messenger

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"go-boilerplate/message"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pkg/logger"
	"github.com/gofrs/uuid"
)

type (
	Worker struct {
		client *mixin.Client
	}
)

func New(client *mixin.Client) *Worker {
	return &Worker{client: client}
}

func (w *Worker) Run(ctx context.Context) error {
	return w.messageLoop(ctx)
}

func (w *Worker) messageLoop(ctx context.Context) error {
	log := logger.FromContext(ctx).WithField("worker", "message")
	ctx = logger.WithContext(ctx, log)

	h := func(ctx context.Context, msg *mixin.MessageView, userID string) error {
		// if there is no valid user id in the message, drop it
		if userID, _ := uuid.FromString(msg.UserID); userID == uuid.Nil {
			return nil
		}

		if msg.Category == mixin.MessageCategorySystemAccountSnapshot {
			// Decode the transfer view from message's content
			data, err := base64.StdEncoding.DecodeString(msg.Data)
			if err != nil {
				log.Println("failed to decode message content:", err)
				return err
			}
			var view mixin.TransferView
			err = json.Unmarshal(data, &view)
			if err != nil {
				log.Println("failed to decode transfer view:", err)
				return err
			}
			return message.HandleTransfer(ctx, msg, &view)

		} else if msg.Category == mixin.MessageCategoryPlainText {
			return message.HandleTextMessage(ctx, msg)
		}

		return nil
	}

	// Start the message loop.
	for {
		// Pass the callback function into the `BlazeListenFunc`
		if err := w.client.LoopBlaze(ctx, mixin.BlazeListenFunc(h)); err != nil {
			log.Printf("LoopBlaze: %v", err)
		}

		// Sleep for a while
		time.Sleep(time.Second)
	}
}
