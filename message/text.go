package message

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/fox-one/mixin-sdk-go"
)

func HandleTextMessage(ctx context.Context, msg *mixin.MessageView) error {
	// Decode the message content
	content, err := base64.StdEncoding.DecodeString(msg.Data)
	if err != nil {
		return nil
	}

	fmt.Println(string(content))
	return nil
}
