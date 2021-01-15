package domain

import (
	"time"
)

// Role
type Role struct {
	// 唯一标识
	Id int64 `json:"id"`
	// 角色名称
	RoleName string `json:"roleName"`
	// 父级Id
	ParentId int64 `json:"parentId"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

type RoleRepository interface {
	Save(dm *Role) (*Role, error)
	Remove(dm *Role) (*Role, error)
	FindOne(queryOptions map[string]interface{}) (*Role, error)
	Find(queryOptions map[string]interface{}) (int64, []*Role, error)
}

func (m *Role) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}

func (m *Role) Update(data map[string]interface{}) error {
	if roleName, ok := data["roleName"]; ok {
		m.RoleName = roleName.(string)
	}
	if parentId, ok := data["parentId"]; ok {
		m.ParentId = parentId.(int64)
	}
	return nil
}

type Roles []*Role

func (roles Roles) RoleMap() map[int64]*Role {
	var retMap = make(map[int64]*Role)
	for _, v := range roles {
		retMap[v.Id] = v
	}
	return retMap
}
func (roles Roles) RoleIds() []int64 {
	var ids []int64
	for _, v := range roles {
		ids = append(ids, v.Id)
	}
	return ids
}

//func(roles Roles)RolesWithMap(roleMap map[int64]*Role,roleIds []int64)string{
//	var role []string
//	for _,id :=range roleIds{
//		if v,ok:=roleMap[id];ok{
//			role = append(role,v.RoleName)
//		}
//	}
//	return strings.Join(role,",")
//}
