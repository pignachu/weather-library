package weather

import "errors"

// Prediction describes a weather prediction.
type Prediction uint8

// The supported weather prediction types (enum type).
const (
	Sunny Prediction = iota
	Rain
	Overcast
	Snow
	Unknown
)

// ToString returns the corresponding weather as a string
func (p Prediction) ToString() string {
	return [...]string{
		"Sunny",
		"Rain",
		"Overcast",
		"Snow",
		"Unknown",
	}[p]
}

// PredictAtCoords returns a weather prediction for the specified GPS coordinates.
func PredictAtCoords(lat, long float64) (Prediction, error) {

	if lat < -90 || lat > 90 {
		return Unknown, errors.New("invalid latitude")
	} else if long < -180 || long > 180 {
		return Unknown, errors.New("invalid longitude")
	} else if lat >= -90 && lat < -60 {
		return Snow, nil
	} else if lat >= -60 && lat < -20 {
		return Sunny, nil
	} else if lat >= -20 && lat < 20 {
		return Overcast, nil
	} else if lat >= 20 && lat < 60 {
		return Rain, nil
	} else if lat >= 60 && lat <= 90 {
		return Snow, nil
	}

	return Unknown, nil
}

// GetVersion returns the package version
func GetVersion() string {
	return "v1.0.0"
}
