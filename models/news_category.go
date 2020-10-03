package models

type NewsCategory struct {
	Id           int    `json:"id" xorm:"not null pk autoincr index(module) SMALLINT(5)"`
	Tid          int    `json:"tid" xorm:"not null comment('栏目类型，0单页，1模块，2外链') index TINYINT(1)"`
	Pid          int    `json:"pid" xorm:"not null default 0 comment('上级id') index(module) SMALLINT(5)"`
	Mid          string `json:"mid" xorm:"not null comment('模块目录') index VARCHAR(20)"`
	Name         string `json:"name" xorm:"not null comment('栏目名称') VARCHAR(30)"`
	Identifier   string `json:"identifier" xorm:"not null default '' comment('地址') index VARCHAR(30)"`
	Path         string `json:"path" xorm:"not null comment('下级所有id') TEXT"`
	Thumb        string `json:"thumb" xorm:"not null comment('栏目图片') VARCHAR(255)"`
	Show         int    `json:"show" xorm:"not null default 1 comment('是否显示') index TINYINT(1)"`
	Content      string `json:"content" xorm:"not null comment('单页内容') MEDIUMTEXT"`
	Setting      string `json:"setting" xorm:"not null comment('属性配置') TEXT"`
	Displayorder int    `json:"displayorder" xorm:"not null default 0 index(module) SMALLINT(5)"`
}
