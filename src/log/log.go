package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Raw     *log.Logger
)

func init() {
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ltime)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ltime)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ltime)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ltime)

	Raw = log.New(os.Stdout,
		"",
		log.Ltime)
	Raw.SetFlags(0)
}
