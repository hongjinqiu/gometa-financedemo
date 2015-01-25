package provider

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type ProviderSupport struct {
	ActionSupport
}

type Provider struct {
	BaseDataAction
}

func (c Provider) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Provider) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Provider) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Provider) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Provider) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c Provider) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Provider) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c Provider) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Provider) LogList(w http.ResponseWriter, r *http.Request) {
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
