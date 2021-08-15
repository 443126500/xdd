package models

import (
	"io"
	"os"
	"regexp"
	"runtime"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

var version = "2099091909"
var AppName = "xdd"
var pname = regexp.MustCompile(`/([^/\s]+)`).FindStringSubmatch(os.Args[0])[1]

func initVersion() {
	if Config.Version != "" {
		version = Config.Version
	}
	logs.Info("检查更新" + version)
	value, err := httplib.Get(GhProxy + "https://raw.githubusercontent.com/cdle/jd_study/main/xdd/models/version.go").String()
	if err != nil {
		logs.Info("更新User-Agent失败")
	} else {
		name := AppName + "_" + runtime.GOOS + "_" + runtime.GOARCH
	}
}
