package match

import (
	"net/http"

	"github.com/labstack/echo"
)

// AddMatchControllerHTTP struct
type AddMatchControllerHTTP struct {
	inputAdapter  AddMatchInputAdapter
	outputAdapter AddMatchOutputAdapter
	useCase       AddMatchUseCase
}

func (a *AddMatchControllerHTTP) postMatch(context echo.Context) error {
	body := context.Get("parsedBody").(string)
	input, err := a.inputAdapter.Handle(body)
	if err != nil {
		return err
	}

	output, err := a.useCase.Handle(input)

	response, err := a.outputAdapter.Handle(output, err)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

// NewAddMatchControllerHTTP factory
func NewAddMatchControllerHTTP(
	routeGroup *echo.Group,
	inputAdapter AddMatchInputAdapter,
	outputAdapter AddMatchOutputAdapter,
	useCase AddMatchUseCase,
) *AddMatchControllerHTTP {
	handler := &AddMatchControllerHTTP{
		inputAdapter:  inputAdapter,
		outputAdapter: outputAdapter,
		useCase:       useCase,
	}
	routeGroup.POST("/match", handler.postMatch)
	return handler
}