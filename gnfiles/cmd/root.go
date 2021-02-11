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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gnames/gnfiles"
	"github.com/gnames/gnsys"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configText = `# Uncomment lines that you want to update.

# APIURL the URL of IPFS API, usually "localhost:5001"
# APIURL: localhost:5001

# KeyName is needed only if you want to save files for your own use, and
# not just download them. If the authors of the directory did not send
# you the exported key, you will not update the their dir, but will create
# you own. Make sure you give newly generated key if you want to share
# the dir.
# KeyName:

# Source is a string that provides information where to get metadata.
# Source can be one of the following:
#
# - Path to a file on a local filesystem
# - URL to metadata file
# - IPFS CID of metadata content
# - IPFS path (/ipfs/Qm...) pointing to metadata
# - IPNS key (k5....) pointint to metadata
# - IPNS path (/ipns/k5....)
#
# Source:

# Dest is the path to the directory where you want to download files.
# Dest:
`
)

var (
	opts    []gnfiles.Option
	cfgFile string
)

type cfgData struct {
	APIURL  string
	KeyName string
	ID      string
	Dir     string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gnfiles",
	Short: "Uses IPFS to store and retrieve directories.",
	Long:  `Uses IPFS to store and download directories`,
	Run: func(cmd *cobra.Command, args []string) {
		if flagVersion(cmd) {
			os.Exit(0)
		}
		flagAPIURL(cmd)
		_ = cmd.Help()
		os.Exit(0)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP("version", "V", false,
		"shows gnfiles version")
	rootCmd.PersistentFlags().StringP(
		"api_url",
		"a",
		"",
		"URL of IPFS API; 'localhost:5001' is default")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configFile := "gnfiles"

	homeConfig, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Cannot find home config directory: %s.", err)
	}

	// Search config in home directory with name ".gnames" (without extension).
	viper.AddConfigPath(homeConfig)
	viper.SetConfigName(configFile)

	configPath := filepath.Join(homeConfig, fmt.Sprintf("%s.yaml", configFile))
	touchConfigFile(configPath, configFile)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	getOpts()
}

func getOpts() []gnfiles.Option {
	cfg := &cfgData{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Cannot deserialize config data: %s.", err)
	}
	if cfg.APIURL != "" {
		opts = append(opts, gnfiles.OptApiURL(cfg.APIURL))
	}
	if cfg.KeyName != "" {
		opts = append(opts, gnfiles.OptKeyName(cfg.KeyName))
	}
	if cfg.ID != "" {
		opts = append(opts, gnfiles.OptSource(cfg.ID))
	}
	if cfg.Dir != "" {
		opts = append(opts, gnfiles.OptDir(cfg.Dir))
	}

	return opts
}

func flagVersion(cmd *cobra.Command) bool {
	version, err := cmd.Flags().GetBool("version")
	if err != nil {
		log.Fatal(err)
	}
	if version {
		fmt.Printf("\nversion: %s\n\n",
			gnfiles.Version)
		return true
	}
	return false
}

func flagAPIURL(cmd *cobra.Command) {
	s, err := cmd.Flags().GetString("api_url")
	if err == nil && s != "" {
		opts = append(opts, gnfiles.OptApiURL(s))
	}
}

// touchConfigFile checks if config file exists, and if not, it gets created.
func touchConfigFile(configPath string, configFile string) {
	exists, err := gnsys.FileExists(configPath)
	if err != nil {
		log.Printf("Cannot use '%s' as config file: %v\n", configPath, err)
	}
	if exists {
		return
	}

	log.Printf("Creating config file: %s.", configPath)
	createConfig(configPath, configFile)
}

// createConfig creates config file.
func createConfig(path string, file string) {
	err := gnsys.MakeDir(filepath.Dir(path))
	if err != nil {
		log.Fatalf("Cannot create dir %s: %s.", path, err)
	}

	err = ioutil.WriteFile(path, []byte(configText), 0644)
	if err != nil {
		log.Fatalf("Cannot write to file %s: %s.", path, err)
	}
}
