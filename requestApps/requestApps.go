package requestApps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type structCategory struct {
	id   int
	name string
}

type structAppList struct {
	Total   int
	Count   int
	Obj     []StructAppInfo
	Msg     string
	Success bool
}

type structAppRate struct {
	AverageRating float32
	RatingCount   int
}

type StructAppInfo struct {
	ApkMd5         string
	ApkURL         string
	ApkPublishTime int
	AppDownCount   int
	AppName        string
	AppRatingInfo  structAppRate
	AuthorName     string
	AverageRating  float32
	CategoryID     int
	CategoryName   string
	Description    string
	EditorIntro    string
	FileSize       int
	PkgName        string
	VersionCode    int
	VersionName    string
	RankId         int
}

// var category = map[string]structCategory{
// 	// "read": structCategory{
// 	// 	102, "阅读",
// 	// },
// 	// "news": structCategory{
// 	// 	110, "新闻",
// 	// },
// 	// "travel": structCategory{
// 	// 	108, "旅游",
// 	// },
// 	// "tool": structCategory{
// 	// 	115, "工具",
// 	// },
// 	"video": structCategory{
// 		103, "视频",
// 	},
// 	"communication": structCategory{
// 		106, "社交",
// 	},
// 	"music": structCategory{
// 		101, "音乐",
// 	},
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getHTTPData(url string) (res []byte) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return body
}

func GetAppList(categoryID int, pageSize int, pageContext int) []StructAppInfo {
	url := fmt.Sprintf("http://sj.qq.com/myapp/cate/appList.htm?orgame=1&categoryId=%d&pageSize=%d&pageContext=%d", categoryID, pageSize, pageContext)
	fmt.Println(url)
	appData := getHTTPData(url)
	var apps structAppList
	json.Unmarshal(appData, &apps)
	if apps.Success {
		return apps.Obj
	} else {
		// checkErr(errors.New("fetch url error"))
	}
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("try getAppList again")
	return GetAppList(categoryID, pageSize, pageContext)
}

// func main() {
// 	for _, v := range category {
// 		categoryId := v.id
// 		pageSize := 10
// 		for i := 0; i < 1; i++ {
// 			pageContext := i * pageSize
// 			apps := getAppList(categoryId, pageSize, pageContext)
// 			fmt.Println(v.name, apps)
// 		}
// 	}
// 	// appUrl = getAppURL()
// }
