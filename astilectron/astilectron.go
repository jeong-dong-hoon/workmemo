package astilectron

import (
	"fmt"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

type Ast struct {
	Option      astilectron.Options
	Woption     *astilectron.WindowOptions
	Window      *astilectron.Window
	Astilectron *astilectron.Astilectron
	Chan1       chan Message
	Chan2       chan Message
}
type Message struct {
	Code    string
	Content map[int]interface{}
}

func (a *Ast) Setoption(name string, Dirpath string) {
	option := astilectron.Options{}
	option.AppName = name
	option.BaseDirectoryPath = Dirpath
	a.Option = option

}
func (a *Ast) SetWoption(height int, width int) {
	option := astilectron.WindowOptions{}
	option.Height = astikit.IntPtr(height)
	option.Width = astikit.IntPtr(width)
	a.Woption = &option
}
func (a *Ast) MakeWindow() {
	ast, err := astilectron.New(nil, a.Option)

	if err != nil {
		fmt.Println(err)
	}
	ast.HandleSignals()
	if err = ast.Start(); err != nil {
		fmt.Println(err)
	}
	a.Astilectron = ast

	if a.Window, err = a.Astilectron.NewWindow("./webresource/index.html", a.Woption); err != nil {
		fmt.Println(err)
	}
	if a.Window.Create(); err != nil {
		fmt.Println(err)
	}
	a.Astilectron.Wait()

}

func (a *Ast) OpenListener() {
	for {

		a.Window.OnMessage(func(m *astilectron.EventMessage) interface{} {
			s := Message{}
			m.Unmarshal(&s)
			a.Chan1 <- s
			return nil
		})
	}
}
func (a *Ast) OpenSocekt() {
	for {
		m := <-a.Chan2
		a.Window.SendMessage(m)
	}
}
