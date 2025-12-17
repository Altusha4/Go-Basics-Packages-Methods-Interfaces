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

func (h *Hotel) AddRoom(number string, roomType string, price float64) {
	h.Rooms[number] = Room{
		RoomNumber:    number,
		Type:          roomType,
		PricePerNight: price,
		IsOccupied:    false,
	}
}

func (h *Hotel) CheckIn(number string) {
	room, exists := h.Rooms[number]
	if !exists {
		return
	}

	if room.IsOccupied {
		return
	}

	room.IsOccupied = true
	h.Rooms[number] = room
}
