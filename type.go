package tgapi

import (
	"encoding/json"
	"sync"
)

// API - Базовый объект
type API struct {
	AccessToken   string
	RetryDontWait bool
	retryCount    int
	sync.Mutex
}

// APIResponse is a response from the Telegram API with the result stored raw.
type APIResponse struct {
	Ok          bool                  `json:"ok"`
	Result      json.RawMessage       `json:"result"`
	ErrorCode   int                   `json:"error_code"`
	Description string                `json:"description"`
	Parameters  APIResponseParameters `json:"parameters"`
}

// APIResponseParameters - параметры ответа
type APIResponseParameters struct {
	RetryAfter      int   `json:"retry_after"`
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
}

/*
	Типы для получения из TG
*/

// Update - Обновление от телеграма
type Update struct {
	UpdateID      int64         `json:"update_id"`
	Message       Message       `json:"message"`
	InlineQuery   InlineQuery   `json:"inline_query"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// User is a user, contained in Message and returned by GetSelf.
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

// Chat is returned in Message, because it's not clear which it is.
type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Message is returned by almost every request, and contains data about almost anything.
type Message struct {
	MessageID int64  `json:"message_id"`
	From      User   `json:"from"`
	Date      int64  `json:"date"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
	Audio     Audio  `json:"audio"`
}

// Audio audio type
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int64  `json:"file_size"`
}

// InlineQuery - Inline запрос
type InlineQuery struct {
	ID     string `json:"id"`
	From   User   `json:"user"`
	Query  string `json:"query"`
	Offset string `json:"offset"`
}

// CallbackQuery запрос
type CallbackQuery struct {
	ID              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	Data            string  `json:"data"`
}

type File struct {
	FileID   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	FilePath string `json:"file_path"`
}

/*
	Типы для отправки в TG
*/

// SendGetFile - инфа о файле
type SendGetFile struct {
	FileID string `json:"file_id"`
}

// ReplyKeyboardMarkup allows the Bot to set a custom keyboard.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard"`
	OneTimeKeyboard bool       `json:"one_time_keyboard"`
	Selective       bool       `json:"selective"`
}

// ReplyKeyboardRemove - удаление клавиатуры
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// SendMessageData - Сообщение
type SendMessageData struct {
	ChatID                interface{} `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview"`
	ReplyToMessageID      int64       `json:"reply_to_message_id"`
	ReplyMarkup           interface{} `json:"reply_markup"`
	DisableNotification   bool        `json:"disable_notification"`
}

// SendAudio - Аудио для отправки
type SendAudio struct {
	ChatID              int64  `json:"chat_id,omitempty"`
	Audio               []byte `json:"audio,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	Performer           string `json:"performer,omitempty"`
	Title               string `json:"title,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
}

// Структура для аплоада данных
type multipartDataObj struct {
	keys     []string
	values   [][]byte
	method   string
	filePath string
	fileName string
}
