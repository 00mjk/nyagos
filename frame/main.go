package frame

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime/debug"

	"github.com/go-ole/go-ole"

	"github.com/zetamatta/nyagos/alias"
	"github.com/zetamatta/nyagos/commands"
	"github.com/zetamatta/nyagos/completion"
	"github.com/zetamatta/nyagos/history"
	"github.com/zetamatta/nyagos/shell"
)

var DefaultHistory *history.Container

func Start(mainHandler func() error) error {
	defer PanicHandler()

	shell.SetHook(func(ctx context.Context, it *shell.Cmd) (int, bool, error) {
		rc, done, err := commands.Exec(ctx, it)
		return rc, done, err
	})
	completion.AppendCommandLister(commands.AllNames)
	completion.AppendCommandLister(alias.AllNames)

	if ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED) == nil {
		defer ole.CoUninitialize()
	}

	alias.Init()

	return mainHandler()
}

func PanicHandler() {
	err := recover()
	if err == nil {
		return
	}
	var dump bytes.Buffer
	w := io.MultiWriter(os.Stderr, &dump)

	fmt.Fprintln(w, "************ Panic Occurred. ***********")
	fmt.Fprintln(w, err)
	w.Write(debug.Stack())
	fmt.Fprintln(w, "*** Please copy these error message ***")
	fmt.Fprintln(w, "*** And hit ENTER key to quit.      ***")

	ioutil.WriteFile("nyagos.dump", dump.Bytes(), 0666)

	var dummy [1]byte
	os.Stdin.Read(dummy[:])
}
