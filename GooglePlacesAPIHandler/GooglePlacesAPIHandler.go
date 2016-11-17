package GooglePlacesAPIHandler

import (
  "github.com/kr/pretty"
  "sort"
)

//getHotels receives the destination TripAdvisor ID &
//returns an array of IDs of the top 10 hotels.
func getHotels(destinationID string) []string{
  hotels := []string{}
  return hotels

}
//getNearRestaurants receives the chosen hotel TripAdvisor ID &
//returns an arrays of IDs of the top 10 nearby restaurants.
func getNearRestaurants(hotelID string) []string{
  restaurants := []string{}
  return restaurants
}




//getAttractions receives the destination &
//returns an array of IDs of the top 20 attractions.
func GetAttractions(destination string) []string{
  results := textSearch("attractions", destination)
  places  := formatResults(results)
  sort.Sort(ByRating(places))

	pretty.Println(places)

  attractions := []string{}
  return attractions
}
//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date &
//returns a schedule with timing allocated for each attraction and for lunch time.
func createSchedule(destinationID string, hotelID string, startDate string, endDate string) string{
  return "";
}
