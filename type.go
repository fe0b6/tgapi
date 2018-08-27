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
	RetryAfter int `json:"retry_after"`
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

/*
	Типы для отправки в TG
*/

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
