package main

import (
	"time"
	"log"
	"fmt"
	"strconv"
	"net/url"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func gSisBot() {
	// Google Sis Bot
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
		b.Send(m.Sender, "This bot is currently not supported interactive mode. Please use the inline mode (@GoogleSisBot your_chat)")
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
				URL: googleTranslateUrl + url.QueryEscape(q.Text),
			}

			results[0] = result
			results[0].SetResultID(strconv.Itoa(0))
		
			err := b.Answer(q, &tb.QueryResponse{
				Results: results,
				CacheTime: 60, // in sec
			})
		
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	b.Start()
}

func the2ndBot() {

	// Second bot
	the2ndBot, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("SECRET_TOKEN_2ND"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	the2ndBot.Handle("/start", func(m *tb.Message) {
		the2ndBot.Send(m.Sender, "2nd bot running too")
	})

	the2ndBot.Start()
}

func main() {

	go gSisBot()
	// go the2ndBot()


	// Second bot
	the2ndBot, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("SECRET_TOKEN_2ND"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	the2ndBot.Handle("/start", func(m *tb.Message) {
		the2ndBot.Send(m.Sender, "2nd bot running too")
	})

	the2ndBot.Start()
	
}
