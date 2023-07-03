package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/MarcelaRamg/FinalBack3.git/internal/dentista"
	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.DentistaService
}

func NewDentistaHandler(s dentista.DentistaService) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

// ListDentista godoc
// @Summary Obtener dentista por ID
// @Description Obtiene un dentista por su ID
// @Tags Dentista
// @Accept json
// @Produce json
// @Param id path int true "ID del dentista"
// @Success 200
// @Failure 400
// @Failure 404
// @Router /dentistas/{id} [get]
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id invalido"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("no se encontró al dentista"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

func validateEmptys(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Nombre == "" || dentista.Apellido == "" || dentista.Matricula == "":
		return false, errors.New("los campos no pueden estar vacios")
	}
	return true, nil
}

// @Summary Crear dentista
// @Description Crea un nuevo dentista
// @Tags Dentista
// @Accept json
// @Produce json
// @Param TOKEN header string true "Token de autenticación"
// @Param dentista body domain.Dentista true "Información del dentista a crear"
// @Success 201
// @Failure 400
// @Failure 401
// @Router /dentistas [post]
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// @Summary Eliminar dentista por ID
// @Description Elimina un dentista por su ID
// @Tags Dentista
// @Accept json
// @Produce json
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del dentista a eliminar"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /dentistas/{id} [delete]
func (h *dentistaHandler) Delete() gin.HandlerFunc {
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

// @Summary Actualizar dentista por ID
// @Description Actualiza los datos de un dentista existente por su ID
// @Tags Dentista
// @Accept json
// @Produce json
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del dentista a actualizar"
// @Param dentista body domain.Dentista true "Datos del dentista a actualizar"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 409
// @Router /dentistas/{id} [put]
func (h *dentistaHandler) Put() gin.HandlerFunc {
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
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista no encontrado"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// @Summary Actualizar parcialmente dentista por ID
// @Description Actualiza parcialmente los datos de un dentista existente por su ID
// @Tags Dentista
// @Accept json
// @Produce json
// @Param TOKEN header string true "Token de autenticación"
// @Param id path int true "ID del dentista a actualizar"
// @Param update body domain.Dentista true "Datos a actualizar del dentista"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 409
// @Router /dentistas/{id} [patch]
func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Apellido  string `json:"apellido,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Matricula string `json:"matricula,omitempty"`
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
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentista{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}

		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
