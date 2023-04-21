package entity

import "time"

type User struct {
	Id                   string        `json:"id"`
	Name                 interface{}   `json:"name"`
	Username             string        `json:"username"`
	Host                 interface{}   `json:"host"`
	AvatarUrl            string        `json:"avatarUrl"`
	AvatarBlurhash       interface{}   `json:"avatarBlurhash"`
	IsBot                bool          `json:"isBot"`
	IsCat                bool          `json:"isCat"`
	Emojis               Emojis        `json:"emojis"`
	OnlineStatus         string        `json:"onlineStatus"`
	BadgeRoles           []interface{} `json:"badgeRoles"`
	Url                  interface{}   `json:"url"`
	Uri                  interface{}   `json:"uri"`
	MovedToUri           interface{}   `json:"movedToUri"`
	AlsoKnownAs          interface{}   `json:"alsoKnownAs"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            interface{}   `json:"updatedAt"`
	LastFetchedAt        interface{}   `json:"lastFetchedAt"`
	BannerUrl            interface{}   `json:"bannerUrl"`
	BannerBlurhash       interface{}   `json:"bannerBlurhash"`
	IsLocked             bool          `json:"isLocked"`
	IsSilenced           bool          `json:"isSilenced"`
	IsSuspended          bool          `json:"isSuspended"`
	Description          interface{}   `json:"description"`
	Location             interface{}   `json:"location"`
	Birthday             interface{}   `json:"birthday"`
	Lang                 interface{}   `json:"lang"`
	Fields               []interface{} `json:"fields"`
	FollowersCount       int           `json:"followersCount"`
	FollowingCount       int           `json:"followingCount"`
	NotesCount           int           `json:"notesCount"`
	PinnedNoteIds        []interface{} `json:"pinnedNoteIds"`
	PinnedNotes          []interface{} `json:"pinnedNotes"`
	PinnedPageId         interface{}   `json:"pinnedPageId"`
	PinnedPage           interface{}   `json:"pinnedPage"`
	PublicReactions      bool          `json:"publicReactions"`
	FfVisibility         string        `json:"ffVisibility"`
	TwoFactorEnabled     bool          `json:"twoFactorEnabled"`
	UsePasswordLessLogin bool          `json:"usePasswordLessLogin"`
	SecurityKeys         bool          `json:"securityKeys"`
	Roles                []interface{} `json:"roles"`
	Memo                 interface{}   `json:"memo"`
	IsAdmin              bool          `json:"isAdmin"`
	IsModerator          bool          `json:"isModerator"`
}
