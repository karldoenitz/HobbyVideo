package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
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
