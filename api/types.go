package api

type MonsterRequestJson struct {
	Name     string `json:"name"`
	GameType int    `json:"game_type"`
}

type MonsterInfo struct {
	ID                 int    `json:"id"`
	MonsterID          int    `json:"monster_id"`
	MonsterName        string `json:"monster_name"`
	MonsterType        int    `json:"monster_type"`
	MonsterAlias       string `json:"monster_alias"`
	MonsterDescription string `json:"monster_description"`
	GameType           int    `json:"game_type"`
}
