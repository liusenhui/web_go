package repositories

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"web_go/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryIF interface {
	CreateRepo(user *models.User) error
	SignInRepo(user *models.User) *models.User
	GetRepo(id string) interface{}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateRepo 创建用户的逻辑
func (ur *UserRepository) CreateRepo(user *models.User) error {
	/*sqlStr := "INTER INTO users(account,name,password,created_by) values (?,?,?,?)"
	db.Exec(sqlStr, user.Account, user.Name, user.Password, user.CreateBy)*/
	zap.L().Debug("create user")
	if err := ur.DB.Create(&user).Error; err != nil {
		zap.L().Error("create user", zap.Error(err))
		return err
	}
	return nil // 示例成功返回
}

// SignInRepo 实现获取用户的逻辑
func (ur *UserRepository) SignInRepo(user *models.User) *models.User {
	//sqlStr := "SELECT account,name,created_by,created_at From users WHERE id = ?"
	var u models.User
	err := ur.DB.Select("id,account,name,password,created_by,created_at").Where("account = ?", user.Account).Find(&u).Error
	if err != nil {
		zap.L().Error("get user info error ", zap.Error(err))
	}
	return &u
}

// GetRepo 实现获取用户的逻辑
func (ur *UserRepository) GetRepo(id string) interface{} {
	//sqlStr := "SELECT account,name,created_by,created_at From users WHERE id = ?"
	var user models.User
	err := ur.DB.Select("id,account,name,created_by,created_at").Where("id = ?", id).Find(&user).Error
	if err != nil {
		zap.L().Error("get user info error ", zap.Error(err))
	}
	return user
}
