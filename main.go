package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

var processor Processor

// 实现处理 @ 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		return processor.ProcessMessage(input, data)
	}
}

func sendMessageUntilTargetDay(ctx context.Context, processor Processor) {
	targetDate := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)

	for {
		now := time.Now()
		messageTime := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())

		if now.After(messageTime) {
			messageTime = messageTime.Add(24 * time.Hour)
		}

		duration := time.Until(messageTime)

		time.Sleep(duration)

		now = time.Now()
		daysUntil := targetDate.Sub(now).Hours() / 24

		msg := fmt.Sprintf("距离荒野发售还有%.0f年\n", daysUntil)
		processor.sendReply(ctx, "665030445", &dto.MessageToCreate{
			Content: msg,
		})
	}
}

func main() {
	ctx := context.Background()
	botToken := token.New(token.TypeBot)
	if err := botToken.LoadFromConfig("config.yaml"); err != nil {
		log.Fatalln(err)
	}

	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)

	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln(err)
	}

	processor = Processor{api: api}

	go sendMessageUntilTargetDay(ctx, processor)

	intent := websocket.RegisterHandlers(
		ATMessageEventHandler(),
	)

	if err = botgo.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}
