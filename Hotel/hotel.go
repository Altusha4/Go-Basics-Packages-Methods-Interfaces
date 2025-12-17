package Hotel

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[string]Room
}

func NewHotel() *Hotel {
	return &Hotel{
		Rooms: make(map[string]Room),
	}
}
