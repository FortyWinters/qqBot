package main

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"
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

func main() {
	ctx := context.Background()
	botToken := token.New(token.TypeBot)
	if err := botToken.LoadFromConfig(getConfigPath("config.yaml")); err != nil {
		log.Fatalln(err)
	}

	api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	// api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)

	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln(err)
	}

	processor = Processor{api: api}

	intent := websocket.RegisterHandlers(
		ATMessageEventHandler(),
		// ReadyHandler(),
		// ErrorNotifyHandler(),
		// GuildEventHandler(),
		// MemberEventHandler(),
		// ChannelEventHandler(),
		// DirectMessageHandler(),
		// CreateMessageHandler(),
	)

	if err = botgo.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}

// 实现处理 @ 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		return processor.ProcessMessage(input, data)
	}
}

// // 感知连接成功事件
// func ReadyHandler() event.ReadyHandler {
// 	return func(event *dto.WSPayload, data *dto.WSReadyData) {
// 		log.Println("ready event receive: ", data)
// 	}
// }

// // 连接关闭回调
// func ErrorNotifyHandler() event.ErrorNotifyHandler {
// 	return func(err error) {
// 		log.Println("error notify receive: ", err)
// 	}
// }

// // 处理频道事件
// func GuildEventHandler() event.GuildEventHandler {
// 	return func(event *dto.WSPayload, data *dto.WSGuildData) error {
// 		fmt.Println(data)
// 		return nil
// 	}
// }

// // 处理成员变更事件
// func MemberEventHandler() event.GuildMemberEventHandler {
// 	return func(event *dto.WSPayload, data *dto.WSGuildMemberData) error {
// 		fmt.Println(data)
// 		return nil
// 	}
// }

// // 处理子频道事件
// func ChannelEventHandler() event.ChannelEventHandler {
// 	return func(event *dto.WSPayload, data *dto.WSChannelData) error {
// 		fmt.Println(data)
// 		return nil
// 	}
// }

// // 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
// func DirectMessageHandler() event.DirectMessageEventHandler {
// 	return func(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
// 		fmt.Println(data)
// 		return nil
// 	}
// }

// // 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
// func CreateMessageHandler() event.MessageEventHandler {
// 	return func(event *dto.WSPayload, data *dto.WSMessageData) error {
// 		fmt.Println(data)
// 		return nil
// 	}
// }

func getConfigPath(name string) string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return fmt.Sprintf("%s/%s", path.Dir(filename), name)
	}
	return ""
}
