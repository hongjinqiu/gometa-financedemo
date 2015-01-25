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
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Provider) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Provider) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Provider) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Provider) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c Provider) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Provider) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c Provider) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c Provider) LogList(w http.ResponseWriter, r *http.Request) {
	result := c.LogListCommon(w, r)

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
