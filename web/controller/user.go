package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro-srv/web/proto/getCaptcha"
	"go-micro-srv/web/utils"
	"go-micro.dev/v4"
	"image/png"
	"net/http"
)

// 获取session信息
func GetSession(ctx *gin.Context) {
	// 初始化map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片验证码
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码 uuid
	uuid := ctx.Param("uuid")

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getcaptcha", consulService.Client())

	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务...")
		return
	}

	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到 浏览器.
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
}
