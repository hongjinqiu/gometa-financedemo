package articletype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type ArticleTypeSupport struct {
	ActionSupport
}

type ArticleType struct {
	BaseDataAction
}

func (c ArticleType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c ArticleType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c ArticleType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c ArticleType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c ArticleType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c ArticleType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ArticleType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c ArticleType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c ArticleType) LogList(w http.ResponseWriter, r *http.Request) {
	result := c.RLogListCommon(w, r)

	format := r.FormValue("format")
	if strings.ToLower(format) == "json" {
		w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
		data, err := json.Marshal(&result)
		if err != nil {
			panic(err)
		}
		w.Write(data)
		return
	}
	return
}
