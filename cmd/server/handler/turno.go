package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/internal/turno"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.TurnoService
}

func NewTurnoHandler(s turno.TurnoService) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// @Summary Obtener turno por ID
// @Description Obtiene un turno por su ID
// @Tags Turno
// @Accept json
// @Produce json
// @Param id path int true "ID del turno a obtener"
// @Success 200 {object} TurnoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /turnos/{id} [get]
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

// @Summary Crear turno
// @Description Crea un nuevo turno
// @Tags Turno
// @Accept json
// @Produce json
// @Param TOKEN header string true "Token de autenticación"
// @Param turno body TurnoRequest true "Datos del turno a crear"
// @Success 201 {object} TurnoResponse
// @Failure 400 {object} ErrorResponse
// @Router /turnos [post]
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

// @Summary Eliminar turno
// @Description Elimina un turno existente por su ID
// @Tags Turno
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del turno a eliminar"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /turnos/{id} [delete]
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

// @Summary Actualizar turno
// @Description Actualiza un turno existente por su ID
// @Tags Turno
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del turno a actualizar"
// @Accept json
// @Produce json
// @Param turno body TurnoRequest true "Datos del turno a actualizar"
// @Success 200 {object} TurnoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Router /turnos/{id} [put]
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

// @Summary Actualizar parcialmente un turno
// @Description Actualiza parcialmente un turno existente por su ID
// @Tags Turno
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del turno a actualizar"
// @Accept json
// @Produce json
// @Param turno body TurnoPatchRequest true "Datos parciales del turno a actualizar"
// @Success 200 {object} TurnoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Router /turnos/{id} [patch]
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
