package storage

import (
	"context"
	"errors"
	"gsurl/log"

	"gorm.io/gorm"
)

func SaveUrl(ctx context.Context, shortUrl *ShortUrl) error {
	// Simulate saving the URL to a database or storage.
	// In a real application, you would use a database connection here.
	log.Logger.Infof("Saving URL: %s with short code: %s", shortUrl.OriginUrl, shortUrl.ShortCode)
	ret := db.WithContext(ctx).Model(shortUrl).Create(shortUrl)
	if ret.Error != nil {
		log.Logger.Errorf("Failed to save URL: %v", ret.Error)
		return ret.Error
	}
	return nil
}

func GetUrlByShortCode(ctx context.Context, shortCode string) (*ShortUrl, error) {
	var shortUrl ShortUrl
	ret := db.WithContext(ctx).Where("short_code = ?", shortCode).First(&shortUrl)
	if ret.Error != nil {
		if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
			log.Logger.Infof("No URL found for short code %s", shortCode)
			return nil, nil
		}
		log.Logger.Errorf("Failed to get URL by short code %s: %v", shortCode, ret.Error)
		return nil, ret.Error
	}
	return &shortUrl, nil
}

func GetUrlByHashCode(ctx context.Context, hashCode string) (*ShortUrl, error) {
	var shortUrl ShortUrl
	ret := db.WithContext(ctx).Where("hash_code = ?", hashCode).First(&shortUrl)
	if ret.Error != nil {
		if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
			log.Logger.Infof("No URL found for hash code %d", hashCode)
			return nil, nil
		}
		log.Logger.Errorf("Failed to get URL by hash code %d: %v", hashCode, ret.Error)
		return nil, ret.Error
	}
	return &shortUrl, nil
}
