/*
 * robolaunch Industry Cloud Platform - OpenAPI 3.0
 *
 * API for robolaunch ICP platform structure.
 *
 * API version: 1.0.11
 * Contact: info@robolaunch.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"io"
	"net/http"
	"net/url"

	"github.com/robolaunch/platform/server/pkg/models"
	"gopkg.in/yaml.v2"
)

func GetPlatform(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(structToJSONByteArray(GetPlatformResponse(r.URL.Query())))
}

func GetPlatformResponse(queryParams url.Values) models.ResponseRobolaunch {

	url := queryParams.Get("url")

	robolaunch, err := GetStructuredPlatform(url)
	if err != nil {
		return models.ResponseRobolaunch{
			Success: false,
			Message: err.Error(),
		}
	}

	return models.ResponseRobolaunch{
		Success:  true,
		Message:  "All active versions of the robolaunch ICP are listed.",
		Response: &robolaunch,
	}
}

func GetStructuredPlatform(url string) (models.Robolaunch, error) {

	robolaunch := models.Robolaunch{}

	response, err := http.Get(url)
	if err != nil {
		return models.Robolaunch{}, err
	}
	defer response.Body.Close()

	manifest, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Robolaunch{}, err
	}

	err = yaml.Unmarshal(manifest, &robolaunch)
	if err != nil {
		return models.Robolaunch{}, err
	}

	return robolaunch, nil
}
