package tree_simple

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"testing"
)

// 定义我们自己的菜单对象
type SystemMenu struct {
	Id       int64  `json:"id"`        //id
	ParentId int64  `json:"parent_id"` //上级菜单id
	Name     string `json:"name"`      //菜单名
	Route    string `json:"route"`     //页面路径
	Icon     string `json:"icon"`      //图标路径
	Children INodes `json:"children"` //子菜单
}

// region 实现ITree 所有接口
func (s SystemMenu) GetId() int64 {
	return s.Id
}

func (s SystemMenu) GetParentId() int64 {
	return s.ParentId
}

func (s SystemMenu) IsRoot() bool {
	// 这里通过ParentId等于0 或者 ParentId等于自身Id表示顶层根节点
	return s.ParentId == 0 || s.ParentId == s.Id
}

func (s *SystemMenu) SetChildren(nodes INodes) {
	s.Children = append(s.Children, nodes...)
}

// endregion

type SystemMenus []*SystemMenu

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s SystemMenus) ConvertToINodeArray() (nodes []INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

func TestGenerateTree(t *testing.T) {
	// 模拟获取数据库中所有菜单，在其它所有的查询中，也是首先将数据库中所有数据查询出来放到数组中，
	// 后面的遍历递归，都在这个 allMenu中进行，而不是在数据库中进行递归查询，减小数据库压力。
	allMenu := []*SystemMenu{
		{Id: 1, ParentId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, ParentId: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},

		{Id: 3, ParentId: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, ParentId: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},

		{Id: 5, ParentId: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, ParentId: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, ParentId: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	// 生成完全树
	resp := GenerateTree(SystemMenus.ConvertToINodeArray(allMenu))
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))

	//root := &SystemMenu{Id: 1, ParentId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"}
	//RecursiveTree(root, SystemMenus.ConvertToINodeArray(allMenu))
	//jsonTree, _ := json.MarshalIndent(root, "", "\t")
	//fmt.Println(string(pretty.Color(pretty.PrettyOptions(jsonTree, pretty.DefaultOptions), nil)))

}
