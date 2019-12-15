package tgapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

// DLFile -
func (tg *API) DLFile(filePath string) (file []byte, err error) {
	resp, err := http.Get(fmt.Sprintf(FileEndpoint, tg.AccessToken, filePath))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	file, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// GetFile -
func (tg *API) GetFile(fileID string) (ans APIResponse) {
	fdata := SendGetFile{FileID: fileID}

	return tg.sendJSONData("getFile", fdata)
}

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

// SendMessage - Отправка сообщения
func (tg *API) SendAudio(msg SendAudio) (ans APIResponse) {
	method := "sendAudio"

	// Добавляем прочие параметры
	keys := []string{"chat_id", "disable_notification", "duration", "performer", "title", "audio"}
	values := [][]byte{
		[]byte(strconv.FormatInt(msg.ChatID, 10)),
		[]byte(strconv.FormatBool(msg.DisableNotification)),
		[]byte(strconv.Itoa(msg.Duration)),
		[]byte(msg.Performer),
		[]byte(msg.Title),
		msg.Audio,
	}

	return tg.sendMultipartData(multipartDataObj{
		keys:   keys,
		values: values,
		method: method,
	}, msg.Title)
}

// Отправляем в телеграм Multipart
func (tg *API) sendMultipartData(data multipartDataObj, title string) (ans APIResponse) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// закрываем multipart
	defer w.Close()

	buf := bytes.NewBuffer(data.values[5])
	part, err := w.CreateFormFile("audio", "test.mp3")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(part, buf)

	for i, k := range data.keys {
		var (
			err error
			fw  io.Writer
		)
		if k == "audio" {
			continue
		} else {
			// Add the other fields
			if fw, err = w.CreateFormField(k); err != nil {
				return
			}
		}

		if _, err = fw.Write(data.values[i]); err != nil {
			return
		}
	}
	// закрываем multipart
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", tg.getRequestURL(data.method), &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	log.Println(w.FormDataContentType())

	// Submit the request
	client := &http.Client{Transport: httpTr}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	return tg.checkAnswer(data.method, resp)
}
