package main

import (
	"encoding/json"
	"fmt"
	"gg/z"
	"github.com/AlecAivazis/survey/v2"
	"github.com/lhlyu/logger/v3/color"
)

const tip = `
----------------------------------------------------------------------------------------
|                                                                                      
|      version: 2.1.2                                                                                   
|                                                                                          
|      1. 构建Gin项目                                                                   
|      2. 构建目录尽量在[ $GOPATH/src ]下
|      3. go建议1.12+                                                                  
|      4. 依赖包下载 git clone http://gogs.miz.so/sdyx-basic/common-lib.git            
|      5. 依赖放在[ $GOPATH/src ]下
|
----------------------------------------------------------------------------------------
`

const success = `
----------------------------------------------------------------------------------------
|                                                                                       
|     SUCCESS !!!
|
|     cd %s
|     go run main.go
|
----------------------------------------------------------------------------------------
`

func main() {
	//createTemplate()
	cli()
}

func createTemplate() {
	// 制作模板
	c := z.NewC()
	c.SetDir(`F:\projects\src\aaaaaa\bbbbbb`)
	c.AddExclude(".idea", "deploy_tmp", "beta")
	c.AddMSS("aaaaaa/", "{{.Workdir}}")
	c.AddMSS("bbbbbb", "{{.Project}}")
	c.AddMSS("cccccc", "{{.Author}}")
	c.AddMSS("dddddd", "{{.Namespace}}")

	c.Create()
}

// Cli
func cli() {
	fmt.Println(color.Cyan(tip))
	qs := z.NewQ()
	an := z.NewAnswer()
	if err := survey.Ask(qs, an); err != nil {
		panic(err)
		return
	}
	node := &z.Node{}
	err := json.Unmarshal([]byte(z.T), node)
	if err != nil {
		panic(err)
	}
	node.Build(an)
	fmt.Println(color.Greenf(success, an.Project))
}
