package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MarcelaRamg/FinalBack3.git/cmd/server/handler"
	"github.com/MarcelaRamg/FinalBack3.git/internal/dentista"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/store"
	"github.com/gin-gonic/gin"
)

// @title           Web Api
// @version         1.0
// @description     This Api is to handle dentistas paciente and turnos
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	// storage := store.NewJsonStore("../../products.json")

	db, err := sql.Open("mysql", "root:Fukurokuju77@tcp/finalBack3")
	if err != nil {
		panic(err)
	}
	storage := store.NewDentistaSql(db)

	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	dentistaHandler := handler.NewDentistaHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PATCH(":id", dentistaHandler.Patch())
		dentistas.PUT(":id", dentistaHandler.Put())
	}
	//docs.SwaggerInfo.Host = os.Getenv("HOST")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
