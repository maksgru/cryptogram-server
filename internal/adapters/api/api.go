package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/maksgru/cryptogram-server/internal/adapters/api/user"
)

func NewApi(router httprouter.Router) {
	user.RegisterUserHandlers(&router)

}
