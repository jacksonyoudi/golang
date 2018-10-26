package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username        string   `validate:"gt=0"`
	// 同上
	PasswordNew     string   `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat  string   `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email           string   `validate:"email"`
}



func validate_bak(req RegisterReq) error {
	err := validator.Validate.Struct(req)
	if err != nil {
	}
	return nil
}


func main()  {

	var req = RegisterReq {
		Username       : "Xargin",
		PasswordNew    : "ohno",
		PasswordRepeat : "ohn",
		Email          : "alex@abc.com",
	}
	validate(req)
}

