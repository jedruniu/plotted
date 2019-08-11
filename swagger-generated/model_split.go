/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Split struct {
	// The average speed of this split, in meters per second
	AverageSpeed float32 `json:"average_speed,omitempty"`
	// The distance of this split, in meters
	Distance float32 `json:"distance,omitempty"`
	// The elapsed time of this split, in seconds
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The elevation difference of this split, in meters
	ElevationDifference float32 `json:"elevation_difference,omitempty"`
	// The pacing zone of this split
	PaceZone int32 `json:"pace_zone,omitempty"`
	// The moving time of this split, in seconds
	MovingTime int32 `json:"moving_time,omitempty"`
	// N/A
	Split int32 `json:"split,omitempty"`
}