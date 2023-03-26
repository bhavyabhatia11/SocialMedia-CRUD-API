package models

type User struct {
	UID   int    `json:"uid" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Users struct {
	Users []User `json:"users"`
}

type Comment struct {
	UID     int    `json:"uid" gorm:"primaryKey"`
	UserUID int    `json:"user_uid"`
	User    User   `json:"user" gorm:"foreignKey:UserUID"`
	PostUID int    `json:"post_uid"`
	Post    Post   `json:"post" gorm:"foreignKey:PostUID"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Post struct {
	UID      int    `json:"uid" gorm:"primaryKey"`
	UserUID  int    `json:"user_uid"`
	User     User   `json:"user" gorm:"foreignKey:UserUID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Likes    uint   `json:"likes"`
	Dislikes uint   `json:"dislikes"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}
