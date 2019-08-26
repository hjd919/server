package file

import (
	"log"
	"testing"
)

// add
func TestCreateFile(t *testing.T) {
	fileDao := NewFileDao("./test.txt")
	fileDao.Create()
}

// query
func TestGetFileInfo(t *testing.T) {
	fileDao := NewFileDao("./test.txt")
	fileDao.GetInfo()
}

// remove
func TestRemoveFile(t *testing.T) {
	fileDao := NewFileDao("./test.txt")
	fileDao.Remove()
}

// 判断存在
func TestExistsFile(t *testing.T) {
	fileDao := NewFileDao("./test.txt")
	fileDao.Exists()
}

// 判断复制
func TestCopyFile(t *testing.T) {
	fileDao := NewFileDao("./test.txt")
	fileDao.Copy("./test_copy.txt")
}

// 写文件
func TestWriteFile(t *testing.T) {
	fileDao := NewFileDao("./test2.txt")
	fileDao.Write("hjd hjd")
}

// 读文件
func TestReadFile(t *testing.T) {
	fileDao := NewFileDao("./test2.txt")
	byteContent := fileDao.Read()
	log.Printf("%#v", string(byteContent))
}

// 读取csv
// 压缩解压缩zip
