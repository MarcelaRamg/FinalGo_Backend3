package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/MarcelaRamg/FinalBack3.git/docs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/MarcelaRamg/FinalBack3.git/cmd/server/handler"
	"github.com/MarcelaRamg/FinalBack3.git/internal/dentista"
	"github.com/MarcelaRamg/FinalBack3.git/internal/paciente"
	"github.com/MarcelaRamg/FinalBack3.git/internal/turno"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/dentistaPkg"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/pacientePkg"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/turnoPkg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error al intentar cargar archivo .env")
	}
	username := os.Getenv("USER_MYSQL")
	password := os.Getenv("PASS_MYSQL")
	dbName := os.Getenv("DB_MYSQL")

	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing.Error())
	}

	//Dentista:
	storageDentista := dentistaPkg.NewSQLDentista(db)
	repoDentista := dentista.NewRepository(storageDentista)
	serviceDentista := dentista.NewService(repoDentista)
	dentistaHandler := handler.NewDentistaHandler(serviceDentista)

	//Paciente
	storagePaciente := pacientePkg.NewSQLPaciente(db)
	repoPaciente := paciente.NewRepository(storagePaciente)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	//Turno
	storageTurno := turnoPkg.NewSQLTurno(db)
	repoTurno := turno.NewTurnoRepository(storageTurno)
	serviceTurno := turno.NewService(repoTurno)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	turnos := r.Group("/turnos")
	{
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.GET("/dni", turnoHandler.GetByDni(servicePaciente))
		turnos.POST("/sinId", turnoHandler.PostByDniAndMatricula(servicePaciente, serviceDentista))
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.PUT(":id", turnoHandler.Put())
	}

	r.Run(":8080")
}
