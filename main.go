package main

import (
	"C"
	"fmt"
	"github.com/mholt/archiver/v3"
	"github.com/yeka/zip"
	"io"
	"strings"
)

/**
 *  The github.com/mholt/archiver/v3/rar.go  has been customized by me.
 */

//export IsRarPasswdCorrect
func IsRarPasswdCorrect(passwd string, filename string) int {
	rar := archiver.Rar{}
	rar.Password = passwd
	return rar.IsPasswdCorrect(filename)
}

//export IsZipPasswdCorrect
func IsZipPasswdCorrect(passwd string, filename string) int {
	// 1、使用zip.OpenReader打开zip文件
	archive, err := zip.OpenReader(filename)
	if err != nil {
		return 2
	}
	defer archive.Close()
	for _, f := range archive.File {
		if f.IsEncrypted() {
			f.SetPassword(passwd)
		}
		if f.FileInfo().IsDir() {
			continue
		}
		fileInArchive, err := f.Open()
		if err != nil {
			return 2
		}

		_, err = io.ReadAll(fileInArchive)
		if err != nil {
			//fmt.Println(err)
			if strings.ContainsRune(err.Error(), 't') {
				return 0
			}
			//flate: corrupt input before offset 1  no
			//unexpected EOF  no
			// checksum error  yes
		}
		fileInArchive.Close()
		break
	}
	return 1
}

// -o libUnArchive.lib -buildmode=c-archive
func main() {
	fmt.Println("abc")
	fmt.Println(IsRarPasswdCorrect("824", "1.rar"))
	fmt.Println(IsZipPasswdCorrect("522", "ziptest.zip"))
}
