package httpapi

import (
	"net/http"
	"strconv"

	"github.com/bachhm.dev/clean-architecture-service/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type apiController struct {
	service service.WeatherService
}

func NewAPIController(s service.WeatherService) apiController {
	return apiController{service: s}
}

type WeatherParam struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (wp *WeatherParam) Bind(r *http.Request) error {
	return nil
}

func (api apiController) getWeather(w http.ResponseWriter, r *http.Request) {
	var param WeatherParam

	longitudeStr := r.URL.Query().Get("longitude")
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid longitude"})
		return
	}
	param.Longitude = longitude

	latitudeStr := r.URL.Query().Get("latitude")
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid latitude"})
		return
	}
	param.Latitude = latitude

	result, err := api.service.GetWeather(r.Context(), param.Latitude, param.Longitude)

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.JSON(w, r, result)
}

func (api apiController) SetUpRoute(router chi.Router) {
	router.Get("/weather", api.getWeather)
}
