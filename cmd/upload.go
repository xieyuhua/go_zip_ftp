/*
Copyright © 2022 DuKang <dukang@dukanghub.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/DuKanghub/upload2ftp/pkg"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
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
			fileName = filepath.Base(args[0])
		}
		fileName = fileName + ".zip"
		err := pkg.ZipFiles(args, fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("打包文件成功", fileName)
		ftpConfig := pkg.FtpConfig{
			Host:     ftpHost,
			Port:     ftpPort,
			User:     ftpUser,
			Password: ftpPass,
		}
		ftpCli := pkg.NewFtpClient(ftpConfig)
		err = ftpCli.UploadFile(fileName, ftpDir)
		if err != nil {
			panic(err)
		}
		fmt.Println("上传文件成功", fileName)
		_ = os.Remove(fileName)

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
