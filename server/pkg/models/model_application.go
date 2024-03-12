/*
 * robolaunch Industry Cloud Platform - OpenAPI 3.0
 *
 * API for robolaunch ICP platform structure.
 *
 * API version: 1.0.11
 * Contact: info@robolaunch.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type Application struct {
	Name string `json:"name"`

	Version string `json:"version"`

	Alias string `json:"alias"`

	Description string `json:"description"`

	Icon string `json:"icon"`
}