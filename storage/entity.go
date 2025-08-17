package storage

import (
	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model        // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	ShortCode  string `gorm:"type:varchar(255);uniqueIndex;not null" json:"short_url"` // 短链接标识
	OriginUrl  string `gorm:"type:varchar(2048);not null" json:"origin_url"`           // 原始链接
	HashCode   string `gorm:"type:varchar(255);uniqueIndex;not null" json:"hash_code"` // 哈希码
}
