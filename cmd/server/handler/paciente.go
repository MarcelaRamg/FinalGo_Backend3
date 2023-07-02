package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/internal/paciente"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/web"

	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		s: s,
	}
}

func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

func validateEmptys3(paciente *domain.Paciente) (bool, error) {
	switch {
	case paciente.Nombre == "" || paciente.Apellido == "" || paciente.Dni == "" || paciente.FechaAlta == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateFechaAlta(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid FechaAlta date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid FechaAlta date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid FechaAlta date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys3(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateFechaAlta(paciente.FechaAlta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(paciente)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *pacienteHandler) Delete() gin.HandlerFunc {
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

func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var paciente domain.Paciente
		err = c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys3(&paciente)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateFechaAlta(paciente.FechaAlta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, paciente)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Nombre    string `json:"name,omitempty"`
		Apellido  string `json:"Apellido,omitempty"`
		Dni       string `json:"Dni,omitempty"`
		FechaAlta string `json:"FechaAlta,omitempty"`
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
			web.Failure(c, 404, errors.New("paciente not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Paciente{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Dni:       r.Dni,
			FechaAlta: r.FechaAlta,
		}
		if update.FechaAlta != "" {
			valid, err := validateFechaAlta(update.FechaAlta)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
