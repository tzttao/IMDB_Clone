package Models

type Cast_list struct {
	Cast_id         uint `gorm:"primary_key"`
	Cast_name       string
	Cast_desciption string
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (Cast_list) TableName() string {
	return "cast_list"
}
