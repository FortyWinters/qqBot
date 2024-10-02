package api

import (
	"errors"
)

func GetMonsterInfo(name string, game int) (string, error) {
	if name == "" {
		return "", errors.New("需要怪物名称")
	}

	url := "http://localhost:9876/api/monster/info/get"
	payload := MonsterRequestJson{
		MonsterName: name,
		Game:        game,
	}
	monsterInfo, err := SendGetRequest(url, payload)
	if err != nil {
		return "", errors.New("internal error")
	}
	if monsterInfo == "Internal server error" {
		monsterInfo = "查询失败"
	}

	return monsterInfo, nil
}
