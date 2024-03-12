package swagger

import (
	"io"
	"net/http"

	"github.com/robolaunch/platform/server/pkg/models"
	"gopkg.in/yaml.v2"
)

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

func GetStructuredPlatformComponents(url string) (models.PlatformComponents, error) {

	robolaunch := models.PlatformComponents{}

	response, err := http.Get(url)
	if err != nil {
		return models.PlatformComponents{}, err
	}
	defer response.Body.Close()

	manifest, err := io.ReadAll(response.Body)
	if err != nil {
		return models.PlatformComponents{}, err
	}

	err = yaml.Unmarshal(manifest, &robolaunch)
	if err != nil {
		return models.PlatformComponents{}, err
	}

	return robolaunch, nil
}

func GetStructuredPlatformImages(url string) (models.PlatformImages, error) {

	robolaunch := models.PlatformImages{}

	response, err := http.Get(url)
	if err != nil {
		return models.PlatformImages{}, err
	}
	defer response.Body.Close()

	manifest, err := io.ReadAll(response.Body)
	if err != nil {
		return models.PlatformImages{}, err
	}

	err = yaml.Unmarshal(manifest, &robolaunch)
	if err != nil {
		return models.PlatformImages{}, err
	}

	return robolaunch, nil
}
