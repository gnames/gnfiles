/*
Copyright © 2021 Dmitry Mozzherin <dmozzherin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/gnames/gnfiles"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Updates IPFS data from local directory.",
	Long: `Takes a directory and a IPNS key name where to place data and
uploads files and metadata. If the key is not given, returns IPFS ID for
the metadata file.

gnfiles <dir> <key_name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := gnfiles.NewConfig(opts...)
		processSyncArgs(cmd, args, cfg)
		if cfg.KeyName == "" {
			log.Print("Did not get IPNS key, perma-link will not be generated.")
		}
		if cfg.Dir == "" {
			log.Fatal("Did not get a directory to sync with IPFS.")
		}
		cfg.WithUpload = true
		gnf := gnfiles.New(cfg)
		err := gnf.Upload()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func processSyncArgs(
	cmd *cobra.Command,
	args []string,
	cfg *gnfiles.Config,
) {
	var dir, keyName string
	switch len(args) {
	case 2:
		dir = args[0]
		keyName = args[1]
	case 1:
		arg := args[0]
		if len(arg) > 40 || strings.HasPrefix(arg, "k5") {
			keyName = arg
		} else {
			dir = arg
		}
	default:
		_ = cmd.Help()
		os.Exit(0)
	}
	if keyName != "" {
		cfg.KeyName = keyName
	}
	if dir != "" {
		cfg.Dir = dir
	}
}
