package taxtype

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type TaxTypeSupport struct {
	ActionSupport
}

type TaxType struct {
	BaseDataAction
}

func (c TaxType) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c TaxType) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c TaxType) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c TaxType) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c TaxType) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c TaxType) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c TaxType) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c TaxType) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = TaxTypeSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c TaxType) LogList(w http.ResponseWriter, r *http.Request) {
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
