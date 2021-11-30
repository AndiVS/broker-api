package model

import "github.com/google/uuid"

// Profile model of user
type Profile struct {
	ProfileID    *uuid.UUID   `param:"profile_id" query:"profile_id" header:"profile_id" form:"profile_id" bson:"profile_id" msg:"profile_id" json:"profile_id"`
	Balance      *float32     `param:"balance" query:"balance" header:"balance" form:"balance" bson:"balance" msg:"balance" json:"balance"`
	PositionList []*uuid.UUID `param:"position_list" query:"position_list" header:"position_list" form:"position_list" bson:"position_list" msg:"position_list" json:"position_list"`
}
