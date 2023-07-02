package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/MarcelaRamg/FinalBack3.git/cmd/server/handler"
	"github.com/MarcelaRamg/FinalBack3.git/internal/dentista"
	"github.com/MarcelaRamg/FinalBack3.git/internal/paciente"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/dentistaPkg"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/pacientePkg"
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
	//Dentista:
	storageDentista := dentistaPkg.NewSQLDentista()
	repoDentista := dentista.NewRepository(storageDentista)
	serviceDentista := dentista.NewService(repoDentista)
	dentistaHandler := handler.NewDentistaHandler(serviceDentista)

	//Paciente
	storagePaciente := pacientePkg.NewSQLPaciente()
	repoPaciente := paciente.NewRepository(storagePaciente)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

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
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}

	//docs.SwaggerInfo.Host = os.Getenv("HOST")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
