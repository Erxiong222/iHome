package handler

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"ihome/service/register/model"
	register "ihome/service/register/proto/register"
	"ihome/service/register/utils"
)

type Register struct{}


func (e *Register) SmsCode(ctx context.Context, req *register.Request, rsp *register.Response) error {
	//写具体业务   uuid   text    mobile
	//验证图片验证码是否输入正确  从redis中获取到存储的图片验证码
	rnd,err := model.GetImgCode(req.Uuid)
	if err != nil {
		rsp.Errno = utils.RECODE_NODATA
		rsp.Errmsg = utils.RecodeText(utils.RECODE_NODATA)
		return err
	}

	//判断输入的图片验证码是否正确
	if req.Text != rnd {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		//返回自定义的error数据
		return errors.New("验证码输入错误")
	}
	//如果成功,发送短信,存储短信验证码
	//阿里云的短信签名和模板没申请下来。。。。
	/*client, err := sdk.NewClientWithAccessKey("default", "LTAI4FexwrAFbn4ua4DHAyXh", "AltI2inQ1I5TqAEwAfrJNgP54VnVOx")
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	获取6位数随机码
	myRnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", myRnd.Int31n(1000000))*/
	vcode := "123456"

	//存储短信验证码  存redis中
	err = model.SaveSmsCode(req.Mobile,vcode)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}

//注册
func (e*Register)Register(ctx context.Context, req *register.RegRequest, rsp*register.RegResponse) error{
	//存储用户数据到Mysql上
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	//要把sha256得到的数据转换之后存储  转换16进制的
	pwdHash := fmt.Sprintf("%x",pwd_hash)


	err := model.SaveUser(req.Mobile,pwdHash)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}

//登录
func (e*Register)Login(ctx context.Context, req*register.RegRequest, rsp*register.RegResponse) error{
	//查询输入手机号和密码是否正确  mysql
	//给密码加密
	pwdByte := sha256.Sum256([]byte(req.Password))
	pwd_hash := string(pwdByte[:])
	//要把sha256得到的数据转换之后存储  转换16进制的
	pwdHash := fmt.Sprintf("%x",pwd_hash)

	user,err := model.CheckUser(req.Mobile,pwdHash)
	if err != nil {
		rsp.Errno = utils.RECODE_LOGINERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_LOGINERR)
		return err
	}

	//查询成功  登录成功  把用户名传给web端
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Name = user.Name

	return nil
}