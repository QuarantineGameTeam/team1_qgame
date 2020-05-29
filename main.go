package main

import (
	"fmt"
	"log"
	"net/http"
	"team1_qgame/conf"
	"team1_qgame/database"
	"strconv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//for registration
var RegFlag bool = false

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(conf.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	bot.SetWebhook(tgbotapi.NewWebhook(conf.WEB_HOOK))
}

func GetUser(msg *tgbotapi.Message) conf.User {
	user := conf.User{Id: msg.Chat.ID, FirstName: msg.Chat.FirstName, ClanName : "empty"}
	fmt.Println(user)
	return user
}

func Registration(msg *tgbotapi.MessageConfig, bot *tgbotapi.BotAPI){
	var chooseClan = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("clan1"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("clan2"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("clan3"),
		),
	)
	msg.ReplyMarkup = chooseClan
	msg.Text = "Оберіть клан"
	bot.Send(msg)
}


func getUpdates(bot *tgbotapi.BotAPI) {
	setWebhook(bot)
	updates := bot.ListenForWebhook("/")
	for update := range updates {
		
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "<empty>")
			var user conf.User = GetUser(update.Message)
			
			//if registration is in progress and we're waiting for clan info
			if RegFlag {
				user.ClanName = update.Message.Text
				db.SaveUser(&user)
				msg.Text = "Готово"
				bot.Send(msg)
				RegFlag = false
			} 
			if update.Message.Text != "/start" {
				user = db.GetUser(strconv.Itoa(int(update.Message.Chat.ID)))
			}
			switch update.Message.Text {
			case "/start":
				if user.ClanName != "empty"{
					Registration(&msg, bot)
					RegFlag = true
				}
			case "/me":
				user = db.GetUser(strconv.Itoa(int(update.Message.Chat.ID)))
				msg.Text = "Привіт, " + user.FirstName + " " + user.ClanName
				bot.Send(msg)
			}
		}
	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":"+conf.BOT_PORT, nil))
	}()

	getUpdates(NewBot)
}