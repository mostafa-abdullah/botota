package models
import(
  "googlemaps.github.io/maps"
)
type Place struct {
	Name     string
	Location maps.LatLng
	Rating   float32
}
