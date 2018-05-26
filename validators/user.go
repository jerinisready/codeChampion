package validators

import (
	"go-webapp/common"
	"go-webapp/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type CodeUserModelValidator struct {
	Username  string      `form:"username" json:"username" binding:"exists,alphanum,min=3,max=255"`
	Password  string      `form:"password" json:"password" binding:"exists,min=4,max=255"`
	CodeUserModel models.CodeUser `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *CodeUserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)

	if err != nil {
		log.WithFields(log.Fields{
				"self": self,
		}).Info("Error parsing", err)
		return err
	}
	self.CodeUserModel.UserName = self.Username
	self.CodeUserModel.Password = self.Password
	return nil
}

// You can put the default value of a Validator here
func NewCodeUserModelValidator() CodeUserModelValidator {
	codeUserModelValidator := CodeUserModelValidator{}
	return codeUserModelValidator
}

func NewCodeUserModelValidatorFillWith(userModel models.CodeUser) CodeUserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.Username = userModel.UserName
	userModelValidator.Password = userModel.Password
	return userModelValidator
}
