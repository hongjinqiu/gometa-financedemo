package systemparameter

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type SystemParameterSupport struct {
	ActionSupport
}

type SystemParameter struct {
	BaseDataAction
}

func (c SystemParameter) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SystemParameter) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SystemParameter) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SystemParameter) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SystemParameter) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c SystemParameter) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SystemParameter) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c SystemParameter) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SystemParameterSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SystemParameter) LogList(w http.ResponseWriter, r *http.Request) {
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
