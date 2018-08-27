package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"./analyseApp"
	"./downloadApp"
	"./requestApps"
)

var SqlLock sync.Mutex //申明一个互斥锁

type structCategory struct {
	id   int
	name string
}

var category = map[string]structCategory{
	"read": structCategory{
		102, "阅读",
	},
	"news": structCategory{
		110, "新闻",
	},
	"video": structCategory{
		103, "视频",
	},
	"tool": structCategory{
		115, "工具",
	},
	"social": structCategory{
		106, "社交",
	},
	"music": structCategory{
		101, "音乐",
	},
	"education": structCategory{
		111, "教育",
	},
	"health": structCategory{
		109, "健康",
	},
	"entertainment": structCategory{
		105, "娱乐",
	},
	"children": structCategory{
		100, "儿童",
	},
	"works": structCategory{
		113, "办公",
	},
	"communication": structCategory{
		116, "通讯",
	},
}

func trackApp(app requestApps.StructAppInfo, db *sql.DB, ch chan string) {
	// 查询数据
	queryString := fmt.Sprintf(`SELECT count(*) as count FROM apps WHERE PkgName = "%s" AND VersionName = "%s"`, app.PkgName, app.VersionName)
	// fmt.Println(queryString)
	// SqlLock.Lock()
	rows, _ := db.Query(queryString)
	count := 0
	if rows != nil {
		for rows.Next() {
			err := rows.Scan(&count)
			if err != nil {
				fmt.Println(err)
			}
			// checkErr(err)
		}
	}
	// SqlLock.Unlock()
	// fmt.Println(rows)
	if count == 0 {
		filePath := downloadApp.DownloadApp(app.ApkURL)
		// fmt.Println(filePath)
		soMap := analyseApp.AnalyseApp(filePath)
		UsedSDK := ""
		for k := range soMap {
			UsedSDK = UsedSDK + " " + k
		}
		// fmt.Println(app.AppName, " downloading...")
		UpdateTime := time.Now().Format("2006/1/2 15:04:05")
		SqlLock.Lock() //写全局数据加互斥锁
		stmt, _ := db.Prepare(`INSERT INTO apps(
				ApkMd5, 
				ApkURL,
				ApkPublishTime,
				AppDownCount,
				AppName,
				AuthorName,
				AverageRating,
				CategoryID,
				CategoryName,
				Description,
				EditorIntro,
				FileSize,
				PkgName,
				VersionCode,
				VersionName,
				UsedSDK,
				RankId,
				UpdateTime
			) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
		// fmt.Println("prepare data")
		// checkErr(err)
		_, err := stmt.Exec(
			app.ApkMd5,
			app.ApkURL,
			app.ApkPublishTime,
			app.AppDownCount,
			app.AppName,
			app.AuthorName,
			app.AverageRating,
			app.CategoryID,
			app.CategoryName,
			app.Description,
			app.EditorIntro,
			app.FileSize,
			app.PkgName,
			app.VersionCode,
			app.VersionName,
			UsedSDK,
			app.RankId,
			UpdateTime)
		// fmt.Println("exec data")
		SqlLock.Unlock() //解锁
		checkErr(err)
	} else {
		fmt.Println(app.AppName, " already exists...")
	}
	ch <- app.AppName
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./app_info.sqlite")
	checkErr(err)

	ch := make(chan string)

	count := 0
	for _, v := range category {
		categoryID := v.id
		pageSize := 60
		for i := 0; i < 1; i++ {
			pageContext := i * pageSize
			apps := requestApps.GetAppList(categoryID, pageSize, pageContext)
			for RankId, app := range apps {
				// fmt.Println(v.name, app.AppName)
				count = count + 1
				app.RankId = RankId
				go trackApp(app, db, ch)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}
	for {
		select {
		case <-ch:
			{
				count = count - 1
				println("No.", count)
				if count <= 0 {
					println("process end exist after 8 seconds ...")
					db.Close()
					time.Sleep(time.Second * 8)
					return
				}
			}
		}
	}
}
