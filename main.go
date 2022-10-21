package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	// 環境変数の読み込み
	godotenv.Load(".env")

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Server Running!!")
	})

	router.GET("/db-connect-test", func(c *gin.Context) {
		dbHost := os.Getenv("DB_HOST")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName
		_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{"msg": "internal server error"})
			return
		}

		c.String(200, "DB connect success!!")
	})

	router.POST("/", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 最後にファイルを閉じる
		defer file.Close()

		// 画像ファイルのデータを全て読み込み
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// セッション開始
		sess := session.Must(session.NewSession())

		// Rekognitionクライアントを作成
		svc := rekognition.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

		// Rekognitionに渡すパラメータを設定
		params := &rekognition.DetectTextInput{
			Image: &rekognition.Image{
				Bytes: bytes,
			},
		}

		// Rekognitionを実行
		resp, err := svc.DetectText(params)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}

		// 結果を出力
		c.JSON(200, resp)
	})

	router.Run()
}
