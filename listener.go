package main

import (
	"sync"
	"workmemo/astilectron"
)

func MessageListener(message chan astilectron.Message, sendmessage chan astilectron.Message, wg *sync.WaitGroup) {
mainloop:
	for {
		m := <-message
		switch m.Code {
		case "take":
		case "save":
			{
				for i := 0; i < len(m.Content); i++ {

				}
			}
		case "done":
			break mainloop
		}
	}
	wg.Done()
}
