package main

import (
	"context"
	"fmt"
	"qqbot/api"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/openapi"
)

const (
	WORLD = 0
	RISE  = 1
	WILD  = 2
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
	case "/世界怪物":
		info, err = api.GetMonsterInfo(cmd.Content, WORLD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/世界装备":
		info, err = api.GetEquipmentInfo(cmd.Content, WORLD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/崛起怪物":
		info, err = api.GetMonsterInfo(cmd.Content, RISE)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/崛起装备":
		info, err = api.GetEquipmentInfo(cmd.Content, RISE)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/荒野怪物":
		info, err = api.GetMonsterInfo(cmd.Content, WILD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/荒野装备":
		info, err = api.GetEquipmentInfo(cmd.Content, WILD)
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
