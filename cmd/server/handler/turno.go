package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/MarcelaRamg/FinalBack3.git/internal/dentista"
	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/internal/paciente"
	"github.com/MarcelaRamg/FinalBack3.git/internal/turno"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.TurnoService
}

// NewTurnoHandler crea un nuevo controller de turno
func NewTurnoHandler(s turno.TurnoService) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// Get obtiene un turno por id
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id invalido"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("no se encontró al turno"))
			return
		}
		web.Success(c, 200, turno)
	}
}

// Get obtiene un turno por DNI
func (h *turnoHandler) GetByDni(servicePaciente paciente.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Query("DNI")
		dni, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			web.Failure(c, 400, errors.New("Dni invalido"))
			return
		}

		turno, err := h.s.GetByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New("no se encontró al turno"))
			return
		}
		web.Success(c, 200, turno)
	}
}

// Post crea un nuevo turno
func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var turno domain.Turno
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		p, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *turnoHandler) PostByDniAndMatricula(servicePaciente paciente.Service, serviceDentista dentista.DentistaService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var turnoAux domain.TurnoByMatriculaAndDni
		var turno domain.Turno

		err := c.ShouldBindJSON(&turnoAux)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		paciente, err := servicePaciente.GetByDni(turnoAux.Dni)
		if err != nil {
			web.Failure(c, 404, errors.New("Paciente inexistente"))

		}
		turno.PacienteID = fmt.Sprint(paciente.ID)
		odontologo, err := serviceDentista.GetByMatricula(turnoAux.Matricula)

		if err != nil {
			web.Failure(c, 404, errors.New("Odontologo inexistente"))

		}
		turno.DentistaID = fmt.Sprint(odontologo.ID)

		turno.FechaHora = turnoAux.FechaHora
		turno.Descripcion = turnoAux.Descripcion

		p, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un turno
func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un turno
func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var turno domain.Turno
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno no encontrado"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		p, err := h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un turno o alguno de sus campos
func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		FechaHora   string `json:"fechaHora,omitempty"`
		Descripcion string `json:"descripcion,omitempty"`
		PacienteID  string `json:"paciente,omitempty"`
		DentistaID  string `json:"odontologo,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Turno{
			FechaHora:   r.FechaHora,
			Descripcion: r.Descripcion,
			PacienteID:  r.PacienteID,
			DentistaID:  r.DentistaID,
		}

		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
