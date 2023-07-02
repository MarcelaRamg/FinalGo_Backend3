package handler

import (
	"errors"
	"strconv"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
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

// Post crea un nuevo turno
func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
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

// Delete elimina un turno
func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		Descripcion string `json:"Descripcion,omitempty"`
		PacienteID  string `json:"paciente_id,omitempty"`
		DentistaID  string `json:"dentista_id,omitempty"`
	}
	return func(c *gin.Context) {
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
