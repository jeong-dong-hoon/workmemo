package main

import (
	"bytes"
	"sync"
	"workmemo/astilectron"
	"workmemo/textread"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ast := astilectron.Ast{}
	chan1 := make(chan astilectron.Message, 100)
	chan2 := make(chan astilectron.Message, 100)
	ast.Chan1 = chan1
	ast.Chan2 = chan2
	ast.Setoption("WorkMemo", "/webresource")
	ast.SetWoption(500, 300)
	go MessageListener(chan1, chan2, &wg)
	go ast.MakeWindow(&wg)
	ast.Chan2 <- tmpread()
	wg.Wait()
}

func tmpread() astilectron.Message {
	msg := astilectron.Message{}
	var txt bytes.Buffer
	var content string
	data := textread.StartRead()
	txt.Write(data)
	content = txt.String()
	msg.Content = content
	msg.Code = "first"
	return msg
}
