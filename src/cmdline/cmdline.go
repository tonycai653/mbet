package cmdline

import (
	"flag"
)

var Debug bool

func init() {
	flag.BoolVar(&Debug, "-debug", true, "turn on debug mode")
}
