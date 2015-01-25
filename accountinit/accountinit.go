package accountinit

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type AccountInitSupport struct {
	ActionSupport
}

type AccountInit struct {
	BaseDataAction
}

func (c AccountInit) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c AccountInit) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c AccountInit) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c AccountInit) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c AccountInit) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c AccountInit) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c AccountInit) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c AccountInit) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = AccountInitSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c AccountInit) LogList(w http.ResponseWriter, r *http.Request) {
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
