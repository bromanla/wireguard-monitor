package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wireguard-monitor/internal/common"
)

// TelegramMessage represents the payload for Telegram Bot API.
type TelegramMessage struct {
	Text      string `json:"text"`
	ChatID    int64  `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
}

// SendTelegram posts a message to the configured Telegram chat.
func SendTelegram(message string) error {
	token := common.GetEnv("TELEGRAM_TOKEN")

	chatIdStr := common.GetEnv("TELEGRAM_CHAT_ID")
	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid chat ID %q: %w", chatIdStr, err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := TelegramMessage{
		Text:      message,
		ChatID:    chatId,
		ParseMode: "HTML",
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal telegram payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram returned non 200 status: %s", resp.Status)
	}

	return nil
}
