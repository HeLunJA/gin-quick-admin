package system

type RoleRelation struct {
	ParentID uint
	ChildID  uint
}

func (RoleRelation) TableName() string {
	return "sys_role_relation"
}
