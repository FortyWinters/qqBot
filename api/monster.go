package api

import (
	"errors"
)

func GetMonsterInfo(name string) (string, error) {
	if name == "" {
		return "", errors.New("需要怪物名称")
	}

	url := "http://localhost:9876/api/monster/info"
	payload := MonsterRequestJson{
		MonsterName: name,
	}
	monsterInfo, err := SendGetRequest(url, payload)
	if err != nil {
		return "", errors.New("internal error")
	}

	return monsterInfo, nil
}
