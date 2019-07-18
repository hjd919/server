package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/gin-gonic/gin/binding"
)

func (s *Service) Record(c *bm.Context) {
	var req struct {
		Filelink string `json:"file_link" form:"file_link" binding:"required"`
	}
	if err := c.BindWith(&req, binding.Form); err != nil {
		c.JSON("", err)
		return
	}

	res, err := s.do(req.Filelink)
	if err != nil {
		c.JSON(err.Error(), err)
		return
	}

	c.JSON(res, nil)
}

func (s *Service) do(fileLink string) (res interface{}, err error) {
	// fileLink, _ = url.QueryUnescape(fileLink)
	// fileLink = url.QueryEscape(fileLink)
	log.Info("fileLink----%s", fileLink)
	// 地域ID，常量内容，请勿改变
	const REGION_ID string = "cn-shanghai"
	const ENDPOINT_NAME string = "cn-shanghai"
	const PRODUCT string = "nls-filetrans"
	const DOMAIN string = "filetrans.cn-shanghai.aliyuncs.com"
	const API_VERSION string = "2018-08-17"
	const POST_REQUEST_ACTION string = "SubmitTask"
	const GET_REQUEST_ACTION string = "GetTaskResult"
	// 请求参数key
	const KEY_APP_KEY string = "appkey"
	const KEY_FILE_LINK string = "file_link"
	const KEY_VERSION string = "version"
	const KEY_ENABLE_WORDS string = "enable_words"
	// 响应参数key
	const KEY_TASK string = "Task"
	const KEY_TASK_ID string = "TaskId"
	const KEY_STATUS_TEXT string = "StatusText"
	const KEY_RESULT string = "Result"
	// 状态值
	const STATUS_SUCCESS string = "SUCCESS"
	const STATUS_RUNNING string = "RUNNING"
	const STATUS_QUEUEING string = "QUEUEING"
	var accessKeyId string = "A07OENwppI9Tfa1B"
	var accessKeySecret string = "1Lh7iVcsM2GlLxExwsFV2cTPPYQXly"
	var appKey string = "wpuZcw0HrDUkV6PD"

	// var fileLink string = "http://xz-yzzc.oss-cn-beijing.aliyuncs.com/audio/%E5%B7%A5%E4%BD%9C5.m4a"
	// fileLink = "https://aliyun-nls.oss-cn-hangzhou.aliyuncs.com/asr/fileASR/examples/nls-sample-16k.wav"
	// fileLink = "http://xz-yzzc.oss-cn-beijing.aliyuncs.com/audio/%E9%93%83%E5%A3%B0_01.mp3"

	client, err := sdk.NewClientWithAccessKey(REGION_ID, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	postRequest := requests.NewCommonRequest()
	postRequest.Domain = DOMAIN
	postRequest.Version = API_VERSION
	postRequest.Product = PRODUCT
	postRequest.ApiName = POST_REQUEST_ACTION
	postRequest.Method = "POST"
	mapTask := make(map[string]string)
	mapTask[KEY_APP_KEY] = appKey
	mapTask[KEY_FILE_LINK] = fileLink
	// 新接入请使用4.0版本，已接入(默认2.0)如需维持现状，请注释掉该参数设置
	mapTask[KEY_VERSION] = "4.0"
	// 设置是否输出词信息，默认为false，开启时需要设置version为4.0
	mapTask[KEY_ENABLE_WORDS] = "false"
	task, err := json.Marshal(mapTask)
	if err != nil {
		panic(err)
	}
	postRequest.FormParams[KEY_TASK] = string(task)
	postResponse, err := client.ProcessCommonRequest(postRequest)
	if err != nil {
		panic(err)
	}
	postResponseContent := postResponse.GetHttpContentString()
	log.Info(postResponseContent)
	if postResponse.GetHttpStatus() != 200 {
		log.Info("录音文件识别请求失败，Http错误码: ", postResponse.GetHttpStatus())
		return
	}
	var postMapResult map[string]interface{}
	err = json.Unmarshal([]byte(postResponseContent), &postMapResult)
	if err != nil {
		panic(err)
	}
	var taskId string = ""
	var statusText string = ""
	statusText = postMapResult[KEY_STATUS_TEXT].(string)
	if statusText == STATUS_SUCCESS {
		log.Info("录音文件识别请求成功响应!")
		taskId = postMapResult[KEY_TASK_ID].(string)
	} else {
		log.Info("录音文件识别请求失败!")
		return
	}
	getRequest := requests.NewCommonRequest()
	getRequest.Domain = DOMAIN
	getRequest.Version = API_VERSION
	getRequest.Product = PRODUCT
	getRequest.ApiName = GET_REQUEST_ACTION
	getRequest.Method = "GET"
	getRequest.QueryParams[KEY_TASK_ID] = taskId
	statusText = ""
	var getMapResult map[string]interface{}
	for true {
		getResponse, err := client.ProcessCommonRequest(getRequest)
		if err != nil {
			panic(err)
		}
		getResponseContent := getResponse.GetHttpContentString()
		log.Info("识别查询结果：", getResponseContent)
		if getResponse.GetHttpStatus() != 200 {
			log.Info("识别结果查询请求失败，Http错误码：", getResponse.GetHttpStatus())
			break
		}
		err = json.Unmarshal([]byte(getResponseContent), &getMapResult)
		if err != nil {
			panic(err)
		}
		statusText = getMapResult[KEY_STATUS_TEXT].(string)
		if statusText == STATUS_RUNNING || statusText == STATUS_QUEUEING {
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	if statusText != STATUS_SUCCESS {
		log.Info("录音文件识别成功！")
		return "", fmt.Errorf("录音文件识别失败！--文件:%v", fileLink)
	}

	return getMapResult["Result"], nil
}
