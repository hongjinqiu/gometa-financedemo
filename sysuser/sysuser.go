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
	modelRenderVO := c.SaveCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c SysUser) DeleteData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}

	modelRenderVO := c.DeleteDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c SysUser) EditData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.EditDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c SysUser) NewData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RNewDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c SysUser) GetData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.GetDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 复制
 */
func (c SysUser) CopyData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.CopyDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 放弃保存,回到浏览状态
 */
func (c SysUser) GiveUpData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.GiveUpDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

/**
 * 刷新
 */
func (c SysUser) RefreshData(w http.ResponseWriter, r *http.Request) {
	c.RActionSupport = SysUserSupport{}
	modelRenderVO := c.RefreshDataCommon(w, r)
	c.RenderCommon(w, r, modelRenderVO)
}

func (c SysUser) LogList(w http.ResponseWriter, r *http.Request) {
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
