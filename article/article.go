package article

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type ArticleSupport struct {
	ActionSupport
}

type Article struct {
	BaseDataAction
}

func (c Article) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Article) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	
	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Article) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Article) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Article) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c Article) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Article) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c Article) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ArticleSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Article) LogList(w http.ResponseWriter, r *http.Request) {
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

