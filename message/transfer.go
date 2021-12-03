package message

import (
	"context"
	"fmt"

	"github.com/fox-one/mixin-sdk-go"
)

func HandleTransfer(ctx context.Context, msg *mixin.MessageView, transfer *mixin.TransferView) error {
	fmt.Println(transfer)
	return nil
}
