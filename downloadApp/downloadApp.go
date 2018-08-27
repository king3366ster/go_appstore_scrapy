package downloadApp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDir(_dir string) (res bool) {
	exist, err := pathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error!%v\n", err)
		return false
	}
	if exist {
		// fmt.Printf("has dir!%v\n", _dir)
	} else {
		// fmt.Printf("no dir!%v\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed!%v\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	return true
}

func downloadFromURL(url, rootDir string) (filePath string) {
	tokens := strings.Split(url, "?")
	queryString := tokens[1]
	querys := strings.Split(queryString, "&")
	var fileName string
	for _, v := range querys {
		if strings.Contains(v, "fsname=") {
			fileName = strings.Replace(v, "fsname=", "", 1)
			fileName = strings.Replace(fileName, ":", "_", -1)
		}
	}
	filePath = rootDir + fileName

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", filePath, "-", err)
		return ""
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return ""
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return ""
	}

	fmt.Println(filePath, n, "bytes downloaded.")
	return filePath
}

func DownloadApp(apkURL string) (newFilePath string) {
	rootDir := "./packages/"
	createDir(rootDir)
	filePath := downloadFromURL(apkURL, rootDir)
	newFilePath = strings.Replace(filePath, ".apk", ".zip", 1)
	err := os.Rename(filePath, newFilePath)
	if err != nil {
		fmt.Println("Error while rename", filePath, "-", err)
		return ""
	}
	return newFilePath
}

// func main() {
// 	apkURL := "http://imtt.dd.qq.com/16891/840A0CB8193DAF1EC82F55CE5FFD1551.apk?fsname=os.xiehou360.im.mei_6.2.0_620.apk&csr=3554"
// 	filePath := DownloadApp(apkURL)
// 	fmt.Println(filePath)
// }
