# Go Basics: Packages, Methods & Interfaces

A console-based Go application.
This project demonstrates core Go concepts such as **packages, structs, methods, interfaces, maps, slices, and loops** through four independent tasks.

---

## Technologies Used

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

---

## Features

- Modular project structure using Go packages
- Clear usage of structs and methods
- Interfaces for polymorphism
- Maps and slices for data storage
- Loop-based interactive console menus
- Clean and readable beginner-friendly code
- Single main menu for all tasks

---

## Project Structure
```
Assignment1/
├── go.mod
├── main.go
│
├── Hotel/
│ └── hotel.go
│
├── Employee/
│ └── employee.go
│
├── Gym/
│ └── gym.go
│
└── Wallet/
└── wallet.go
```

---

## Implemented Tasks

### Task 1: Hotel Room Reservation System
- Room and Hotel structs
- Room storage using map
- Methods: AddRoom, CheckIn, CheckOut, ListVacantRooms

---

### Task 2: Employee & Interfaces
- Employee interface
- Structs: FullTime, PartTime, Contractor, Intern
- Custom salary calculation logic
- Employees stored in a slice and processed with loops

---

### Task 3: Gym Membership Management
- Member interface with GetDetails()
- Structs: BasicMember, PremiumMember
- Members stored in a map
- Methods: AddMember, ListMembers

---

### Task 4: Wallet Transaction Simulation
- Wallet struct
- Methods: AddMoney, SpendMoney, GetBalance
- Transactions stored in a slice
- Interactive wallet menu loop

---

## How to Run the Project

1. Open the project folder
```bash
cd Assignment1
```
2. Run the application
```bash
go run .
```

## How It Works

1. The application starts with a main menu
2. User selects one of the four tasks
3. Each task demonstrates specific Go concepts
4. User can return to the menu or exit the program

## Author

**Altynay Yertay**
Software Engineering Student
Astana IT University