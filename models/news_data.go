package models

type NewsData struct {
	Id      int    `json:"id" xorm:"not null pk unique INT(10)"`
	Uid     int    `json:"uid" xorm:"not null comment('作者uid') index MEDIUMINT(8)"`
	Catid   int    `json:"catid" xorm:"not null comment('栏目id') index SMALLINT(5)"`
	Content string `json:"content" xorm:"comment('内容') MEDIUMTEXT"`
}
