package migrate

type Migrate interface {
	AutoMigrate()
}

func AutoMigrate() {
	new(option).AutoMigrate()
	new(user).AutoMigrate()
	new(asset).AutoMigrate()
	new(role).AutoMigrate()
	new(authAccess).AutoMigrate()
	new(AdminMenu).AutoMigrate()
}
