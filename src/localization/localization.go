package localization

var MonsterPartMap map[string]string
var MonsterWeaknessMap map[int]string

func init() {
	MonsterPartMap = make(map[string]string)
	MonsterWeaknessMap = make(map[int]string)

	MonsterPartMap["Snout"] = "鼻头"
	MonsterPartMap["Head"] = "龙头"
	MonsterPartMap["Neck"] = "脖颈"
	MonsterPartMap["Chest"] = "胸部"
	MonsterPartMap["Legs"] = "龙腿"
	MonsterPartMap["Arms"] = "龙手"
	MonsterPartMap["Forefeet"] = "前足"
	MonsterPartMap["Forelegs"] = "前腿"
	MonsterPartMap["Body"] = "胴体"
	MonsterPartMap["Wings"] = "全翅"
	MonsterPartMap["Wing Webbings"] = "翼膜"
	MonsterPartMap["Wing Tips"] = "翅尖"
	MonsterPartMap["Back"] = "脊背"
	MonsterPartMap["Hindfeet"] = "后足"
	MonsterPartMap["Hindlegs"] = "后腿"
	MonsterPartMap["Tail"] = "尻尾"
	MonsterPartMap["Tail Tip"] = "尾尖"
	MonsterPartMap["Forearms"] = "前臂"
	MonsterPartMap["Horns"] = "龙角"
	MonsterPartMap["Horn"] = "龙角"
	MonsterPartMap["Stomach"] = "肚子"
	MonsterPartMap["Broken"] = "破"
	MonsterPartMap["Critical"] = "界"
	MonsterPartMap["Dragon"] = "龙"
	MonsterPartMap["Ice"] = "冰"
	MonsterPartMap["Fire"] = "火"
	MonsterPartMap["No Element"] = "无"
	MonsterPartMap["White"] = "白"
	MonsterPartMap["Gloss White"] = "白"
	MonsterPartMap["Black"] = "黑"
	MonsterPartMap["Gloss Black"] = "黑"
	MonsterPartMap["Wounded"] = "伤"

	MonsterWeaknessMap[0] = "斩"
	MonsterWeaknessMap[1] = "击"
	MonsterWeaknessMap[2] = "弹"
	MonsterWeaknessMap[3] = "火"
	MonsterWeaknessMap[4] = "水"
	MonsterWeaknessMap[5] = "雷"
	MonsterWeaknessMap[6] = "冰"
	MonsterWeaknessMap[7] = "龙"
	MonsterWeaknessMap[8] = "异常"
	MonsterWeaknessMap[9] = "耐力"
}
