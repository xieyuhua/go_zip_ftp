/*
 * @Descripttion:
 * @version:
 * @Author: seaslog
 * @Date: 2022-03-12 14:31:28
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-03-12 14:32:44
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string
	ftpHost string
	ftpPort string
	ftpUser string
	ftpPass string
	ftpDir  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "upload2ftp",
	Short: "上传到ftp服务器",
	Long:  `将文件上传到ftp服务器`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.upload2ftp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&ftpHost, "server", "s", "", "ftp地址，必传")
	rootCmd.PersistentFlags().StringVarP(&ftpPort, "port", "P", "21", "ftp端口，默认21")
	rootCmd.PersistentFlags().StringVarP(&ftpUser, "user", "u", "", "ftp用户名，必传")
	rootCmd.PersistentFlags().StringVarP(&ftpPass, "password", "p", "", "ftp密码，必传")
	rootCmd.PersistentFlags().StringVarP(&ftpDir, "ftpDir", "d", "/", "ftp远程目录，非必须，默认为根目录")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".upload2ftp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".upload2ftp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
