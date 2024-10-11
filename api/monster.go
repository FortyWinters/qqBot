package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	SERVER_URL       = "http://localhost:9876/api/monster"
	INFO_GET_API     = "/info/get"
	WEAKNESS_GET_API = "/weakness/get"
	MATERIAL_GET_API = "/material/get"
)

func GetMonsterInfoHandler(argStr string, game_type int) (MonsterResJson, error) {
	if argStr == "" {
		return MonsterResJson{Info: "", Img: ""}, errors.New("需要怪物名称")
	}

	url := SERVER_URL
	argv := strings.Split(argStr, " ")
	payload := MonsterReqJson{
		Name:     argv[0],
		GameType: game_type,
	}

	var apiUrl string
	var errorMessage string

	if len(argv) == 1 {
		apiUrl = INFO_GET_API
		errorMessage = "也没这怪啊"
	} else if len(argv) == 2 {
		switch argv[1] {
		case "弱点":
			apiUrl = WEAKNESS_GET_API
			errorMessage = "弱点查询失败"
		case "素材", "掉落":
			apiUrl = MATERIAL_GET_API
			errorMessage = "素材查询失败"
		default:
			return MonsterResJson{Info: "参数错误", Img: ""}, nil
		}
	}

	url += apiUrl
	monsterInfo, err := SendGetRequest(url, payload)
	if err != nil {
		return MonsterResJson{Info: "", Img: ""}, errors.New("internal error")
	}
	if monsterInfo == "Internal server error" {
		return MonsterResJson{Info: errorMessage, Img: ""}, nil
	}

	var info, img string

	if len(argv) == 1 {
		info, img = monsterInfoResult(monsterInfo)
		return MonsterResJson{Info: info, Img: img}, nil
	} else if len(argv) == 1 {
		switch argv[1] {
		case "弱点":
			info, img = monsterPartResult(monsterInfo)
			return MonsterResJson{Info: info, Img: img}, nil
		case "素材", "掉落":
			info, img = monsterMaterialResult(monsterInfo)
			return MonsterResJson{Info: info, Img: img}, nil
		default:
			return MonsterResJson{Info: info, Img: img}, nil
		}
	}

	return MonsterResJson{Info: "", Img: ""}, nil
}

func monsterInfoResult(monsterInfoStr string) (string, string) {
	var monsterInfo MonsterInfo

	err := json.Unmarshal([]byte(monsterInfoStr), &monsterInfo)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return "", ""
	}

	return fmt.Sprintf("%s\n%s", monsterInfo.MonsterName, monsterInfo.MonsterDescription), monsterInfo.MonsterIconUrl
}

func monsterPartResult(monsterInfoStr string) (string, string) {
	return monsterInfoStr, ""
}

func monsterMaterialResult(monsterInfoStr string) (string, string) {
	return monsterInfoStr, ""
}
