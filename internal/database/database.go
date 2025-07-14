package database

import (
	"os"

	"github.com/lin-snow/ech0/internal/config"
	commonModel "github.com/lin-snow/ech0/internal/model/common"
	connectModel "github.com/lin-snow/ech0/internal/model/connect"
	echoModel "github.com/lin-snow/ech0/internal/model/echo"
	todoModel "github.com/lin-snow/ech0/internal/model/todo"
	userModel "github.com/lin-snow/ech0/internal/model/user"

	util "github.com/lin-snow/ech0/internal/util/err"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql" // 使用 mysql 驱动

	// "github.com/glebarez/sqlite" // 使用 glebarez/sqlite 作为 SQLite 驱动
	"gorm.io/gorm"
)

// DB 全局数据库连接变量
var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	// 读取数据库类型和保存路径
	dbType := config.Config.Database.Type
	dbPath := config.Config.Database.Path

	dir := dbPath[:len(dbPath)-len("/ech0.db")] // 提取目录部分
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		util.HandlePanicError(&commonModel.ServerError{
			Msg: commonModel.CREATE_DB_PATH_PANIC,
			Err: err,
		})
	}

	if dbType == "sqlite" {
		// 添加 PRAGMA 参数，例如 WAL 模式和外键支持
		// pragma := config.Config.Database.Pragma // 从配置读取
		// dsn := dbPath + "?" + pragma
		var err error
		DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			util.HandlePanicError(&commonModel.ServerError{
				Msg: commonModel.INIT_DATABASE_PANIC,
				Err: err,
			})
		}
	}

	if dbType == "mysql" {
		dsn := dbPath // MySQL DSN 格式为 "username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			util.HandlePanicError(&commonModel.ServerError{
				Msg: commonModel.INIT_DATABASE_PANIC,
				Err: err,
			})
		}
	}

	if err := MigrateDB(); err != nil {
		util.HandlePanicError(&commonModel.ServerError{
			Msg: commonModel.MIGRATE_DB_PANIC,
			Err: err,
		})
	}

	// 从 1.x 迁移到 2.x
	UpdateMigration()
}

// MigrateDB 执行数据库迁移
func MigrateDB() error {
	models := []interface{}{
		&userModel.User{},
		&echoModel.Echo{},
		&echoModel.Image{},
		&commonModel.KeyValue{},
		&todoModel.Todo{},
		&connectModel.Connected{},
	}

	return DB.AutoMigrate(
		models...,
	)
}
