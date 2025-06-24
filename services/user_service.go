package services

import (
	"encoding/json"
	"fmt"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/models"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/pkg/erros"
	request2 "github.com/dhiegoemmanuel2006/picpay-simplificado-go/pkg/request"
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/pkg/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{DB: db}
}

func (us *UserService) CreateUser(userRequest request2.CreateUserRequest) (uuid.UUID, error) {
	fmt.Printf("models.LOJISTA= %s", models.LOJISTA)
	fmt.Printf("models.USUARIO= %s", models.USUARIO)
	fmt.Printf("Role request= %s", userRequest.Role)

	var model models.Users
	if userRequest.Role != models.USUARIO && userRequest.Role != models.LOJISTA {
		return uuid.Nil, erros.ErrThisRoleIsNotAllowed
	}
	model.ID = uuid.New()
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	model.Email = userRequest.Email
	model.Role = userRequest.Role
	model.Password = userRequest.Password
	model.FullName = userRequest.FullName
	model.Document = userRequest.Document
	model.Balance = userRequest.Balance

	if err := us.DB.Create(&model).Error; err != nil {
		return uuid.Nil, err
	}
	return model.ID, nil
}

func (us *UserService) GetUserByID(id uuid.UUID) (models.Users, error) {
	var user models.Users
	if err := us.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.Users{}, erros.ErrUserNotFound
	}
	return user, nil
}

func (us *UserService) ValidTransaction(payer, payee models.Users, value float64) (bool, error) {
	if payer.Role == models.LOJISTA {
		return false, erros.ErrLojistaNotCanSendMoney
	}
	if payer.Balance < value {
		return false, erros.ErrPayerDontHaveSufficientBalance
	}
	externAuthorization, err := http.Get("https://util.devi.tools/api/v2/authorize")
	if err != nil {
		return false, err
	}
	var authorization response.ExternAuthorizationResponse
	err = json.NewDecoder(externAuthorization.Body).Decode(&authorization)
	if err != nil {
		return false, err
	}

	if authorization.Data.Authorization == false {
		return false, erros.ErrTranferUnauthorized
	}

	return true, nil
}

func (us *UserService) DoTransaction(payer, payee models.Users, value float64) {
	payer.Balance -= value
	payee.Balance += value

	us.DB.Save(&payer)
	us.DB.Save(&payee)
}
