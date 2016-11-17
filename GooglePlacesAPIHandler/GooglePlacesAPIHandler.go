package GooglePlacesAPIHandler

import (
	"botota/models"
	"botota/utils"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"sort"
)

const (
	MaxRestDist = 2000
)

//GetHotels receives the destination city &
//returns an array of Places for the top 20 hotels.
func GetHotels(destination string) []models.Place {
	r := textSearch("hotels", destination)
	h := formatResults(r)
	sort.Sort(models.ByRating(h))
	// pretty.Println(h)
	return h
}

//GetNearRestaurants receives the chosen hotel as a Place struct
//returns an arrays of Places for the top 20 nearby restaurants.
func GetNearRestaurants(hotel models.Place) []models.Place {
	client := CreateClient()
	r := &maps.NearbySearchRequest{
		Location: &hotel.Location,
		Type:     "restaurant",
		Radius:   MaxRestDist}
	resp, err := client.NearbySearch(context.Background(), r)
	utils.Check(err)
	rest := formatResults(resp.Results)
	sort.Sort(models.ByRating(rest))
	// pretty.Println(rest)
	return rest
}

//GetAttractions receives the destination &
//returns an array of Places for the top 20 attractions.
func GetAttractions(destination string) []models.Place {
	r := textSearch("attractions", destination)
	a := formatResults(r)
	sort.Sort(models.ByRating(a))
	// pretty.Println(a)
	return a
}

//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date &
//returns a schedule with timing allocated for each attraction and for lunch time.
func CreateSchedule(destination string, hotel models.Place, startDate string, endDate string) string {
	return ""
}
