package sale_mrp

import (
	"github.com/hexya-erp/pool/h"
)

//vars

//rights
func init() {
	h.Mrp.ModelMrpBom().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Sale.ModelSaleOrder().Methods().Load().AllowGroup(GroupMrpUser)
	h.Sale.ModelSaleOrder().Methods().Write().AllowGroup(GroupMrpUser)
	h.Sale.ModelSaleOrderLine().Methods().Load().AllowGroup(GroupMrpUser)
	h.Sale.ModelSaleOrderLine().Methods().Write().AllowGroup(GroupMrpUser)
	h.Mrp.ModelMrpProduction().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpProduction().Methods().Write().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpProduction().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpWorkorder().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpWorkorder().Methods().Create().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpBom().Methods().Load().AllowGroup(GroupSaleSalesman)
	h.Mrp.ModelMrpBomLine().Methods().Load().AllowGroup(GroupSaleSalesman)
}
