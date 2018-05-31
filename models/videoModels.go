package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
	"fmt"
)

type VideoSource struct {
	Id int
	Video_Name string
	Tag string
	Url string
	Download_Url string
	Summary string
	Screen_Shot_Pic string
}

func (videoSource *VideoSource) TableName() string {
	return "video_source"
}

func (videoSource *VideoSource) TableIndex() [][]string {
    return [][]string{
        []string{"Video_Name"},
    }
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "sqlite3", "db.sqlite3")

    // register model
    orm.RegisterModel(new(VideoSource))

    // create table
    orm.RunSyncdb("default", false, true)
}

func GetVideoByName(name string) (lists []orm.ParamsList) {
	o := orm.NewOrm()
	num, err := o.QueryTable("video_source").Filter("Video_Name__contains", name).ValuesList(&lists)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		return lists
	}
	return lists
}

func GetVideosInFront() (lists []orm.ParamsList) {
	o := orm.NewOrm()
	num, err := o.QueryTable("video_source").Limit(9).ValuesList(&lists)
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		return lists
	}
	return lists
}
