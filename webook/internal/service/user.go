package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
)

var (
	ErrUserDuplicate          = repository.ErrUserDuplicate
	ErrInvalidEmailOrPassword = errors.New("邮箱或密码错误")
	ErrUserNotFound           = repository.ErrUserNotFound
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(encrypted)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return domain.User{}, ErrInvalidEmailOrPassword
		}
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidEmailOrPassword
	}
	return u, nil
}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	return svc.repo.FindById(ctx, id)
}

func (svc *UserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	// 快路径，大部分请求都会进来这里
	u, err := svc.repo.FindByPhone(ctx, phone)
	if !errors.Is(err, ErrUserNotFound) {
		// 注意 err == nil 也会来这里，返回 u
		return u, err
	}
	// 触发降级之后不执行慢路径
	// if ctx.Value("降级") == "true" {
	// 	return domain.User{}, errors.New("触发系统降级")
	// }
	// 慢路径
	// 执行注册
	err = svc.repo.Create(ctx, domain.User{Phone: phone})
	// ErrUserDuplicate 错误表明新用户已经存在，可能是并发情况下的重复创建
	if err != nil && !errors.Is(err, ErrUserDuplicate) {
		return domain.User{}, err
	}
	// 这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}
