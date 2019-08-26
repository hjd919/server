package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

type FileDao struct {
	filepath string
}

func NewFileDao(filepath string) (fileDao *FileDao) {
	fileDao = &FileDao{
		filepath,
	}
	return
}

func (fileDao *FileDao) Create() {
	newFile, err = os.Create(fileDao.filepath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}

func (fileDao *FileDao) Copy(dstPath string) {
	// 打开原始文件
	originalFile, err := os.Open(fileDao.filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (fileDao *FileDao) GetInfo() {
	// 如果文件不存在，则返回错误
	fileInfo, err := os.Stat(fileDao.filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

// remove
func (fileDao *FileDao) Remove() {
	// 如果文件不存在，则返回错误
	err := os.Remove(fileDao.filepath)
	if err != nil {
		log.Fatal(err)
	}
}

// exists
func (fileDao *FileDao) Exists() {
	// 文件不存在则返回error
	fileInfo, err := os.Stat(fileDao.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist. File information:")
	log.Println(fileInfo)
}

// 检查读写权限
func (fileDao *FileDao) CheckPermission(permissionType string) {
	switch permissionType {
	case "write":
		file, err := os.OpenFile(fileDao.filepath, os.O_WRONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Println("Error: Write permission denied.")
			}
		}
		file.Close()
	case "read":
		file, err := os.OpenFile(fileDao.filepath, os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Println("Error: Read permission denied.")
			}
		}
		file.Close()
	}
	return
}

// 写文件
func (fileDao *FileDao) Write(content string) {
	// 可写方式打开文件
	file, err := os.OpenFile(
		fileDao.filepath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte(content)
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
	return
}

// 读文件
func (fileDao *FileDao) Read() []byte {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
	return data
	//    2.读取全部
	/* 	data, err := ioutil.ReadAll(file)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Printf("Data as hex: %x\n", data)
	   	fmt.Printf("Data as string: %s\n", data)
	   	fmt.Println("Number of bytes read:", len(data))
		   return data */
	//    1.读取字节
	/* 	// 从文件中读取len(b)字节的文件。
	   	// 返回0字节意味着读取到文件尾了
	   	// 读取到文件会返回io.EOF的error
	   	byteSlice := make([]byte, 16)
	   	bytesRead, err := file.Read(byteSlice)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	log.Printf("Number of bytes read: %d\n", bytesRead)
	   	log.Printf("Data read: %s\n", byteSlice)
	return byteSlice */
}
