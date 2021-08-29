package main

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
)

func main() {
	cap :=captcha.New()

	//设置字符集
	cap.SetFont("source/comic.ttf")
	//if err := cap.SetFont("comic.ttf"); err != nil {
	//	panic(err.Error())
	//}
	//设置验证码图片大小
	cap.SetSize(128, 64)
	//设置混淆程度
	cap.SetDisturbance(captcha.NORMAL)
	//设置字体颜色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255},color.RGBA{255, 0, 0, 255})
	//设置背景色  background
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//生成验证码图片
	//rand.Seed(time.Now().UnixNano())
	img,rnd := cap.Create(4,captcha.NUM)
	fmt.Println(img, rnd)
}
