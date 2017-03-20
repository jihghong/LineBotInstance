// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
			    if strings.Contains(message.Text, "我是誰") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您是昭宗大哥嗎？")).Do();
				} else if strings.Contains(message.Text, "不對") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您是信宏兄嗎？")).Do();
				} else if strings.Contains(message.Text, "不是") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您是家興嗎？")).Do();
				} else if strings.Contains(message.Text, "再猜") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我知道了，您是志鴻！")).Do();
				} else if strings.Contains(message.Text, "答對") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("需要我做些什麼嗎？")).Do();
				} else if strings.Contains(message.Text, "HIGGS.csv") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("HIGGS.csv 檔案超大，要跑很久喔！")).Do();
				} else if strings.Contains(message.Text, "算了") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("再下一個命令？")).Do();
				} else if strings.Contains(message.Text, "pca") {
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://raw.githubusercontent.com/jihghong/LineBotInstance/master/images/pca.PNG", "https://raw.githubusercontent.com/jihghong/LineBotInstance/master/images/pca.PNG")).Do();
				} else {
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎下次再來")).Do();
				}
			
			case *linebot.ImageMessage:
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哈哈哈！有創意")).Do();
			｝
		}
	}
}
