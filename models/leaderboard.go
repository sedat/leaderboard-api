package models

type Leaderboard struct {
	DisplayName string `bson:"display_name,omitempty" json:"display_name,omitempty"`
	Country     string `bson:"country,omitempty" json:"country,omitempty"`
	Points      int    `bson:"points,omitempty" json:"points,omitempty"`
	Rank        int    `bson:"rank,omitempty" json:"rank,omitempty"`
}
