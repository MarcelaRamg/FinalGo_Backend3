# FinalBack3
Adding go.mod:
go mod init github.com/MarcelaRamg/FinalGo_Backend3.git

Adding go.sum:
go get -u github.com/gin-gonic/gin
go get "github.com/go-sql-driver/mysql"
go get -u github.com/joho/godotenv
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g FinalGo_Backend3/cmd/main.go

add documentation:
swag init ../doc

si hay problemas con el docs.go:
go get -u github.com/swaggo/swag 




docker-compose up
go run .\cmd\main.go


port: 8080

Endpoints:
- Get dentistas by ID: http://localhost:8080/dentistas/:id
- Get All dentistas: http://localhost:8080/dentistas
- Post New dentistas: http://localhost:8080/dentistas
- Delete dentistas by ID: http://localhost:8080/dentistas/:id

- Get pacientes by ID: http://localhost:8080/pacientes/:id
- Get All pacientes: http://localhost:8080/pacientes
- Post New pacientes: http://localhost:8080/pacientes
- Delete pacientes by ID: http://localhost:8080/pacientes/:id

- Get turnos by ID: http://localhost:8080/turnos/:id
- Get All turnos: http://localhost:8080/turnos
- Post New turnos: http://localhost:8080/turnos
- Delete turnos by ID: http://localhost:8080/turnos/:id


- Documentation: http://localhost:8080/swagger/index.html
