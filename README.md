# 使用说明

## 编译

```sh
go build main.go
```

## 使用

```sh
upload2ftp upload -s FTP服务器域名或IP -P FTP服务器端口 -u FTP用户名 -p FTP密码 需要上传的文件或目录
```

参数解释：

> -P FTP服务器端口，默认是21
>
> -d 远程目录，默认是/，这个是ftp的根而非系统的根目录
>
> 除了这两参数可不传，其他的必须传

如果需要上传多个目录或文件，则在命令后接多个目录即可，会将这些目录或文件打包成一个压缩文件，文件名格式为`2006-01-02-15-04-05.zip`

如果只有一个文件，压缩后的文件名为`原文件名.zip`

如果需要指定上传的远程ftp目录，可以加上-d参数，如：

```shell
upload2ftp upload -s FTP服务器域名或IP -P FTP服务器端口 -u FTP用户名 -p FTP密码 -d /test/  /test/test.txt
```