package models

import (
	"googlemaps.github.io/maps"
)

type Place struct {
	Name     string
	Location maps.LatLng
	Rating   float32
}

// ByRating implements sort.Interface for []maps.PlacesSearchResult based on
// the Rating field.
type ByRating []Place

func (a ByRating) Len() int           { return len(a) }
func (a ByRating) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRating) Less(i, j int) bool { return a[i].Rating > a[j].Rating }
