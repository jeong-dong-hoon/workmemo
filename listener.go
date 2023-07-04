package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
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
		case "test":
		case "done":

			if data, err := UnescapeHTML(m.Content); err != nil {
				fmt.Println(err)
			} else {
				textread.SaveTmp(data)

			}
			break mainloop

		}
	}
	wg.Done()
}

func UnescapeHTML(v string) ([]byte, error) {

	Buffer := &bytes.Buffer{}
	Encoder := json.NewEncoder(Buffer)
	Encoder.SetEscapeHTML(false)
	err := Encoder.Encode(strings.ReplaceAll(v, "\\", ""))
	fmt.Println(Buffer.String())
	return Buffer.Bytes(), err
}
