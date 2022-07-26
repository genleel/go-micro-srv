package handler

import (
	"context"
	"fmt"
	"user/model"
	"user/utils"

	"math/rand"
	"time"
	pb "user/proto"
)

type User struct{}

func (e *User) SendSms(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	result := model.CheckImgCode(req.Uuid, req.ImgCode)

	if result {
		// 发送短信
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		smsCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
		templateParam := `{"code":"` + smsCode + `"}`

		fmt.Println(templateParam)

		// 验证码存入redis
		err := model.SaveSmsCode(req.Phone, smsCode)
		if err != nil {
			// 存储验证码失败
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		}

		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	} else {
		// 发送错误信息
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	return nil
}

func (e *User) Register(ctx context.Context, req *pb.RegReq, rsp *pb.CallResponse) error {
	// 校验验证码是否正确
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)
	if err == nil {
		// 注册用户写入数据库
		err := model.RegisterUser(req.Mobile, req.Password)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		}
	} else {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	return nil
}
