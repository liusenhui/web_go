package services

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"web_go/internal/models"
	"web_go/internal/repositories"
	"web_go/pkg/jwt"
)

type UserService struct {
	userRepository repositories.UserRepositoryIF
}

type UserServiceIF interface {
	CreateService(user *models.User) error
	SignInService(user *models.User) (interface{}, error)
	GetService(id string) interface{}
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// CreateService 注册用户
func (us UserService) CreateService(user *models.User) error {
	if len(user.Password) < 6 {
		return errors.New("密码不能少于六位")
	}
	if len(user.Account) < 4 {
		return errors.New("账号不能少于四位")
	}
	hashedPassword, err := generateBcrypt(user.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	user.Password = hashedPassword

	if len(user.Name) == 0 {
		user.Name = user.Account
	}

	user.CreatedBy = "system"

	err = us.userRepository.CreateRepo(user)
	if err != nil {
		return errors.New("创建失败")
	}
	return nil
}

func (us UserService) SignInService(user *models.User) (interface{}, error) {
	// 获取数据库内用户信息
	u := us.userRepository.SignInRepo(user)

	// 比较密码和哈希值
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("账号不存在或密码错误")
	}
	zap.L().Debug("check password:", zap.String("password", user.Password))
	// 处理session
	token, err := jwt.GenerateToken(u.Id, u.Account)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (us UserService) GetService(id string) interface{} {
	return us.userRepository.GetRepo(id)
}

// generateBcrypt 加密密码
func generateBcrypt(password string) (p string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 加密密码
	if err != nil {
		zap.L().Error("bcrypt generate from password error:", zap.Error(err))
		return "", err
	}
	zap.L().Debug("bcrypt generate from password ", zap.String("password", password), zap.String("hashedPassword", string(hashedPassword)))
	return string(hashedPassword), nil
}
