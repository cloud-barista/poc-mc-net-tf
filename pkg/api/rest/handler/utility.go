/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handler

import (
	"net/http"

	"github.com/cloud-barista/mc-terrarium/pkg/api/rest/model"
	"github.com/cloud-barista/mc-terrarium/pkg/readyz"
	"github.com/cloud-barista/mc-terrarium/pkg/tofu"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Readyz func is for checking mc-terrarium server is ready.
// Readyz godoc
// @Summary Check mc-terrarium server is ready
// @Description Check mc-terrarium server is ready
// @Tags [System] Utility
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ResponseText
// @Failure 503 {object} model.ResponseText
// @Router /readyz [get]
func Readyz(c echo.Context) error {
	res := model.ResponseText{}
	if !readyz.IsReady() {
		res.Success = false
		res.Text = "mc-terrarium server is NOT ready"
		return c.JSON(http.StatusServiceUnavailable, &res)
	}
	res.Success = true
	res.Text = "mc-terrarium server is ready"
	return c.JSON(http.StatusOK, &res)
}

// HTTPVersion godoc
// @Summary Check HTTP version of incoming request
// @Description Checks and logs the HTTP version of the incoming request to the server console.
// @Tags [System] Utility
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ResponseText
// @Failure 404 {object} model.ResponseText
// @Failure 500 {object} model.ResponseText
// @Router /httpVersion [get]
func HTTPVersion(c echo.Context) error {
	// Access the *http.Request object from the echo.Context
	req := c.Request()

	// Determine the HTTP protocol version of the request
	res := model.ResponseText{Success: true, Text: req.Proto}

	return c.JSON(http.StatusOK, res)
}

// TofuVersion godoc
// @Summary Check Tofu version
// @Description Check Tofu version
// @Tags [System] Utility
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ResponseText
// @Failure 503 {object} model.ResponseText
// @Router /tofuVersion [get]
func TofuVersion(c echo.Context) error {

	ret, err := tofu.GetTofuVersion()
	if err != nil {
		res := model.ResponseText{Success: false, Text: "failed to get Tofu version"}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := model.ResponseText{Success: true, Text: ret}

	log.Debug().Msgf("%+v", res) // debug

	return c.JSON(http.StatusOK, res)
}