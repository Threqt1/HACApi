package controllers

import (
	"fmt"
	"time"

	"github.com/Threqt1/HACApi/app/models"
	"github.com/Threqt1/HACApi/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

// PostIPR handles POST requests to the IPR endpoint.
//
//	@Description	Returns the IPR(s) for the user. If the date parameter is not passed into the body or is invalid, the most recent IPR is returned.
//	@Description	It is important the format of the date follows the format "01/02/2006" (01 = month, 02 = day, 2006 = year), with leading zeros like shown in the format.
//	@Description	For all possible dates, refer to the "/ipr/all" endpoint.
//	@Tags			ipr
//	@Param			request	body	models.IprRequestBody	false	"Body Params"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.IPRResponse
//	@Router			/ipr [post]
func PostIPR(server *repository.Server, ctx *fiber.Ctx) error {
	// Parse body.
	params := new(models.IprRequestBody)

	// Check if parsing succeeded.
	if err := ctx.BodyParser(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.IPRResponse{
			HTTPError: models.HTTPError{
				Error:   true,
				Message: repository.ErrorBadBodyParams.Error(),
			},
		})
	}

	// Verify the validity of the body parameters.
	valid := true

	if err := server.Validator.Struct(params); err != nil {
		valid = false
	}

	// Confirm the date is valid.
	if _, err := time.Parse("01/02/2006", params.Date); len(params.Date) > 0 && err != nil {
		valid = false
	}

	// If they aren't valid, send back an error.
	if !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.IPRResponse{
			HTTPError: models.HTTPError{
				Error:   true,
				Message: repository.ErrorBadBodyParams.Error(),
			},
		})
	}

	// Form a cache key.
	cacheKey := fmt.Sprintf("%s\n%s\n%s", params.Username, params.Password, params.Base)

	// Try logging in, or grab the cached collector.
	collector, err := server.Cache.GetOrLogin(cacheKey)

	// Check if login succeeded.
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.IPRResponse{
			HTTPError: models.HTTPError{
				Error:   true,
				Message: repository.ErrorInvalidAuthentication.Error(),
			},
		})
	}

	// Get IPR.
	ipr, err := server.Querier.GetIPR(collector, *params)

	// Check if getting IPR succeeded.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.IPRResponse{
			HTTPError: models.HTTPError{
				Error:   true,
				Message: repository.ErrorInternalError.Error(),
			},
		})
	}

	// Return the IPR.
	return ctx.Status(fiber.StatusOK).JSON(models.IPRResponse{
		IPR: ipr,
	})
}
