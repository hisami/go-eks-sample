package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {
	// // 引数から画像ファイルのURLを取得
	// f := flag.String("url", "http://examplea.com/", "URL")
	// flag.Parse()

	// // 画像ファイルを取得
	// image, err := http.Get(*f)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	file, err := os.Open("img/montes.png")
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

	// DetectFacesに渡すパラメータを設定
	params := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: bytes,
		},
	}

	// DetectFacesを実行
	resp, err := svc.DetectText(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 結果を出力
	fmt.Println(resp)
}
