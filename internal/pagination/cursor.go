package pagination

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	DefaultLimit = 50
	MaximumLimit = 200
)

type Cursor struct {
	Time time.Time `json:"time"`
	ID   string    `json:"id"`
}

type Page[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor,omitempty"`
}

func NormalizeLimit(value int) int {
	if value <= 0 {
		return DefaultLimit
	}
	if value > MaximumLimit {
		return MaximumLimit
	}
	return value
}

func Encode(cursor Cursor) (string, error) {
	if cursor.Time.IsZero() || strings.TrimSpace(cursor.ID) == "" {
		return "", fmt.Errorf("cursor requires time and id")
	}
	payload, err := json.Marshal(cursor)
	if err != nil {
		return "", fmt.Errorf("encode cursor payload: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(payload), nil
}

func Decode(value string) (Cursor, error) {
	if strings.TrimSpace(value) == "" {
		return Cursor{}, nil
	}
	payload, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return Cursor{}, fmt.Errorf("decode cursor: %w", err)
	}
	var cursor Cursor
	if err := json.Unmarshal(payload, &cursor); err != nil {
		return Cursor{}, fmt.Errorf("parse cursor: %w", err)
	}
	if cursor.Time.IsZero() || strings.TrimSpace(cursor.ID) == "" {
		return Cursor{}, fmt.Errorf("cursor is incomplete")
	}
	return cursor, nil
}
