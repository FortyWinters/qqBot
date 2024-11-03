package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"qqbot/localization"
	"sort"
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

	argv := strings.Split(argStr, " ")
	monster_name := argv[0]

	if len(argv) == 1 {
		res, err := getMonsterInfo(monster_name, game_type)
		if err != nil {
			return MonsterResJson{Info: "也没这怪啊", Img: ""}, nil
		}

		return res, nil
	} else if len(argv) == 2 {
		switch argv[1] {
		case "弱点":
			res, err := getMonsterWeakness(monster_name, game_type)
			if err != nil {
				return MonsterResJson{Info: "也没这怪啊", Img: ""}, nil
			}
			return res, nil
		case "素材", "掉落":
			res, err := getMonsterMaterial(monster_name, game_type)
			if err != nil {
				return MonsterResJson{Info: "也没这怪啊", Img: ""}, nil
			}
			return res, nil
		default:
			return MonsterResJson{Info: "参数错误", Img: ""}, nil
		}
	}
	return MonsterResJson{Info: "查询失败QAQ", Img: ""}, nil
}

func getMonsterInfo(name string, game_type int) (MonsterResJson, error) {
	var res string
	var err error
	var monsterInfo MonsterInfoRes

	item := MonsterInfoReqJson{
		Name:     name,
		GameType: game_type,
	}

	res, err = SendGetInfoRequest(SERVER_URL+INFO_GET_API, item)
	if err != nil {
		return MonsterResJson{Info: res, Img: ""}, nil
	}

	err = json.Unmarshal([]byte(res), &monsterInfo)
	if err != nil {
		return MonsterResJson{Info: "我晕了", Img: ""}, nil
	}

	return MonsterResJson{Info: fmt.Sprintf("%s\n%s", monsterInfo.MonsterName, monsterInfo.MonsterDescription), Img: monsterInfo.MonsterIconUrl}, nil
}

func partNameLocalization(partNameEnStr string) (string, string) {
	partNameEnStr = strings.TrimSpace(partNameEnStr)
	if strings.Contains(partNameEnStr, "(") && strings.Contains(partNameEnStr, ")") {
		partName := strings.TrimSpace(partNameEnStr[:strings.Index(partNameEnStr, "(")])
		partStatus := strings.TrimSpace(partNameEnStr[strings.Index(partNameEnStr, "(")+1 : strings.Index(partNameEnStr, ")")])

		partNameCn, nameExists := localization.MonsterPartMap[partName]
		partStatusCn, statusExists := localization.MonsterPartMap[partStatus]

		if nameExists {
			partName = partNameCn
		}
		if statusExists {
			partStatus = partStatusCn
		}
		return partName, fmt.Sprintf("(%s)", partStatus)
	}

	partName, nameExsits := localization.MonsterPartMap[partNameEnStr]
	if nameExsits {
		return partName, ""
	} else {
		return partNameEnStr, ""
	}
}

func getMonsterWeakness(monster_name string, game_type int) (MonsterResJson, error) {
	var res string
	var err error
	var monsterWeaknessRes MonsterWeaknessRes

	item := MonsterWeaknessReqJson{
		MonsterName: monster_name,
		GameType:    game_type,
	}

	res, err = SendGetWeaknessRequest(SERVER_URL+WEAKNESS_GET_API, item)
	if err != nil {
		return MonsterResJson{Info: res, Img: ""}, nil
	}

	err = json.Unmarshal([]byte(res), &monsterWeaknessRes)
	if err != nil {
		return MonsterResJson{Info: "json解析失败", Img: ""}, nil
	}

	var weaknessSlice []string
	weaknessSlice = append(weaknessSlice, monsterWeaknessRes.MonsterName)

	for _, p := range monsterWeaknessRes.MonsterParts {
		var pSlice []string
		if p.PartName == "" {
			break
		}
		partNameCn, partStatusCn := partNameLocalization(p.PartName)
		pSlice = append(pSlice, partNameCn)
		pSlice = append(pSlice, "\t")

		var middleWeaknesses []MonsterWeakness

		for i, w := range p.MonsterWeaknesses {
			weaknessTypeStr := localization.MonsterWeaknessMap[w.WeaknessType]

			if i < 2 {
				pSlice = append(pSlice, fmt.Sprintf("%02d(%s)\t\t", w.WeaknessValue, weaknessTypeStr))
			} else if i == 2 {
				pSlice = append(pSlice, fmt.Sprintf("%02d(%s)", w.WeaknessValue, weaknessTypeStr))
			} else if i >= 3 && i < 8 {
				middleWeaknesses = append(middleWeaknesses, w)
			} else if i >= 8 {
				break
			}
		}

		if partStatusCn == "" {
			pSlice = append(pSlice, "\n\t\t")
		} else {
			pSlice = append(pSlice, fmt.Sprintf("\n%s\t", partStatusCn))
		}

		sort.Slice(middleWeaknesses, func(i, j int) bool {
			return middleWeaknesses[i].WeaknessValue > middleWeaknesses[j].WeaknessValue
		})

		for i, v := range middleWeaknesses[:3] {
			weaknessTypeStr := localization.MonsterWeaknessMap[v.WeaknessType]
			if i == 2 {
				pSlice = append(pSlice, fmt.Sprintf("%02d(%s)", v.WeaknessValue, weaknessTypeStr))
			} else {
				pSlice = append(pSlice, fmt.Sprintf("%02d(%s)\t\t", v.WeaknessValue, weaknessTypeStr))
			}
		}

		pStr := strings.Join(pSlice, "")
		weaknessSlice = append(weaknessSlice, pStr)
	}

	weaknessStr := strings.Join(weaknessSlice, "\n")

	return MonsterResJson{Info: weaknessStr, Img: ""}, nil
}

func getMonsterMaterial(monster_name string, game_type int) (MonsterResJson, error) {
	fmt.Println("{} {}", monster_name, game_type)
	return MonsterResJson{Info: "功能尚在开发", Img: ""}, nil
}
