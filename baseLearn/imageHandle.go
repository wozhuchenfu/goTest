package baseLearn

import (
	"os"
	"fmt"
	"image/png"
	"image/jpeg"
	"image"
	"image/draw"
	"io"
	"net/http"
)

func HandleImage()  {
	f,err := os.Open("C:\\Users\\edianzu\\Desktop\\test.png")
	if err!=nil {
		fmt.Println(err.Error())
	}
	png,err1 := png.Decode(f)
	if err1 != nil{
		fmt.Println(err1)
	}
	defer f.Close()

	f2,err2 := os.Open("C:\\Users\\edianzu\\Desktop\\test2.jpg")
	if err2 !=nil {
		fmt.Println(err2)
	}
	jpg,err2 := jpeg.Decode(f2)
	if err2!=nil {
		fmt.Println(err2)
	}
	defer f2.Close()

	offset := image.Pt(200,200)
	b:=png.Bounds()
	m:=image.NewRGBA(b)
	draw.Draw(m,b,png,image.ZP,draw.Src)
	draw.Draw(m,jpg.Bounds().Add(offset),jpg,image.ZP,draw.Over)

	imgw ,_:= os.Create("C:\\Users\\edianzu\\Desktop\\result.jpg")
	jpeg.Encode(imgw,m,&jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()

}

func Test(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	imgb, _ := os.Open("C:\\Users\\edianzu\\Desktop\\test2.jpg")
	b := make([]byte,1024)
	for {
		_,err := imgb.Read(b)
		if err == io.EOF {
			break
		}
	}
	fmt.Fprint(w,b)
	defer r.Body.Close()
	defer imgb.Close()

	/*img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("C:\\Users\\edianzu\\Desktop\\test.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	offset := image.Pt(200, 200)
	b := img.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create("C:\\Users\\edianzu\\Desktop\\result.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()*/

}

func Test2()  {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("C:\\Users\\edianzu\\Desktop\\test2.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	img, _ := jpeg.Decode(file1)

	file2, err := os.Open("C:\\Users\\edianzu\\Desktop\\test3.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()
	img2, _ := jpeg.Decode(file2)

	jpg := image.NewRGBA(image.Rect(0, 0, 300, 300))

	draw.Draw(jpg, jpg.Bounds(), img2, img2.Bounds().Min, draw.Over)                   //首先将一个图片信息存入jpg
	draw.Draw(jpg, jpg.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over)   //将另外一张图片信息存入jpg

	// draw.DrawMask(jpg, jpg.Bounds(), img, img.Bounds().Min, img2, img2.Bounds().Min, draw.Src) // 利用这种方法不能够将两个图片直接合成？目前尚不知道原因。

	jpeg.Encode(file, jpg, nil)
}

