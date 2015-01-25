package providertype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type ProviderTypeSupport struct {
	ActionSupport
}

type ProviderType struct {
	BaseDataAction
}

func (c ProviderType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ProviderType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ProviderType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ProviderType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ProviderType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c ProviderType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c ProviderType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c ProviderType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = ProviderTypeSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c ProviderType) LogList(w http.ResponseWriter, r *http.Request) {
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
