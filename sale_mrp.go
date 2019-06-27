package sale_mrp

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.MrpProduction().DeclareModel()

	h.MrpProduction().AddFields(map[string]models.FieldDefinition{
		"SaleName": models.CharField{
			Compute: h.MrpProduction().Methods().ComputeSaleNameSaleRef(),
			String:  "Sale Name",
			Help:    "Indicate the name of sales order.",
		},
		"SaleRef": models.CharField{
			Compute: h.MrpProduction().Methods().ComputeSaleNameSaleRef(),
			String:  "Sale Reference",
			Help:    "Indicate the Customer Reference from sales order.",
		},
	})
	h.MrpProduction().Methods().ComputeSaleNameSaleRef().DeclareMethod(
		`ComputeSaleNameSaleRef`,
		func(rs h.MrpProductionSet) h.MrpProductionData {
			//        def get_parent_move(move):
			//            if move.move_dest_id:
			//                return get_parent_move(move.move_dest_id)
			//            return move
			//        for production in self:
			//            move = get_parent_move(production.move_finished_ids[:1])
			//            production.sale_name = move.procurement_id and move.procurement_id.sale_line_id and move.procurement_id.sale_line_id.order_id.name or False
			//            production.sale_ref = move.procurement_id and move.procurement_id.sale_line_id and move.procurement_id.sale_line_id.order_id.client_order_ref or False
		})
	h.SaleOrderLine().DeclareModel()

	h.SaleOrderLine().Methods().GetDeliveredQty().DeclareMethod(
		`GetDeliveredQty`,
		func(rs m.SaleOrderLineSet) {
			//        self.ensure_one()
			//        precision = self.env['decimal.precision'].precision_get(
			//            'Product Unit of Measure')
			//        bom_delivered = {}
			//        bom = self.env['mrp.bom']._bom_find(product=self.product_id)
			//        if bom and bom.type == 'phantom':
			//            bom_delivered[bom.id] = False
			//            product_uom_qty_bom = self.product_uom._compute_quantity(
			//                self.product_uom_qty, bom.product_uom_id) / bom.product_qty
			//            boms, lines = bom.explode(self.product_id, product_uom_qty_bom)
			//            for bom_line, data in lines:
			//                qty = 0.0
			//                for move in self.procurement_ids.mapped('move_ids'):
			//                    if move.state == 'done' and move.product_id.id == bom_line.product_id.id:
			//                        qty += move.product_uom._compute_quantity(
			//                            move.product_uom_qty, bom_line.product_uom_id)
			//                if float_compare(qty, data['qty'], precision_digits=precision) < 0:
			//                    bom_delivered[bom.id] = False
			//                    break
			//                else:
			//                    bom_delivered[bom.id] = True
			//        if bom_delivered and any(bom_delivered.values()):
			//            return self.product_uom_qty
			//        elif bom_delivered:
			//            return 0.0
			//        return super(SaleOrderLine, self)._get_delivered_qty()
		})
	h.SaleOrderLine().Methods().GetBomComponentQty().DeclareMethod(
		`GetBomComponentQty`,
		func(rs m.SaleOrderLineSet, bom interface{}) {
			//        bom_quantity = self.product_uom._compute_quantity(
			//            1, bom.product_uom_id)
			//        boms, lines = bom.explode(self.product_id, bom_quantity)
			//        components = {}
			//        for line, line_data in lines:
			//            product = line.product_id.id
			//            uom = line.product_uom_id
			//            qty = line.product_qty
			//            if components.get(product, False):
			//                if uom.id != components[product]['uom']:
			//                    from_uom = uom
			//                    to_uom = self.env['product.uom'].browse(
			//                        components[product]['uom'])
			//                    qty = from_uom._compute_quantity(qty, to_uom)
			//                components[product]['qty'] += qty
			//            else:
			//                # To be in the uom reference of the product
			//                to_uom = self.env['product.product'].browse(product).uom_id
			//                if uom.id != to_uom.id:
			//                    from_uom = uom
			//                    qty = from_uom._compute_quantity(qty, to_uom)
			//                components[product] = {'qty': qty, 'uom': to_uom.id}
			//        return components
		})
	h.AccountInvoiceLine().DeclareModel()

	h.AccountInvoiceLine().Methods().GetAngloSaxonPriceUnit().DeclareMethod(
		`GetAngloSaxonPriceUnit`,
		func(rs m.AccountInvoiceLineSet) {
			//        price_unit = super(AccountInvoiceLine,
			//                           self)._get_anglo_saxon_price_unit()
			//        if self.product_id.invoice_policy == "delivery":
			//            for s_line in self.sale_line_ids:
			//                # qtys already invoiced
			//                qty_done = sum([x.uom_id._compute_quantity(x.quantity, x.product_id.uom_id)
			//                                for x in s_line.invoice_lines if x.invoice_id.state in ('open', 'paid')])
			//                quantity = self.uom_id._compute_quantity(
			//                    self.quantity, self.product_id.uom_id)
			//                # Put moves in fixed order by date executed
			//                moves = s_line.mapped(
			//                    'procurement_ids.move_ids').sorted(lambda x: x.date)
			//                # Go through all the moves and do nothing until you get to qty_done
			//                # Beyond qty_done we need to calculate the average of the price_unit
			//                # on the moves we encounter.
			//                bom = s_line.product_id.product_tmpl_id.bom_ids and s_line.product_id.product_tmpl_id.bom_ids[
			//                    0]
			//                if bom.type == 'phantom':
			//                    average_price_unit = 0
			//                    components = s_line._get_bom_component_qty(bom)
			//                    for product_id in components.keys():
			//                        factor = components[product_id]['qty']
			//                        prod_moves = [
			//                            m for m in moves if m.product_id.id == product_id]
			//                        prod_qty_done = factor * qty_done
			//                        prod_quantity = factor * quantity
			//                        average_price_unit += factor * \
			//                            self._compute_average_price(
			//                                prod_qty_done, prod_quantity, prod_moves)
			//                    price_unit = average_price_unit or price_unit
			//                    price_unit = self.product_id.uom_id._compute_price(
			//                        price_unit, self.uom_id)
			//        return price_unit
		})
}
