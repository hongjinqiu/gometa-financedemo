package bank

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type BankSupport struct {
	ActionSupport
}

type Bank struct {
	BaseDataAction
}

func (c Bank) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Bank) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Bank) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Bank) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Bank) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c Bank) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c Bank) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c Bank) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = BankSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c Bank) LogList(w http.ResponseWriter, r *http.Request) {
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
