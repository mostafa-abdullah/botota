package GooglePlacesAPIHandler

import (
  "sort"
  "googlemaps.github.io/maps"
  "golang.org/x/net/context"
  "botota/utils"
  "github.com/kr/pretty"
)
const(
  MaxRestDist = 2000
)
//getHotels receives the destination TripAdvisor ID &
//returns an array of IDs of the top 10 hotels.
func getHotels(destinationID string) []string{
  hotels := []string{}
  return hotels

}
//getNearRestaurants receives the chosen hotel TripAdvisor ID &
//returns an arrays of IDs of the top 10 nearby restaurants.
func GetNearRestaurants(hotel Place) []Place{
  client := CreateClient()
  r := &maps.NearbySearchRequest{
    Location:   &hotel.Location,
    Type:       "restaurant",
    Radius:     MaxRestDist }
  resp, err := client.NearbySearch(context.Background(), r)
  utils.Check(err)
  rest := formatResults(resp.Results)
  sort.Sort(ByRating(rest))
  pretty.Println(rest)

  return rest
}




//getAttractions receives the destination &
//returns an array of IDs of the top 20 attractions.
func GetAttractions(destination string) []Place{
  r := textSearch("attractions", destination)
  a  := formatResults(r)
  sort.Sort(ByRating(a))
  pretty.Println(a)
  return a
}
//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date &
//returns a schedule with timing allocated for each attraction and for lunch time.
func createSchedule(destinationID string, hotelID string, startDate string, endDate string) string{
  return "";
}
