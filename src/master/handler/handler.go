package handler

import (
	"log"

	"fly"

	"html/template"
	"master/api"
	"master/data/user"

	"bytes"
	"fmt"
	"io/ioutil"
	"master/data"
	"master/util"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var PublicPath string

func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	PublicPath = dir + "/../public"
}

func DeleteHandler(ctx *fly.Context) {
	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	data.CleanReportCacheByUid(u1.Uid)

	ctx.Redirect(302, "/show")
}

func CreatTaskPostHandler(ctx *fly.Context) {
	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	ctx.Request.ParseForm()
	urlstr := strings.TrimSpace(ctx.Request.FormValue("url"))

	urls := strings.Split(urlstr, ";")
	for i, v := range urls {
		urls[i] = strings.TrimSpace(v)
	}

	if _, err := url.ParseRequestURI(urls[0]); err != nil {
		ctx.WriteString(200, "url参数错误1")
		return
	}

	num := ctx.Request.FormValue("num")
	num1, e := strconv.Atoi(num)
	if e != nil || num1 == 0 {
		ctx.WriteString(200, "num参数错误")
		return
	}

	conc := ctx.Request.FormValue("conc")
	conc1, e := strconv.Atoi(conc)
	if e != nil || conc1 == 0 {
		ctx.WriteString(200, "conc 参数错误")
		return
	}

	qps := ctx.Request.FormValue("qps")
	qps1, e := strconv.Atoi(qps)
	if e != nil {
		ctx.WriteString(200, "qps参数错误")
		return
	}

	body := ctx.Request.FormValue("body")
	method := ctx.Request.FormValue("method")

	user := ctx.Request.FormValue("user")
	pwd := ctx.Request.FormValue("pwd")

	timeout := ctx.Request.FormValue("timeout")
	timeout1, e := strconv.Atoi(timeout)
	if e != nil {
		ctx.WriteString(200, "timeout参数错误")
		return
	}

	accept := ctx.Request.FormValue("accept")
	contentType := ctx.Request.FormValue("contentType")
	keepAlive1 := ctx.Request.FormValue("keepalive")
	keepAlive := true
	if keepAlive1 == "false" {
		keepAlive = false
	}

	host := ctx.Request.FormValue("host")

	proxyaddr := ctx.Request.FormValue("proxyaddr")

	if proxyaddr != "" {
		if !strings.HasPrefix(proxyaddr, "http://") {
			ctx.WriteString(200, "proxy参数错误")
			return
		}

		if len(proxyaddr) <= 7 {
			ctx.WriteString(200, "proxy参数错误")
			return
		}
	}

	headerstr := ctx.Request.FormValue("headers")
	var h []data.HeaderKV

	if ctx.Request.FormValue("gzip") == "on" {
		h = append(h, data.HeaderKV{K: "Accept-Encoding", V: "gzip"})
	}

	if headerstr == "" {
		h = make([]data.HeaderKV, 0)
	} else {
		h1, err := url.ParseQuery(headerstr)
		if err != nil {
			ctx.WriteString(200, "header错误")
			return
		}

		for k, v := range h1 {
			if len(v) <= 0 {
				continue
			}
			h = append(h, data.HeaderKV{K: k, V: v[0]})
		}
	}

	cookie := ctx.Request.FormValue("cookie")
	if cookie != "" {
		h = append(h, data.HeaderKV{K: "Cookie", V: cookie})
	}

	_, e = url.Parse(proxyaddr)
	if e != nil && proxyaddr != "" {
		ctx.WriteString(200, "proxy参数错误")
		return
	}

	bargs := data.BoomArgs{
		Url:         urlstr,
		Num:         int32(num1),
		Qps:         int32(qps1),
		Conc:        int32(conc1),
		TimeOut:     int32(timeout1),
		UserName:    user,
		Password:    pwd,
		Body:        body,
		Accept:      accept,
		ContentType: contentType,
		Method:      method,
		Headers:     h,
		KeepAlive:   keepAlive,
		ProxyAddr:   proxyaddr,
		Host:        host,
	}

	taskid, result := api.CreateTask(&bargs, u1.Uid)
	if result {
		ctx.WriteString(200, fmt.Sprintf("%s 创建成功 id: %d ", urlstr, taskid))
	} else {
		ctx.WriteString(200, "无法创建任务")
	}
}

func Index(ctx *fly.Context) {
	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	format := `<a href="/logout">注销[%s]</a>`
	renderTemplate(ctx, "create.html", fmt.Sprintf(format, u1.Name))
	return
}

func StaticHandler(ctx *fly.Context) {
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir(PublicPath)))
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}

func ShowStats(ctx *fly.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.Writer.Write(api.GetStatus())
	return
}

func ShowChart(ctx *fly.Context) {
	taskid := ctx.Query("taskid")

	path := filepath.Join(PublicPath, "./chart.html")
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		ctx.WriteString(200, "error tmpl")
		return
	}
	err = t.Execute(ctx.Writer, api.GetChartData(taskid))
	if err != nil {
		log.Println(err)
	}
}

func ShowResult(ctx *fly.Context) {
	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	path := filepath.Join(PublicPath, "./resultTMPL.html")
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		ctx.WriteString(200, "error tmpl")
		return
	}

	userstr := fmt.Sprintf(`<a href="/logout">注销[%s]</a>`, u1.Name)

	m := map[string]interface{}{
		"html": template.HTML(api.GetReport(u1.Uid)),
		"user": template.HTML(userstr),
	}

	err = t.Execute(ctx.Writer, m)
	if err != nil {
		log.Println(err)
	}
}

func GetQueueHandler(ctx *fly.Context) {
	path := filepath.Join(PublicPath, "./queueTMPL.html")
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		ctx.WriteString(200, "error tmpl")
		return
	}

	err = t.Execute(ctx.Writer, template.HTML(api.GetTaskQueueStatus()))
	if err != nil {
		log.Println(err)
	}
	return
}

func UploadfileGet(ctx *fly.Context) {
	path := filepath.Join(PublicPath, "./upload.html")
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
	}
	t.Execute(ctx.Writer, template.HTML(""))
}

func Uploadfile(ctx *fly.Context) {

	data.UploadLock.Lock()
	defer data.UploadLock.Unlock()

	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	fakeFileName := "fake_" + u1.Uid + ".txt"

	fakeFileName = "file1.txt"

	ctx.Request.ParseForm()
	err := ctx.Request.ParseMultipartForm(100000)
	if err != nil {
		ctx.WriteString(200, "file too big")
		return
	}

	b := bytes.NewBufferString("")

	if savefile(ctx.Request, "file1", fakeFileName) {
		b.WriteString(fmt.Sprintf(`<span class="label label-success">%s 成功</span>`, fakeFileName))
		data.FakeQueue <- fakeFileName
	} else {
		b.WriteString(fmt.Sprintf(`<span class="label label-danger">%s 失败</span>`, fakeFileName))
	}

	path := filepath.Join(PublicPath, "./upload.html")
	t, err := template.ParseFiles(path)
	t.Execute(ctx.Writer, template.HTML(b.String()))
}

func savefile(req *http.Request, uploadname, filename string) bool {
	infile, _, err := req.FormFile(uploadname)
	if err != nil {
		log.Println(filename, err)
		return false
	}
	defer infile.Close()

	b, err := ioutil.ReadAll(infile)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	err = ioutil.WriteFile(dir+"/../public/fake/"+filename, b, 0644)

	return err == nil
}

func DeleteTask(ctx *fly.Context) {
	taskid := ctx.Query("taskid")

	data.TaskGlobalLock.Lock()

	id, _ := strconv.Atoi(taskid)
	if data.CurrTaskId == int64(id) {
		api.StopCurrent()
	} else {
		data.TaskQueue[int64(id)] = nil
	}

	data.TaskGlobalLock.Unlock()

	ctx.Redirect(302, "/taskqueue")

}

func DeleteReport(ctx *fly.Context) {
	u, _ := ctx.Get("user")
	u1 := u.(*user.User)

	taskid := ctx.Query("taskid")
	id, _ := strconv.Atoi(taskid)

	r, err := data.Get(int64(id))

	log.Println(r.Uid, u1.Uid)

	if err == nil && r.Uid == u1.Uid {
		data.DelReport(int64(id))
	}

	ctx.Redirect(302, "/show")
}

func LoginGet(ctx *fly.Context) {
	path := filepath.Join(PublicPath, "./login.html")
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
	}
	t.Execute(ctx.Writer, template.HTML(""))
}

func Logout(ctx *fly.Context) {
	ctx.SetCookie(data.Token, "", 0, "/", "", false, true)
	ctx.SetCookie(data.Uid, "", 0, "/", "", false, true)
	ctx.Redirect(302, "/")
}

func LoginPost(ctx *fly.Context) {

	ctx.Request.ParseForm()
	name := ctx.Request.FormValue("username")
	pwd := ctx.Request.FormValue("password")

	u, ok := user.GetUserByName(name)
	if !ok || u.Password != pwd {
		ctx.Redirect(302, "/login")
		return
	}

	ctx.SetCookie(data.Token, util.Hash(name, pwd), 60*60*24, "/", "", false, true)
	ctx.SetCookie(data.Uid, u.Uid, 60*60*24, "/", "", false, true)

	log.Println(name, pwd)

	ctx.Redirect(302, "/")
}

func Register(ctx *fly.Context) {
	ctx.Request.ParseForm()
	name := ctx.Request.FormValue("username")
	pwd := ctx.Request.FormValue("password")

	log.Println(name, pwd)

	_, ok := user.GetUserByName(name)
	if ok {
		ctx.WriteString(200, "用户名重复")
		return
	}

	u := user.AddUserWithId(&user.User{Name: name, Password: pwd})
	ctx.SetCookie(data.Token, util.Hash(name, pwd), 60*60*24, "/", "", false, true)
	ctx.SetCookie(data.Uid, u.Uid, 60*60*24, "/", "", false, true)

	ctx.Redirect(302, "/")
}

func renderTemplate(ctx *fly.Context, html, content string) {
	path := filepath.Join(PublicPath, html)
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		ctx.WriteString(200, "error tmpl")
		return
	}
	err = t.Execute(ctx.Writer, template.HTML(content))
	if err != nil {
		log.Println(err)
	}
}

func ARCH(ctx *fly.Context) {
	buf := bytes.NewBufferString("")
	user.UserCache.ForeachRead(func(k int64, v interface{}) {
		u := v.(*user.User)
		buf.WriteString(u.Uid + " " + u.Name + " " + u.Password + "\n")
	})

	ctx.WriteString(200, buf.String())
}
