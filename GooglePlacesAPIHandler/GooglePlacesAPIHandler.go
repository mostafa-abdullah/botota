package GooglePlacesAPIHandler

import (
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
func GetNearRestaurants(hotelLat string, hotelLon string) []string{
  client := CreateClient()

  //prepare location parameter
  location := hotelLat + "," + hotelLon
  l, err := maps.ParseLatLng(location)
  utils.Check(err)

  r := &maps.NearbySearchRequest{
    Location:   &l,
    Type:       "restaurant",
    Radius:     MaxRestDist }
  resp, err := client.NearbySearch(context.Background(), r)
  utils.Check(err)
  pretty.Println(resp)

  restaurants := []string{}
  return restaurants
}




//getAttractions receives the destination &
//returns an array of IDs of the top 20 attractions.
func GetAttractions(destination string) []Place{
  r := textSearch("attractions", destination)
  h  := formatResults(r)
  sort.Sort(ByRating(h))
  //pretty.Println(h)
  return h
}
//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date &
//returns a schedule with timing allocated for each attraction and for lunch time.
func createSchedule(destinationID string, hotelID string, startDate string, endDate string) string{
  return "";
}
