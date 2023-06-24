package response

import "time"

type GetRoomUsers struct {
	Users []*RoomUser `json:"users"`
}

type RoomUser struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Profile    string    `json:"profile"`
	PictureURL string    `json:"pictureUrl"`
	LastReadAt time.Time `json:"lastReadAt"`
}
