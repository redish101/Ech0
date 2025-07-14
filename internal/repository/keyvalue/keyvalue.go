package keyvalue

import (
	model "github.com/lin-snow/ech0/internal/model/common"
	"gorm.io/gorm"
)

type KeyValueRepository struct {
	db *gorm.DB
}

func NewKeyValueRepository(db *gorm.DB) KeyValueRepositoryInterface {
	return &KeyValueRepository{
		db: db,
	}
}

// GetKeyValue 根据键获取值
func (keyvalueRepository *KeyValueRepository) GetKeyValue(key string) (interface{}, error) {
	var kv model.KeyValue
	if err := keyvalueRepository.db.Where("`key` = ?", key).First(&kv).Error; err != nil {
		return nil, err
	}

	return kv.Value, nil
}

// AddKeyValue 添加键值对
func (keyvalueRepository *KeyValueRepository) AddKeyValue(key string, value interface{}) error {
	if err := keyvalueRepository.db.Create(&model.KeyValue{
		Key:   key,
		Value: value.(string),
	}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteKeyValue 删除键值对
func (keyvalueRepository *KeyValueRepository) DeleteKeyValue(key string) error {
	if err := keyvalueRepository.db.Where("key = ?", key).Delete(&model.KeyValue{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateKeyValue 更新键值对
func (keyvalueRepository *KeyValueRepository) UpdateKeyValue(key string, value interface{}) error {
	if err := keyvalueRepository.db.Model(&model.KeyValue{}).Where("key = ?", key).Update("value", value.(string)).Error; err != nil {
		return err
	}

	return nil
}
