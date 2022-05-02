package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {
	// 引数から画像ファイルのURLを取得
	f := flag.String("url", "http://examplea.com/", "URL")
	flag.Parse()

	// 画像ファイルを取得
	image, err := http.Get(*f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 画像ファイルのデータを全て読み込み
	bytes, err := ioutil.ReadAll(image.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// セッション開始
	sess := session.Must(session.NewSession())

	// Rekognitionクライアントを作成
	svc := rekognition.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	// DetectFacesに渡すパラメータを設定
	params := &rekognition.DetectFacesInput{
		Image: &rekognition.Image{
			Bytes: bytes,
		},
		Attributes: []*string{
			aws.String("ALL"),
		},
	}

	// DetectFacesを実行
	resp, err := svc.DetectFaces(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 結果を出力
	fmt.Println(resp)
}
