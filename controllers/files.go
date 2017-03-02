package controllers

import (
	"books/utils"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileController struct {
	AppController
}

func (this *FileController) Post() {
	fileName := this.GetString(":splat")
	fmt.Println("filename: ", fileName)
	if !strings.HasSuffix(fileName, ".jpg") && !strings.HasSuffix(fileName, ".png") && !strings.HasSuffix(fileName, ".jpeg") {
		this.Fail(400, "非法文件,只能上传jpg jpeg png格式图片")
	}
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	t := time.Now()
	path := "upload/" + fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
	exp := filepath.Ext(fileName)
	fileName = utils.EncodePasswd(fileName+fmt.Sprintf("%d", t.Unix())) + exp
	_, err := SaveAsFile(body, path, fileName)
	if err != nil {
		this.Fail(400, "上传失败")
	}
	this.Succuess(map[string]string{"src": path + "/" + fileName})
}

func SaveAsFile(b []byte, path, filename string) (int, error) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return 0, err
	}
	f, err := os.Create(path + "/" + filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		os.Remove(filename)
	}
	return n, err
}

func (this *FileController) TestFile() {
	this.Data["token"] = this.User.Token
	this.TplName = "file.html"
}

func (this *FileController) UploadFileByFrom() {
	req := this.Ctx.Request
	f, h, err := req.FormFile("file")
	if err != nil {
		this.Fail(2, err.Error())
	}
	exp := filepath.Ext(h.Filename)
	fmt.Println("exp: ", exp)
	if exp != ".jpg" && exp != ".jpeg" && exp != ".png" {
		this.Fail(400, "非法文件,只能上传jpg jpeg png格式图片")
	}
	upload := "upload/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day())
	err = os.MkdirAll(upload, 0777)
	if err != nil {
		this.Fail(2, err.Error())
	}
	fs, err := os.Create(upload + "/" + h.Filename)
	io.Copy(fs, f)
	m := map[string]string{"src": upload + "/" + h.Filename}
	this.Succuess(m)
}
