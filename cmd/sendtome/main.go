package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AnnonaOrg/sendtome/internal/log"

	"github.com/AnnonaOrg/osenv"
	"github.com/gin-gonic/gin"
	"github.com/AnnonaOrg/sendtome/internal/constvar"
	_ "github.com/AnnonaOrg/sendtome/internal/dotenv"
	"github.com/AnnonaOrg/sendtome/router"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("run time panic:%v\n", err)
		}
	}()

	fmt.Printf("%s %s\n%s\n",
		constvar.APPName(), constvar.APPVersion(), constvar.APPDesc(),
	)
	runAt := "运行在"
	if osenv.IsInDocker() {
		runAt = runAt + "(Docker)"
	}
	runAt = runAt + ": " + osenv.Getwd()
	fmt.Println(runAt)
	time.Sleep(time.Second * 3)

	// Set gin mode.
	ginMode := osenv.GetServerGinRunmode()
	gin.SetMode(ginMode)
	//Create the Gin engine.
	g := gin.New()
	//Routes.
	router.Load(
		g,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatalf(
				"(%s)没有响应，请检查配置及网络状态: %v",
				constvar.APPName(), err,
			)
		}
		log.Infof("(%s)成功部署，服务地址:%s", constvar.APPName(), osenv.GetServerUrl())
	}()

	go mainTask()

	addr := ":" + osenv.GetServerPort()
	if err := http.ListenAndServe(addr, g); err != nil {
		log.Errorf(
			"(%s)出错了，需要重启: %v",
			constvar.APPName(), err,
		)
	}
}
