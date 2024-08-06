package dotenv

import (
	"fmt"

	"github.com/AnnonaOrg/pkg/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env 配置文件加载失败❌: %v\n", err)
		return
	}
}
