package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gsurl/log"
	"gsurl/storage"
	"strings"
)

const (
	dict      = "abcdefghijklmnopqrstuvwxyzABCEDFGHIJKLMNOPQRSTUVWXYZ0123456789"
	urlPrefix = "http://127.0.0.0/"
)

func GenShortUrl(ctx context.Context, url string) (string, error) {
	hashVal := hashString(url)
	shortUrl, err := storage.GetUrlByHashCode(ctx, hashVal)
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
	shortUrl, err := storage.GetUrlByShortCode(ctx, shortCode)
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
	return fmt.Sprintf("%s%s", urlPrefix, encode(GenId()))
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
