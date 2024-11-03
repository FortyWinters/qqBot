package api

type MonsterInfo struct {
	ID                 int    `json:"id"`
	MonsterID          int    `json:"monster_id"`
	MonsterName        string `json:"monster_name"`
	MonsterType        int    `json:"monster_type"`
	MonsterDescription string `json:"monster_description"`
	MonsterIconUrl     string `json:"monster_icon_url"`
	GameType           int    `json:"game_type"`
}

type MonsterInfoRes struct {
	ID                 int    `json:"id"`
	MonsterID          int    `json:"monster_id"`
	MonsterName        string `json:"monster_name"`
	MonsterType        int    `json:"monster_type"`
	MonsterDescription string `json:"monster_description"`
	MonsterIconUrl     string `json:"monster_icon_url"`
	GameType           int    `json:"game_type"`
}

type MonsterWeaknessRes struct {
	MonsterID    int           `json:"monster_id"`
	MonsterName  string        `json:"monster_name"`
	MonsterType  int           `json:"monster_type"`
	GameType     int           `json:"game_type"`
	MonsterParts []MonsterPart `json:"monster_parts"`
}

type MonsterPart struct {
	PartName          string            `json:"part_name"`
	MonsterWeaknesses []MonsterWeakness `json:"monster_weaknesses"`
}

type MonsterWeakness struct {
	WeaknessType  int `json:"weakness_type"`
	WeaknessValue int `json:"weakness_value"`
}

type MonsterInfoReqJson struct {
	Name     string `json:"name"`
	GameType int    `json:"game_type"`
}

type MonsterWeaknessReqJson struct {
	MonsterName string `json:"monster_name"`
	GameType    int    `json:"game_type"`
}

type MonsterReqJson struct {
	Name     string `json:"name"`
	GameType int    `json:"game_type"`
}

type MonsterResJson struct {
	Info string `json:"info"`
	Img  string `json:"img"`
}
