package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type UploadController struct {
	Ctx iris.Context
}

func (p *UploadController) PostImages() {
	result := struct {
		Code  string `json:"code"`
		Url   string `json:"url"`
		Msg   string `json:"msg"`
		Thumb string `json:"thumb"`
		Type  string `json:"type"`
		Id    string `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
	}{"", "", "", "", "", "", "", ""}

	file, info, err := p.Ctx.FormFile("file")
	if err != nil {
		golog.Error(err)
		p.Ctx.JSON(result)
	}

	dir := "/upload/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		golog.Error(err)
		p.Ctx.JSON(result)
	}
	newUuid := uuid.New().String()
	url := dir + newUuid + strings.ToLower(path.Ext(info.Filename))
	out, err := os.OpenFile("./public/"+url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		golog.Error(err)
		p.Ctx.JSON(result)
	}

	defer file.Close()
	defer out.Close()
	io.Copy(out, file)

	result.Id = "1"
	result.Code = "1"
	result.Link = url
	result.Url = url
	result.Thumb = url
	result.Msg = url
	result.Type = "image"
	result.Name = newUuid
	p.Ctx.JSON(result)
}

func (p *UploadController) PostVideos() {
	result := struct {
		Code  string `json:"code"`
		Url   string `json:"url"`
		Msg   string `json:"msg"`
		Thumb string `json:"thumb"`
		Type  string `json:"type"`
		Id    string `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
	}{"", "", "", "", "", "", "", ""}

	file, info, err := p.Ctx.FormFile("file")
	if err != nil {
		golog.Error(err)
		_, _ = p.Ctx.JSON(result)
	}

	dir := "/upload/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		golog.Error(err)
		_, _ = p.Ctx.JSON(result)
	}
	newUuid := uuid.New().String()
	url := dir + newUuid + strings.ToLower(path.Ext(info.Filename))
	out, err := os.OpenFile("./public/"+url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		golog.Error(err)
		_, _ = p.Ctx.JSON(result)
	}

	defer file.Close()
	defer out.Close()
	_, _ = io.Copy(out, file)

	result.Id = "1"
	result.Code = "1"
	result.Link = url
	result.Url = url
	result.Thumb = url
	result.Msg = url
	result.Type = "video"
	result.Name = newUuid
	p.Ctx.JSON(result)
}
