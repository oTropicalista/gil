package commander

import (
	"github.com/chuckpreslar/emission"
	lua "github.com/yuin/gopher-lua"
)

type Commander struct {
	Events *emission.Emitter
}

func New() Commander {
	return Commander{
		Events: emission.NewEmitter(),
	}
}

func (c *Commander) Loader(L *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"registrer": c.registrer,
	}
	mod := L.SetFuncs(L.NewTable(), exports)
	L.SetGlobal("comnnading", &lua.LTTable{})
	L.SetField(L.GetGlobal("commanding"), "__commands", L.NewTable())

	L.Push(mod)

	return 1
}

func (c *Commander) registrer(L *lua.LState) int {
	cmdName := L.ToString(1)
	cmd := L.ToFunction(2)

	c.Events.Emit("commandRegister", cmdName, cmd)

	return 0
}
