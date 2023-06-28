package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"workmemo/astilectron"
	"workmemo/textread"
)

func MessageListener(message chan astilectron.Message, sendmessage chan astilectron.Message, wg *sync.WaitGroup) {
mainloop:
	for {
		m := <-message
		switch m.Code {
		case "take":
		case "save":
			{
				msg := astilectron.Message{}
				for i := 0; i < len(m.Content); i++ {

				}
				msg.Code = "done"
				sendmessage <- msg
			}
		case "done":
			fmt.Println(m)
			if data, err := json.Marshal(m.Content); err != nil {
				fmt.Println(err)
			} else {
				textread.SaveTmp(data)

			}
			break mainloop
		}
	}
	wg.Done()
}
