package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ybalcin/user-management/pkg/err"
)

type response struct {
	data interface{}
	e    *err.Error

	c *fiber.Ctx
}

func New(c *fiber.Ctx) *response {
	return &response{
		c: c,
	}
}

func (r *response) Data(data interface{}) *response {
	r.data = data
	return r
}

func (r *response) Error(e *err.Error) *response {
	r.e = e
	return r
}

func (r *response) JSON() error {
	if r.e != nil {
		return r.c.Status(r.e.Code).JSON(r.e)
	}

	return r.c.JSON(r.data)
}
