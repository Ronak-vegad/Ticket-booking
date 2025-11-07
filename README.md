Ticket Booking System (Go)
🧾 Project Overview

This is a simple Go-based console application for managing airline ticket bookings.
This project is part of my Go learning journey — I’m continuously updating and improving it.
It currently runs as a command-line tool, handling basic ticket reservations, input validation, and booking summaries.

🚀 Features

✅ User-friendly console input system
✅ Dynamic ticket booking using Go slices
✅ Input validation for:

First and last names (minimum length check)

Email format (must contain "@")

Ticket quantity (positive & within remaining limit)
✅ Tracks total and remaining tickets dynamically
✅ Displays list of first names of booked users
✅ Graceful handling of invalid input

🧩 Technologies Used

Language: Go (Golang)

Core Packages:

fmt → Input/output operations

strings → String manipulation and validation
How It Works

Run the program in your terminal using:

go run main.go


The program welcomes you with a message.

Enter your details:

First Name

Last Name

Email Address

Number of tickets you want to book

If input is valid, your booking is confirmed and added to the list.

If invalid, an error message will explain what went wrong.

Booking continues in a loop until tickets are sold out.

📦 Example Output
--------------- Welcome to Air Tarsai ---------------
Enter your first name: Ronak
Enter your last name: Vegad
Enter your email: ronak@example.com
Enter number of tickets you want: 2

Booking confirmed for Ronak Vegad!
You booked 2 tickets. 48 remaining.
Current bookings: [Ronak]

🧠 Future Enhancements (Work in Progress)

 Add support for JSON file storage (save and load bookings)

 Add goroutines for concurrent booking simulation

 Add email confirmation simulation

 Introduce a simple web interface (using Go templates)

 Improve input validation (regex-based email validation)

💻 How to Run Locally

Clone this repository:

git clone https://github.com/<your-username>/air-tarsai-booking.git


Navigate into the project folder:

cd Tikcet-booking


Run the program:

go run main.go

📂 Project Structure
Ticket-booking/
│
├── main.go          # Main Go file containing logic
├── README.md        # Project documentation (you’re reading it)
└── go.mod           # (Optional) Go module file if initialized

🧑‍💻 Author

Ronak Vegad
🎓 Engineering IT Student | 💡 Exploring Go, Python, and DevOps
📍 Vallabh Vidyanagar, Gujarat
🔗 GitHub Profile



⭐ Support

If you find this project helpful or want to track updates:

Star ⭐ this repo

Fork it and try extending it yourself

Open issues or suggestions for improvement
