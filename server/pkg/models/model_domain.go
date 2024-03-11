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

type Domain struct {
	Alias string `json:"alias,omitempty"`

	Description string `json:"description,omitempty"`

	Environments []Environment `json:"environments,omitempty"`
}
