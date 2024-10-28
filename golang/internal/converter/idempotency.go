package converter

import (
	"database/sql"
	"encoding/json"
    "golang.org/x/exp/slog"
	"time"
)

func IsProcessed(db *sql.DB, videoID int) bool {
	var isProcessed bool
	query := "SELECT EXISTS(SELECT 1 FROM processed_videos WHERE video_id = $1 AND status = 'success')"
	err := db.QueryRow(query, videoID).Scan(&isProcessed)
	if err != nil {
		slog.Error("Error checking if video is processed", slog.Int("video_id", videoID), slog.Any("error", err)) // Use slog.Any para erros
		return false
	}
	return isProcessed
}

func MarkProcessed(db *sql.DB, videoID int) error {
	query := "INSERT INTO processed_videos (video_id, status, processed_at) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, videoID, "success", time.Now())
	if err != nil {
		slog.Error("Error marking video as processed", slog.Int("video_id", videoID), slog.Any("error", err)) // Use slog.Any para erros
		return err
	}
	return nil
}

func RegisterError(db *sql.DB, errorData map[string]interface{}, err error) {
	serializedError, _ := json.Marshal(errorData)
	query := "INSERT INTO process_errors_log (error_details, created_at) VALUES ($1, $2)"
	_, dbErr := db.Exec(query, string(serializedError), time.Now()) 
	if dbErr != nil {
		slog.Error("Error registering error", slog.String("error_details", string(serializedError)), slog.Any("error", dbErr)) // Use slog.Any para erros
	}
}
