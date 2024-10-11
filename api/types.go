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

type MonsterReqJson struct {
	Name     string `json:"name"`
	GameType int    `json:"game_type"`
}

type MonsterResJson struct {
    Info string `json:"info"`
    Img  string `json:"img"`
}
