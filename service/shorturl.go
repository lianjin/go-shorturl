package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gsurl/config"
	"gsurl/log"
	"gsurl/storage"
	"strings"
	"time"
)

const (
	dict                  = "abcdefghijklmnopqrstuvwxyzABCEDFGHIJKLMNOPQRSTUVWXYZ0123456789"
	cacheExpireTime4Write = time.Second * 3
	cacheExpireTime4Read  = time.Hour * 24
)

func GenShortUrl(ctx context.Context, url string) (string, error) {
	hashVal := hashString(url)
	if val, exist := GetFromCache(hashVal); exist {
		return formstShortUrl(val), nil
	}
	shortUrl, err := storage.GetUrlByHashCode(ctx, hashVal)
	defer func() {
		if err == nil && shortUrl != nil {
			PutCache(hashVal, shortUrl.ShortCode, cacheExpireTime4Write)
		}
	}()
	if err != nil {
		log.Logger.Errorf("Error retrieving URL by hash code: %v", err)
		return "", err
	}
	if shortUrl != nil {
		log.Logger.Warnf("Found existing short URL: %s for original URL: %s", shortUrl.ShortCode, url)
		return formstShortUrl(shortUrl.ShortCode), nil
	}
	shortUrl = &storage.ShortUrl{
		ShortCode: encode(GenId()),
		OriginUrl: url,
		HashCode:  hashVal,
	}
	if err := storage.SaveUrl(ctx, shortUrl); err != nil {
		log.Logger.Errorf("Error saving short URL: %v", err)
		return "", err
	}
	log.Logger.Infof("Generated new short URL: %s for original URL: %s", shortUrl.ShortCode, url)
	return formstShortUrl(shortUrl.ShortCode), nil
}

func GetShortUrl(ctx context.Context, shortCode string) (string, error) {
	if val, exist := GetFromCache(shortCode); exist {
		return val, nil
	}
	shortUrl, err := storage.GetUrlByShortCode(ctx, shortCode)
	defer func() {
		if err == nil && shortUrl != nil {
			PutCache(shortCode, shortUrl.OriginUrl, cacheExpireTime4Read)
		}
	}()
	if err != nil {
		log.Logger.Errorf("Error retrieving URL by short code: %v", err)
		return "", err
	}
	if shortUrl == nil {
		log.Logger.Warnf("No URL found for short code: %s", shortCode)
		return "", nil
	}
	return shortUrl.OriginUrl, nil
}

func formstShortUrl(code string) string {
	return fmt.Sprintf("%s/%s", config.AppConfig.Host, code)
}

func encode(num int64) string {
	scale := int64(len(dict))
	var ret strings.Builder
	for num > 0 {
		idx := num % scale
		ret.WriteByte(dict[idx])
		num = num / scale
	}
	return ret.String()
}

func hashString(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
