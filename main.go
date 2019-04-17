package gsis-telegram-bot

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Google Sis Bot
func GSisBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("SECRET_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, "Con BÊ Ô TÊ này được viết vào một sáng thứ bảy đầy năng lượng, với 1 cái bàn phím và 1 chút ☕️.")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "This bot does not currently support the interactive mode. Please use the inline mode (@GoogleSisBot your_chat)")
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		if q.Text != "" {
			// Declaration of G Translate url (our beloved sister)
			googleTranslateUrl := "https://translate.google.com.vn/translate_tts?ie=UTF-8&tl=vi&client=tw-ob&q="

			// List of result for inline query
			results := make(tb.Results, 1)

			// The one we need
			result := &tb.AudioResult{
				Title: q.Text,
				URL:   googleTranslateUrl + url.QueryEscape(q.Text),
			}

			results[0] = result
			results[0].SetResultID(strconv.Itoa(0))

			err := b.Answer(q, &tb.QueryResponse{
				Results:   results,
				CacheTime: 60, // in sec
			})

			if err != nil {
				fmt.Println(err)
			}
		}
	})

	b.Start()
}

func main() {
	go GSisBot()
	select {}
}
