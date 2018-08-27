package analyseApp

import (
	"archive/zip"
	"regexp"
	"strings"
)

var LibSO = map[string]string{
	"libyoume":  "游密云",
	"libqav":    "腾讯云",
	"libycloud": "欢聚云",
	"libqqlive": "腾讯云",
	"libqy":     "爱奇艺",
}

func AnalyseApp(zipFile string) map[string]bool {
	soMap := make(map[string]bool)
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return soMap
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return soMap
		}
		defer rc.Close()
		reg := regexp.MustCompile("\\.so$")
		if reg.Match([]byte(file.Name)) {
			arr := strings.Split(file.Name, "/")
			so := arr[len(arr)-1]
			so = strings.ToLower(so)
			for k, v := range LibSO {
				if strings.Contains(so, k) {
					soMap[v] = true
				}
			}
		}
	}
	return soMap
}

// func main() {
// 	filePath := "./packages/os.xiehou360.im.mei_6.2.0_620.zip"
// 	soList := analyseApp(filePath)
// 	fmt.Println(soList)
// }
