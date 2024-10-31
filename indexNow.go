package main

import (
	"bytes"
	"os"

	"github.com/axiaoxin-com/urlsubmitter"
	"github.com/mnhkahn/gogogo/logger"
	"github.com/sasbury/mini"
)

const (
	configPath = "./indexNow.conf"
	urlFile    = "./urls.txt"
)

func main() {
	appconfig := new(mini.Config)
	err := appconfig.InitializeFromPath(configPath)
	if err != nil { // Handle errors reading the config file
		logger.Info("no config file", configPath)
		os.Exit(1)
	}

	baiduAPI := appconfig.String("BAIDU_API", "")
	bingKey := appconfig.String("BING_KEY", "")
	bingHost := appconfig.String("BING_SUBMIT_HOST", "")
	bingKeyLocation := bingHost + "/" + bingKey + ".txt"
	googleCredentialsFile := appconfig.String("GOOGLE_CREDENTIALS_FILE", "")
	logger.Info(bingKey, bingHost, bingKeyLocation, baiduAPI, googleCredentialsFile)

	// 初始化 Bing 提交器
	// bingSubmitter := urlsubmitter.NewBingSubmitter(
	// 	bingKey,
	// 	bingKeyLocation,
	// 	bingHost,
	// )
	// 初始化 Baidu 提交器
	// baiduSubmitter := urlsubmitter.NewBaiduSubmitter(baiduAPI)
	// 初始化 Google 提交器
	googleSubmitter := urlsubmitter.NewGoogleSubmitter("/path/to/your-svc-account-keys.json")

	urlBytes, err := os.ReadFile(urlFile)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	} else if len(urlBytes) == 0 {
		logger.Error(urlFile, "is empty")
		os.Exit(1)
	}

	urls := []string{}
	for _, u := range bytes.Split(urlBytes, []byte("\n")) {
		urls = append(urls, string(u))
	}
	logger.Info(urls)

	// // 提交 URL 到 Bing
	// bingResult, err := bingSubmitter.SubmitURLs(urls)
	// if err != nil {
	// 	logger.Error("Error submitting to Bing:", err)
	// 	os.Exit(1)
	// }
	// logger.Info("Bing Result:", bingResult)

	// // 提交 URL 到 Baidu
	// baiduResult, err := baiduSubmitter.SubmitURLs(urls)
	// if err != nil {
	// 	logger.Error("Error submitting to Bing:", err)
	// 	os.Exit(1)
	// }
	// logger.Info("Bing Result:", baiduResult)

	// 提交 URL 到 Google
	googleResult, err := googleSubmitter.SubmitURLs(urls)
	if err != nil {
		logger.Error("Error submitting to Bing:", err)
		os.Exit(1)
	}
	logger.Info("Bing Result:", googleResult)
}
