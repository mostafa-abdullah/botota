package GooglePlacesAPIHandler

import (
	"botota/models"
	"botota/utils"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"sort"
	"time"
	"strconv"
	"fmt"
)

const (
	MaxRestDist = 2000
	HOTEL_DEPARTURE = 10
	BUFFER_TIME = 2
	RETURN_TIME = 15
	TIME_FORMAT = "02/01/2006"
	NANO_TO_HOUR = 3600000000000

)

//GetHotels receives the destination city &
//returns an array of Places for the top 20 hotels.
func GetHotels(destination string) []models.Place {
	r := textSearch("hotels", destination)
	h := formatResults(r)
	sort.Sort(models.ByRating(h))

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
		return rest
	}

	//GetAttractions receives the destination &
	//returns an array of Places for the top 20 attractions.
	func GetAttractions(destination string) []models.Place {
		r := textSearch("attractions", destination)
		a := formatResults(r)
		sort.Sort(models.ByRating(a))
		return a
	}


	func CreateSchedule(destination string, hotel models.Place, startDate string, endDate string) []models.Message{
		attractions := GetAttractions(destination)
		restaurants := GetNearRestaurants(hotel)

		t1, _ := time.Parse(TIME_FORMAT, startDate)
		t2, _ := time.Parse(TIME_FORMAT, endDate)

		duration := t2.Sub(t1)
		days := int(duration.Hours()) / 24

		// start from -1 and increment at each iteration to prevent revisiting the same attraction on consecutive days
		curPlaceIdx := -1

		res := []models.Message{}
		for i := 0; i < days; i++ {
			//Message attributes
			highlight	:= ""
			value				:= ""
			image				:= ""

			curHour := HOTEL_DEPARTURE

			curPlaceIdx++
			curPlaceIdx %= len(attractions)

			highlight = fmt.Sprintf("Day %d\n", i + 1)
			for curHour < RETURN_TIME {
				value += attractionStr(i + 1, attractions[curPlaceIdx].Name, curHour)

				curHour += BUFFER_TIME

				transportaionTime := timeBetweenTwoPlaces(
					attractions[curPlaceIdx],
					attractions[(curPlaceIdx+1)%len(attractions)])

				if curHour + transportaionTime >= RETURN_TIME {
					break
				}

				curHour += transportaionTime
				curPlaceIdx++
				curPlaceIdx %= len(attractions)
			}

			curHour += timeBetweenTwoPlaces(attractions[curPlaceIdx], restaurants[i % len(restaurants)])
			value += restaurantStr(i + 1, restaurants[i % len(restaurants)].Name, curHour)

			res = append(res, models.Message{highlight,value,image})
		}

		return res;
	}

	//formVisit receives the day number, attractionID and hour of the visit &
	//returns a message that informs the user about it.
	func attractionStr(day int, attraction string, startHour int) string{
		return attraction + " from " + strconv.Itoa(startHour) + ":00 until " + strconv.Itoa(startHour + BUFFER_TIME) +":00 .\n"
	}

	//formVisit receives the day number, attractionID and hour of the visit &
	//returns a message that informs the user about it.
	func restaurantStr(day int, rest string, startHour int) string{
		return "Finally eat at " + rest + " near your hotel.\n"
	}

	func timeBetweenTwoPlaces(s models.Place,d models.Place) int{
		client := CreateClient()
		origin := []string {floatToStr(s.Location.Lat) +","+ floatToStr(s.Location.Lng)}
		dest := []string {floatToStr(d.Location.Lat) +","+ floatToStr(d.Location.Lng)}

		r := &maps.DistanceMatrixRequest{
			Origins: origin,
			Destinations: dest,
		}

		resp, err := client.DistanceMatrix(context.Background(), r)
		utils.Check(err)

		// Ceil the Nanoseconds result to the nearest higher hour
		return int((resp.Rows[0].Elements[0].Duration.Nanoseconds() + NANO_TO_HOUR - 1 )/ NANO_TO_HOUR);
	}

	func floatToStr(input_num float64) string {
		return strconv.FormatFloat(input_num, 'f', -1, 64)
	}
