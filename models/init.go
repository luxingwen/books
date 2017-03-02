package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	AppConfig config.Configer
	SqlServer string
	Dev       bool
	engine    *xorm.Engine
)

func init() {
	AppConfig = beego.AppConfig
	SqlServer = AppConfig.DefaultString("mysql::server", "")
	Dev = AppConfig.DefaultBool("dev", true)
	var err error
	engine, err = GetEngine()
	if err != nil {
		log.Fatal(err)
	}
}

func SyncModels() {
	engine, err := GetEngine()
	if err != nil {
		log.Fatal("get engine err:", err)
	}
	defer engine.Close()
	err = engine.Sync(new(User), new(Book), new(MyBook), new(Category), new(ShopCar), new(Community))
	if err != nil {
		log.Fatal(err)
	}
}

func insert(obj interface{}) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	_, err = engine.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

func update(obj interface{}, id int) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	affected, err := engine.Where("id=?", id).Update(obj)
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("no update.")
	}
	return nil
}

func delete(obj interface{}, id int) error {
	engine, err := GetEngine()
	if err != nil {
		return err
	}
	defer engine.Close()
	_, err = engine.Where("id=?", id).Unscoped().Delete(obj)
	if err != nil {
		return err
	}
	return nil
}

func GetEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", SqlServer)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(true)
	return engine, nil
}
