package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type Version struct {
	Version      string `json:"version"`      // git 版本
	GitCommit    string `json:"gitCommit"`    // git 提交
	GitTreeState string `json:"gitTreeState"` // git 树状态
	GitBranch    string `json:"gitBranch"`    // git 分支
	BuildDate    string `json:"buildDate"`    // 编译时间
	GoVersion    string `json:"goVersion"`    // go 版本
	Platform     string `json:"platform"`     // 编译平台和架构
}

var (
	gitVersion   = "v0.0.0-master+$Format:%H$" // 标记和跟踪代码库的特定版本
	gitCommit    = "$Format:%H$"               // 提交的完整信息
	gitTreeState = ""                          // git 树的状态
	gitBranch    = ""                          // git 分支

	buildDate = "1970-01-01T00:00:00Z" // 编译时间

	version = &Version{
		Version:      gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		GitBranch:    gitBranch,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
)

// Print() 打印应用的 version
func Print() {
	data, _ := json.MarshalIndent(version, "", "    ")
	fmt.Println(string(data))
}
