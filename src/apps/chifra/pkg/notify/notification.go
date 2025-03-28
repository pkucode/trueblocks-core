package notify

import "github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"

// TODO: BOGUS - NOTIFY CODE
type Message string

type Notification[T NotificationPayload] struct {
	Msg     Message         `json:"msg"`
	Meta    *types.MetaData `json:"meta"`
	Payload T               `json:"payload"`
}

type NotificationPayload interface {
	[]NotificationPayloadAppearance |
		[]NotificationPayloadChunkWritten |
		NotificationPayloadChunkWritten |
		string
}
