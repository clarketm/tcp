/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
    tcp – list open TCP connections.

SYNOPSIS:
    tcp [ opts... ]

OPTIONS:
    -h, --help                  print usage information.
    -v, --version               print version number.
    -a, --all                   same as -f, -s, -e.
    -e, --est                   include `established` connections in the output.
    -f, --ftp                   include `ftp` connections in the output.
    -s, --ssh                   include `ssh` connections in the output.
    -n, --num                   show network numbers instead of host names.
    -p, --pretty                pretty print output.

EXAMPLES:
    TODO

*/

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/clarketm/tcp/utils"
	flag "github.com/spf13/pflag"
)

// VERSION - current version number
const VERSION = "v1.0.0"

// Flags
var version bool

// init () - initialize command-line flags
func init() {
	// -i, --input
	// flag.VarP(&inputFormat, "input", "i", "input `format`: (a)scii, (b)inary, (o)ctal, (d)ecimal, (h)exadecimal")

	// -o, --output
	// flag.VarP(&outputFormat, "output", "o", "output `format`: (a)scii, (b)inary, (o)ctal, (d)ecimal, (h)exadecimal")

	// -v, --version
	// flag.BoolVarP(&version, "version", "v", false, "print version number")

	// Usage
	flag.Usage = func() {
		println()
		fmt.Printf("NAME:\n")
		fmt.Printf("\ttcp – list open TCP connections.\n")
		println()
		fmt.Printf("SYNOPSIS:\n")
		fmt.Printf("\t%v [ opts... ]\n", utils.Bold("tcp"))
		println()
		fmt.Printf("OPTIONS:\n")
		flag.PrintDefaults()
		println()
		os.Exit(utils.StatusCode)
	}
}

func printVersion() {
	fmt.Printf("\n%s %v\n", utils.Bold("Version:"), VERSION)
	os.Exit(0)
}

// main ()
func main() {
	// flag.Parse()

	// if version {
	// printVersion() // version and EXIT
	// }

	// Mac - lsof -i tcp | grep -i "listen"
	lsof := exec.Command("lsof", "-i", "tcp")
	grep := exec.Command("grep", "-i", "listen", "--color=always")
	grep.Stdin, _ = lsof.StdoutPipe()
	grep.Stdout = os.Stdout
	_ = grep.Start()
	_ = lsof.Run()
	_ = lsof.Wait()

	// Linux - TODO

	// Windows - TODO

}
