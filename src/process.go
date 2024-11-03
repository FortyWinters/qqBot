package main

import (
	"context"
	"fmt"
	"log"
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

func (p Processor) sendReply(ctx context.Context, channelID string, toCreate *dto.MessageToCreate) {
	if _, err := p.api.PostMessage(ctx, channelID, toCreate); err != nil {
		log.Println(err)
	}
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
		Image: "",
	}

	var res api.MonsterResJson
	var err error
	var info string

	switch cmd.Cmd {
	case "/世界怪物":
		res, err = api.GetMonsterInfoHandler(cmd.Content, WORLD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = res.Info
			toCreate.Image = res.Img
		}
	case "/世界装备":
		info, err = api.GetEquipmentInfoHandler(cmd.Content, WORLD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/崛起怪物":
		res, err = api.GetMonsterInfoHandler(cmd.Content, RISE)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = res.Info
			toCreate.Image = res.Img
		}
	case "/崛起装备":
		info, err = api.GetEquipmentInfoHandler(cmd.Content, RISE)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	case "/荒野怪物":
		res, err = api.GetMonsterInfoHandler(cmd.Content, WILD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = res.Info
			toCreate.Image = res.Img
		}
	case "/荒野装备":
		info, err = api.GetEquipmentInfoHandler(cmd.Content, WILD)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = info
		}
	default:
		gptRes, err := api.GPTHandler(cmd.Cmd)
		if err != nil {
			toCreate.Content = fmt.Sprintf("参数错误: %v", err)
		} else {
			toCreate.Content = gptRes
		}
	}

	p.sendReply(ctx, data.ChannelID, toCreate)

	return nil
}
