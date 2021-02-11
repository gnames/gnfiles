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

	"github.com/gnames/gnfiles"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Downloads a directory from IPFS",
	Long: `Takes IPFS or IPNS ID to download a directory from IPFS.

gnfiles get [id] [dir]
`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := gnfiles.NewConfig(opts...)
		processGetArgs(cmd, args, cfg)
		if cfg.Source == "" {
			log.Fatal("Did not get the source.")
		}
		if cfg.Dir == "" {
			log.Fatal("Did not get destination directory.")
		}
		cfg.WithUpload = false
		gnf := gnfiles.New(cfg)
		err := gnf.Sync()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func processGetArgs(
	cmd *cobra.Command,
	args []string,
	cfg *gnfiles.Config,
) {
	var source, dir string
	switch len(args) {
	case 2:
		source = args[0]
		dir = args[1]
	case 1:
		if cfg.Source == "" {
			source = args[0]
		} else {
			dir = args[0]
		}
	case 0:
		if cfg.Source == "" || cfg.Dir == "" {
			_ = cmd.Help()
			os.Exit(0)
		}
	default:
		_ = cmd.Help()
		os.Exit(0)
	}
	if source != "" {
		cfg.Source = source
	}
	if dir != "" {
		cfg.Dir = dir
	}
}