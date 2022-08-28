package model

//执行数据迁移
func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&Article{}, &Role{}, &User{})
	if err != nil {
		return
	}
	CreateDefaultRoles()
}
