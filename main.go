package main

import (
	"sync"
	"workmemo/astilectron"
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
	ast.SetWoption(400, 400)
	go ast.MakeWindow()
	go ast.OpenSocekt()
	go ast.OpenListener()
	go MessageListener(chan1, chan2, &wg)
	wg.Wait()
}
