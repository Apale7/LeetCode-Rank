package model

type Rank struct {
	Name        string `json:"name"`
	Easy        int    `json:"easy"`
	Medium      int    `json:"medium"`
	Hard        int    `json:"hard"`
	TotalAC     int    `json:"total_ac"`
	TotalAC7Day int    `json:"total_ac_7_day"`
}
