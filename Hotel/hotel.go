package Hotel

import "fmt"

type Room struct {
	Number string
	Type   string
	Price  float64
	Busy   bool
}

type Hotel struct {
	Rooms map[string]Room
}

func NewHotel() *Hotel {
	return &Hotel{Rooms: map[string]Room{}}
}

func (h *Hotel) AddRoom(number, roomType string, price float64) {
	h.Rooms[number] = Room{
		Number: number,
		Type:   roomType,
		Price:  price,
	}
}

func (h *Hotel) CheckIn(number string) {
	room := h.Rooms[number]
	room.Busy = true
	h.Rooms[number] = room
}

func (h *Hotel) CheckOut(number string) {
	room := h.Rooms[number]
	room.Busy = false
	h.Rooms[number] = room
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant rooms:")
	for _, r := range h.Rooms {
		if !r.Busy {
			fmt.Printf("Room %s | %s | %.2f\n", r.Number, r.Type, r.Price)
		}
	}
}
