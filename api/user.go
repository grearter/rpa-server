package api

type Role string

const (
	RoleAdmin    Role = "admin"    // 管理员
	RoleBusiness Role = "business" // 业务人员
	RoleOps      Role = "ops"      // 运维
)

type User struct {
	ID    string `json:"name" bson:"_id"`
	Nick  string `json:"nick" bson:"nick"`
	Mail  string `json:"mail" bson:"main"`
	Phone string `json:"phone" bson:"phone"`
	Role  Role   `json:"role" bson:"role"`
	Auth  string `json:"-" bson:"auth"`
}
