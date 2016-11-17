package GooglePlacesAPIHandler
import (
  "googlemaps.github.io/maps"
  "golang.org/x/net/context"
  "botota/utils"
)
type Place struct{
  Name string
  Location maps.LatLng
  Rating float32
}
const (
  APIKey = "AIzaSyCNRXCIOJkenWGvhiIgu58ncqL6W9VOc3Y"
)

// ByRating implements sort.Interface for []maps.PlacesSearchResult based on
// the Rating field.
type ByRating []Place
func (a ByRating) Len() int           { return len(a) }
func (a ByRating) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRating) Less(i, j int) bool { return a[i].Rating > a[j].Rating }

func CreateClient() *maps.Client{
  c, err := maps.NewClient(maps.WithAPIKey(APIKey))
  utils.Check(err);
  return c;
}

func formatResults(results []maps.PlacesSearchResult) []Place{
  places := []Place{}
  for _,r := range results {
    p := Place{r.Name, r.Geometry.Location, r.Rating}
    places = append(places, p)
  }
  return places
}

func textSearch(placeType string, destination string) []maps.PlacesSearchResult {
  client := CreateClient()

  q := placeType + " in " + destination

  r := &maps.TextSearchRequest{Query: q}

  resp, err := client.TextSearch(context.Background(), r)
  utils.Check(err)

  results := resp.Results
  return results
}