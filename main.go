package main

import (
  "botota/server"
)

func main(){
  server.StartServer()
}

//createSchedule receives the TripAdvisor IDs of the destination,the chosen hotel, start date & end date
//returns a schedule with timing allocated for each attraction and for lunch time
//with travelling distance into consideration using Google Maps
func createSchedule(destinationID string, hotelID string, startDate string, endDate string) string{

  return "";
}
