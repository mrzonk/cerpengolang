package configs

import "cerpengolang/models/user"
import "cerpengolang/models/cerpen"
import "cerpengolang/models/comment"
func initMigrate() {
	DB.AutoMigrate(&user.User{},&comment.Comment{},&cerpen.Cerpen{},)
}
