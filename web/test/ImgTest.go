package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	cap := captcha.New()

	cap.SetFont("comic.ttf")
	cap.SetSize(128, 64)

	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 128})
	cap.SetBkgColor(color.RGBA{0, 0, 255, 255}, color.RGBA{0, 255, 255, 255})

	http.HandleFunc("/r", func(writer http.ResponseWriter, request *http.Request) {
		image, s := cap.Create(4, captcha.ALL)
		png.Encode(writer, image)
		println(s)
	})
	http.ListenAndServe(":8081", nil)

}
