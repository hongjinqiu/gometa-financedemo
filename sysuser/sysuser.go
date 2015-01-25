package sysuser

import (
	"encoding/json"
	. "github.com/hongjinqiu/gometa/controllers"
	"net/http"
	"strings"
)

func init() {
}

type SysUserSupport struct {
	ActionSupport
}

type SysUser struct {
	BaseDataAction
}

func (c SysUser) SaveData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RSaveCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SysUser) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}

	modelRenderVO := c.RDeleteDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SysUser) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.REditDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SysUser) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SysUser) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RGetDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c SysUser) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RCopyDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SysUser) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RGiveUpDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c SysUser) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RRefreshDataCommon(w, r)
	c.RRenderCommon(w, r, modelRenderVO)
}

func (c SysUser) LogList(w http.ResponseWriter, r *http.Request) {
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
