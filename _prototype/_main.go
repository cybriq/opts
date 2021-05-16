package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/cybriq/qu"

	//_ "github.com/p9c/p9/pkg/gel/gio/app/permission/networkstate" // todo: integrate this into routeable package
	//_ "github.com/p9c/p9/pkg/gel/gio/app/permission/storage"      // this enables the home folder appdata directory to work on android (and ios)

	log2 "github.com/cybriq/log"
	"github.com/cybriq/opts/_prototype/config"
	"github.com/cybriq/opts/_prototype/podcfgs"
	"github.com/cybriq/opts/_prototype/podhelp"
	"github.com/cybriq/opts/_prototype/state"
	"github.com/cybriq/opts/version"

	// This ensures the database drivers get registered
	//_ "github.com/p9c/p9/pkg/database/ffldb"

	// _ "github.com/p9c/p9/pkg/gel/gio/app/permission/bluetooth"
	// _ "github.com/p9c/p9/pkg/gel/gio/app/permission/camera"
)

func main() {
	<-Main()
}

func Main() (quit qu.C) {
	quit = qu.T()
	go func() {
		log2.SetLogLevel("trace")
		log.T.Ln(os.Args)
		log.T.Ln(version.Get())
		var cx *state.State
		var e error
		if cx, e = state.GetNew(podcfgs.GetDefaultConfig(), podhelp.HelpFunction, quit); E.Chk(e) {
			fail()
		}

		// fail()
		// if e = debugConfig(cx.Config); E.Chk(e) {
		// }

		log.D.Ln("running command:", cx.Config.RunningCommand.Name)
		if e = cx.Config.RunningCommand.Entrypoint(cx); E.Chk(e) {
			fail()
		}
		quit.Q()
	}()
	return quit
}

func fail() {
	os.Exit(1)
}

func debugConfig(c *config.Config) (e error) {
	c.ShowAll = true
	defer func() { c.ShowAll = false }()
	var j []byte
	if j, e = c.MarshalJSON(); E.Chk(e) {
		return
	}
	var b []byte
	jj := bytes.NewBuffer(b)
	if e = json.Indent(jj, j, "", "\t"); log.E.Chk(e) {
		return
	}
	log.T.Ln("\n"+jj.String())
	return
}
