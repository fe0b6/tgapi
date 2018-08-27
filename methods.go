package tgapi

import (
	"strings"
	"unicode/utf8"
)

// GetMe - Получаем инфу о боте
func (tg *API) GetMe() (ans APIResponse) {
	method := "getMe"

	return tg.sendJSONData(method, nil)
}

// SetWebhook - Установка Webhook
func (tg *API) SetWebhook(url string) (ans APIResponse) {
	method := "setWebhook"

	m := map[string]string{"url": url}

	return tg.sendJSONData(method, m)
}

// SendMessageBig - Отправка сообщения с проверкой на длину
func (tg *API) SendMessageBig(msg SendMessageData) (ans []APIResponse) {
	ans = []APIResponse{}

	// Если длина текста влезет в одно сообщение - просто отправляем
	if utf8.RuneCountInString(msg.Text) < TextMaxSize {
		ans = append(ans, tg.SendMessage(msg))
		return
	}

	// Разбиваем текст на блоки нужной длины
	texts := []string{}
	var tmp string
	for _, v := range strings.Split(msg.Text, " ") {
		// Если длина куска будет больше чем максимум - сохраняем предыдущий и начинаем новый кусок
		if utf8.RuneCountInString(tmp+v) > (TextMaxSize - 1) {
			texts = append(texts, tmp)
			tmp = ""
		}

		tmp += v + " "
	}
	// Не забываем добавить остаток текста
	if len(tmp) > 0 {
		texts = append(texts, tmp)
	}

	// Отправляем куски
	for _, text := range texts {
		msg.Text = text
		ans = append(ans, tg.SendMessage(msg))
	}

	return
}

// SendMessage - Отправка сообщения
func (tg *API) SendMessage(msg SendMessageData) (ans APIResponse) {
	method := "sendMessage"

	// Если клавиатура не указана - делаем пустую
	if msg.ReplyMarkup == nil {
		msg.ReplyMarkup = ReplyKeyboardMarkup{Keyboard: [][]string{}}
	}

	return tg.sendJSONData(method, msg)
}
