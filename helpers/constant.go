package helpers

import "final-project/models"

func GetConstant() models.Constant {
	var constant models.Constant
	constant.AppJSON = "application/json"
	return constant
}
