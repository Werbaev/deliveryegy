package api

import (
	"github.com/gin-gonic/gin"
	"github.com/werbaev/deliveryegy/config"
	"github.com/werbaev/deliveryegy/pkg/logger"
	"github.com/werbaev/deliveryegy/storage"

	_ "github.com/werbaev/deliveryegy/api/docs"
	v1 "github.com/werbaev/deliveryegy/api/handlers/v1"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(CORSMiddleware())

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	//config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")

	//router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")
	//apiV1.Use(cors.New(config))
	{

		//user
		apiV1.POST("/user/create", handlerV1.CreateUser)
		apiV1.POST("/user/login", handlerV1.LoginUser)
		apiV1.GET("/user/get/:guid", handlerV1.GetUser)
		apiV1.GET("/user/list", handlerV1.ListUsers)
		apiV1.PUT("/user/update", handlerV1.UpdateUser)
		apiV1.DELETE("/user/delete", handlerV1.DeleteUser)
		apiV1.GET("/user/:guid/orders", handlerV1.GetUserOrders)

		//category
		apiV1.POST("/category/create", handlerV1.CreateCategory)
		apiV1.GET("/category/get/:guid", handlerV1.GetCategory)
		apiV1.GET("/category/list", handlerV1.ListCategories)
		apiV1.PUT("/category/update", handlerV1.UpdateCategory)
		apiV1.DELETE("/category/delete", handlerV1.DeleteCategory)

		//product
		apiV1.POST("/product/create", handlerV1.CreateProduct)
		apiV1.GET("/product/get/:guid", handlerV1.GetProduct)
		apiV1.GET("/product/list", handlerV1.ListProducts)
		apiV1.PUT("/product/update", handlerV1.UpdateProduct)
		apiV1.DELETE("/product/delete", handlerV1.DeleteProduct)

		//order
		apiV1.POST("/order/create", handlerV1.CreateOrder)
		apiV1.GET("/order/get/:guid", handlerV1.GetOrder)
		apiV1.GET("/order/list", handlerV1.ListOrders)
		apiV1.PUT("/order/update", handlerV1.UpdateOrder)
		apiV1.PUT("/order/update-status", handlerV1.UpdateOrderStatus)
		apiV1.DELETE("/order/delete", handlerV1.DeleteOrder)

		//merchant
		apiV1.POST("/merchant/create", handlerV1.CreateMerchant)
		apiV1.GET("/merchant/get/:guid", handlerV1.GetMerchant)
		apiV1.GET("/merchant/get-branches/:guid", handlerV1.GetMerchantBranches)
		apiV1.GET("/merchant/list", handlerV1.ListMerchants)
		apiV1.PUT("/merchant/update", handlerV1.UpdateMerchant)
		apiV1.DELETE("/merchant/delete", handlerV1.DeleteMerchant)
		apiV1.GET("/merchant/products/:guid", handlerV1.GetMerchantProducts)
		apiV1.GET("/merchant/orders/:guid", handlerV1.GetMerchantOrders)

		//merchant-branch
		apiV1.POST("/merchant-branch/create", handlerV1.CreateMerchantBranch)
		apiV1.GET("/merchant-branch/get/:guid", handlerV1.GetMerchantBranch)
		apiV1.GET("/merchant-branch/list", handlerV1.ListMerchantBranches)
		apiV1.PUT("/merchant-branch/update", handlerV1.UpdateMerchantBranch)
		apiV1.DELETE("/merchant-branch/delete", handlerV1.DeleteMerchantBranch)
		apiV1.GET("/merchant-branch/:guid/orders", handlerV1.GetMerchantBranchOrders)

		//courier
		apiV1.POST("/courier/create", handlerV1.CreateCourier)
		apiV1.POST("/courier/login", handlerV1.LoginCourier)
		apiV1.GET("/courier/get/:guid", handlerV1.GetCourier)
		apiV1.GET("/courier/list", handlerV1.ListCouriers)
		apiV1.GET("/courier/:guid/list-orders", handlerV1.ListCourierOrders)
		apiV1.PUT("/courier/update", handlerV1.UpdateCourier)
		apiV1.DELETE("/courier/delete", handlerV1.DeleteCourier)

		//vendor-user
		apiV1.POST("/vendor-user/create", handlerV1.CreateVendorUser)
		apiV1.POST("/vendor-user/login", handlerV1.LoginVendorUser)
		apiV1.GET("/vendor-user/get/:guid", handlerV1.GetVendorUser)
		apiV1.GET("/vendor-user/list", handlerV1.ListVendorUsers)
		apiV1.PUT("/vendor-user/update", handlerV1.UpdateVendorUser)
		apiV1.DELETE("/vendor-user/delete", handlerV1.DeleteVendorUser)
	}

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,upload-offset, upload-metadata, upload-length, tus-resumable, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, HEAD, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
