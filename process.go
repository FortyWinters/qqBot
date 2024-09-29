package main

import (
	"context"
	"fmt"
	"qqbot/api"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/openapi"
)

type Processor struct {
	api openapi.OpenAPI
}

func (p Processor) ProcessMessage(input string, data *dto.WSATMessageData) error {
	ctx := context.Background()
	cmd := message.ParseCommand(input)
	toCreate := &dto.MessageToCreate{
		Content: "",
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}

	var info string
	var err error

	switch cmd.Cmd {
	case "/怪物":
		info, err = api.GetMonsterInfo(cmd.Content)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/武器":
		info, err = api.GetWeaponInfo(cmd.Content)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/装备":
		info, err = api.GetEquipmentInfo(cmd.Content)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	default:
		toCreate.Content = "指令不存在"
	}
	p.sendReply(ctx, data.ChannelID, toCreate)

	return nil
}
