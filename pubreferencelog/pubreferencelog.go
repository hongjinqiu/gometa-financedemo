package pubreferencelog

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type PubReferenceLogSupport struct {
	ActionSupport
}

type PubReferenceLog struct {
	BaseDataAction
}

func (c PubReferenceLog) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PubReferenceLog) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PubReferenceLog) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PubReferenceLog) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PubReferenceLog) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c PubReferenceLog) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PubReferenceLog) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c PubReferenceLog) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PubReferenceLogSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PubReferenceLog) LogList(w http.ResponseWriter, r *http.Request) {
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
