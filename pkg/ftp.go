package pkg

import (
	"github.com/jlaffaye/ftp"
	"os"
	"time"
)

type FtpConfig struct {
	Host     string
	User     string
	Port     string
	Password string
}

type FtpSaver struct {
	client *ftp.ServerConn
}

func NewFtpClient(a FtpConfig) *FtpSaver {
	conn, err := ftp.Dial(a.Host+":"+a.Port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	err = conn.Login(a.User, a.Password)
	if err != nil {
		panic(err)
	}
	return &FtpSaver{
		client: conn,
	}
}

// UploadFile 上传文件到ftp服务器
func (cli *FtpSaver) UploadFile(filePath string, remotePath string) error {
	defer cli.client.Quit()
	err := cli.client.ChangeDir(remotePath)
	if err != nil {
		return err
	}
	// 读取文件内容
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 上传文件
	err = cli.client.Stor(filePath, file)
	if err != nil {
		return err
	}
	return nil
}
