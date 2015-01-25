package paypact

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type PayPactSupport struct {
	ActionSupport
}

type PayPact struct {
	BaseDataAction
}

func (c PayPact) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PayPact) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PayPact) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PayPact) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PayPact) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c PayPact) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c PayPact) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c PayPact) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = PayPactSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c PayPact) LogList(w http.ResponseWriter, r *http.Request) {
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
