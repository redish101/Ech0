package model

// UserStatus 用于存储用户状态信息
type UserStatus struct {
	UserID   uint   `json:"user_id"`  // 用户ID
	UserName string `json:"username"` // 用户名
	IsAdmin  bool   `json:"is_admin"` // 是否是管理员
}

// Status 用于存储Echo状态信息
type Status struct {
	SysAdminID uint         `json:"sys_admin_id"` // 系统管理员ID
	Username   string       `json:"username"`     // 系统管理员用户名
	Users      []UserStatus `json:"users"`        // 所有用户
	Logo       string       `json:"logo"`         // 站点logo
	TotalEchos int          `json:"total_echos"`  // 总共发布数量
}

// Heatmap 用于存储热力图数据
type Heatmap struct {
	Date  string `json:"date"`  // 日期
	Count int    `json:"count"` // Echo数量
}

// File 相关
type UploadFileType string
type FileStorageType string

const (
	// ImageType  图片类型
	ImageType UploadFileType = "image"
	// AudioType  音频类型
	AudioType UploadFileType = "audio"
)

const (
	// LOCAL_FILE 本地存储类型
	LOCAL_FILE FileStorageType = "local"
	// S3_FILE   S3 存储类型
	S3_FILE FileStorageType = "s3"
	// R2_FILE   R2 存储类型
	R2_FILE FileStorageType = "r2"
)

// key value表
type KeyValue struct {
	Key   string `json:"key" gorm:"primaryKey"`
	Value string `json:"value"`
}

// 键值对相关
const (
	// SystemSettingsKey 是系统设置的键
	SystemSettingsKey = "system_settings"
	// MigrationKey 是数据库迁移的标记键
	MigrationKey = "db_migration:message_to_echo:v1"
)

// PageQueryResult 用于分页查询的结果数据传输对象
type PageQueryResult[T any] struct {
	Total int64 `json:"total"`
	Items T     `json:"items"`
}

const (
	// InitInstallCode 是初始化安装的标志
	InitInstallCode = 666
)

const (
	// Version 是当前版本号
	Version = "2.2.8"
)

const (
	// GreetingBanner 是控制台横幅
	GreetingBanner = `
███████╗     ██████╗    ██╗  ██╗     ██████╗ 
██╔════╝    ██╔════╝    ██║  ██║    ██╔═████╗
█████╗      ██║         ███████║    ██║██╔██║
██╔══╝      ██║         ██╔══██║    ████╔╝██║
███████╗    ╚██████╗    ██║  ██║    ╚██████╔╝
╚══════╝     ╚═════╝    ╚═╝  ╚═╝     ╚═════╝ 
Redish101 Patched Edition                                           
`
)
