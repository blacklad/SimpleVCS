package util

import "github.com/jinzhu/gorm"

//Config is the config struct
type Config struct {
	Name  string
	Value string
}

//File is the file struct
type File struct {
	Hash    string
	Content string
}

//Tree is the tree struct
type Tree struct {
	Hash  string
	Files string
}

//Commit is the commit struct
type Commit struct {
	Hash    string
	Tree    string
	Parent  string
	Author  string
	Time    string
	Message string
}

//Branch is the branch struct
type Branch struct {
	Name   string
	Commit string
}

//Tag is the tag struct
type Tag struct {
	Name   string
	Commit string
}

//Stash is the stash struct
type Stash struct {
	Name  string
	Files string
}

//DB is the database var
var DB *gorm.DB

//InitDB initializes the database
func InitDB() error {
	var err error
	DB, err = gorm.Open("sqlite3", "svcs.db")
	if err != nil {
		return err
	}
	DB.AutoMigrate(&Config{}, &File{}, &Tree{}, &Commit{}, &Branch{}, &Tag{}, &Stash{})
	return nil
}
