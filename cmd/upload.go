/*
 * @Descripttion:
 * @version:
 * @Author: seaslog
 * @Date: 2022-03-12 14:31:28
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-03-12 15:22:31
 */
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DuKanghub/upload2ftp/pkg"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件",
	Long:  `上传文件到FTP服务器`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := time.Now().Local().Format("2006-01-02-15-04-05")
		if len(args) == 0 {
			fmt.Println("请输入文件路径")
			return
		} else if len(args) == 1 {
			// fileName = args[0]
			fileName = filepath.Base(args[0])
		}
		// fileName = fileName + ".zip"
		// err := pkg.ZipFiles(args, fileName)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		fmt.Println("打包文件成功", args[0])
		ftpConfig := pkg.FtpConfig{
			Host:     ftpHost,
			Port:     ftpPort,
			User:     ftpUser,
			Password: ftpPass,
		}
		ftpCli := pkg.NewFtpClient(ftpConfig)
		err := ftpCli.UploadFile(args[0], ftpDir)
		if err != nil {
			panic(err)
		}
		fmt.Println("上传文件成功", fileName)
		_ = os.Remove(args[0])

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
