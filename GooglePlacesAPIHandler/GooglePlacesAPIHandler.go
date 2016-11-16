package GooglePlacesAPIHandler

import (
  "googlemaps.github.io/maps"
  "github.com/kr/pretty"
  "golang.org/x/net/context"
  "botota/utils"
)
const (
  APIKey = "AIzaSyCNRXCIOJkenWGvhiIgu58ncqL6W9VOc3Y"
)

func CreateClient() *maps.Client{
  c, err := maps.NewClient(maps.WithAPIKey(APIKey))
  utils.Check(err);
  return c;
}

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
//getAttractions receives the destination TripAdvisor ID &
//returns an array of IDs of the top 10 attractions.
func GetAttractions(destination string) []string{
  client := CreateClient()

  q := "attractions in " + destination

  r := &maps.TextSearchRequest{Query: q}

  resp, err := client.TextSearch(context.Background(), r)
	utils.Check(err)

	pretty.Println(resp)

  attractions := []string{}
  return attractions
}
//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date &
//returns a schedule with timing allocated for each attraction and for lunch time.
func createSchedule(destinationID string, hotelID string, startDate string, endDate string) string{
  return "";
}
