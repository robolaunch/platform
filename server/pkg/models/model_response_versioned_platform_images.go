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

type ResponseVersionedPlatformImages struct {
	// Success of the request.
	Success bool `json:"success"`
	// Response message.
	Message string `json:"message"`

	Response *VersionedPlatformImages `json:"response"`
}
