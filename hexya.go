package sale_mrp

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "sale_mrp"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
