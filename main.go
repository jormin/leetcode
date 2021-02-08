package main

import (
	"gitlab.wcxst.com/jormin/leetcode/config"
	"gitlab.wcxst.com/jormin/leetcode/ctl"
)

// 初始化
func init() {
	config.Init()
}

func main() {
	ctl.SubmitFetchJob(
		&ctl.FetchJob{
			Tag:   ctl.TagProblemList,
			Title: "问题列表",
			Url:   "https://leetcode-cn.com/problemset/all/",
		},
	)
	ctl.Run()
}
