package routes

import(
	"github.com/gin-gonic/gin"
	controller "github.com/andymartinezot/go-restaurant-management/server/internal/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("invoices/:invoices_id", controller.UpdateInvoice())
}