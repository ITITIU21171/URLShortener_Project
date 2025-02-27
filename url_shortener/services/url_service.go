package services

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	"github.com/ITITIU21171/url-shortener/database"
	"github.com/ITITIU21171/url-shortener/models"
)

func GetOriginalURL(shortURL string) (models.URL, error) {
	var url models.URL

	cacheURL, err := GetCacheURL(shortURL)
	if err != nil {
		log.Println("retrived from redis cahe:", cacheURL)
	}

	err = database.DB.QueryRow(
		"SELECT id, short_url, original_url, created_at, expires_at, visit_count FROM urls WHERE short_url = $1",
		shortURL,
	).Scan(&url.ID, &url.ShortURL, &url.OriginalURL, &url.CreatedAt, &url.ExpiresAt, &url.VisitCount)
	return url, err
}

func CacheURL(shortURL, originalURL string) {
	err := database.RedisClient.Set(context.Background(), shortURL, originalURL, 24*time.Hour).Err()
	if err != nil {
		log.Println("fail to connect URL in redis!", err)
	}
}

func GetCacheURL(shortURL string) (string, error) {
	return database.RedisClient.Get(context.Background(), shortURL).Result()
}

func GenerateShortURL(OriginalURL string) string {
	hash := sha256.Sum256([]byte(OriginalURL))
	return base64.URLEncoding.EncodeToString(hash[:])[:8] // take first 8 characters
}

func SaveURL(originalURL string, shortURL string) error {
	var existingShortURL string

	// Check if the original URL already exists
	err := database.DB.QueryRow(
		"SELECT short_url FROM urls WHERE original_url = $1", originalURL,
	).Scan(&existingShortURL)

	if err == nil {
		log.Println("🔍 URL already exists, returning existing short URL:", existingShortURL)
		return nil // No need to insert again, URL already exists
	} else if err != sql.ErrNoRows {
		log.Println("❌ Database query error:", err)
		return err
	}

	// If the URL is not in the database, insert it
	_, err = database.DB.Exec(
		"INSERT INTO urls (short_url, original_url, created_at) VALUES ($1, $2, NOW())",
		shortURL, originalURL,
	)
	if err != nil {
		log.Println("❌ Error inserting into database:", err)
		return err
	}

	log.Println("✅ URL saved successfully:", shortURL)
	CacheURL(shortURL, originalURL)

	return nil
}
