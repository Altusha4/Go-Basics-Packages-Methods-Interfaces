package main

import (
	"Assignment1/Hotel"
)

func main() {
	hotel := Hotel.NewHotel()

	hotel.AddRoom("101", "Single", 100)
	hotel.AddRoom("102", "Double", 150)

	hotel.CheckIn("101")

	hotel.ListVacantRooms()
}
