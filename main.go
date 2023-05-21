package main

import (
	"os"
	"work/config"
	"work/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := config.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := gin.New()
	r.Static("/static", "./static")
	// r.MaxMultipartMemory = 8 << 20  // 上传文件大小 8mb
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)

	}
}
