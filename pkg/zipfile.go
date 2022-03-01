package pkg

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"os"
)

func ZipFiles(srcFiles []string, compressFileName string) error {
	_, err := os.Stat(compressFileName)
	if err == nil {
		_ = os.Remove(compressFileName)
	}
	z := archiver.Zip{
		OverwriteExisting: true,
		MkdirAll:          true,
	}
	err = z.Archive(srcFiles, compressFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
