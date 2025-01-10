package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rickj1ang/RRS_Client/cmd"
)

var loggedIn bool
var authToken string

func main() {
	fmt.Println("Welcome to the RRS")
	reader := bufio.NewReader(os.Stdin)

	for {
		if !loggedIn {
			login(reader)
		} else {
			showMenu(reader)
		}
	}
}

func login(reader *bufio.Reader) {
	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	token, err := cmd.Login(email, password)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		return
	}

	authToken = token
	loggedIn = true
	fmt.Println("Login successful!")
}

func showMenu(reader *bufio.Reader) {
	fmt.Println("\nPlease choose an option:")
	fmt.Println("1. Show all records")
	fmt.Println("2. Record reading")
	fmt.Println("3. Recommand book")
	fmt.Println("4. Add a book")
	fmt.Println("5. Delete user")
	fmt.Println("6. Logout")
	fmt.Println("7. Exit")

	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		getAllRecords()
	case "2":
		getUserByID(reader)
	case "3":
		createUser(reader)
	case "4":
		updateUser(reader)
	case "5":
		deleteUser(reader)
	case "6":
		logout()
	case "7":
		fmt.Println("Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func logout() {
	loggedIn = false
	authToken = ""
	fmt.Println("Logged out successfully.")
}

func getAllRecords() {
	response, err := cmd.Get("/records", authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(response)
}

func getUserByID(reader *bufio.Reader) {
	fmt.Print("Enter user ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	response, err := cmd.Get("/users/"+id, authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(response)
}

func createUser(reader *bufio.Reader) {
	fmt.Print("Enter user name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter user email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	userData := fmt.Sprintf(`{"name": "%s", "email": "%s"}`, name, email)
	response, err := cmd.Post("/users", userData, authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(response)
}

func updateUser(reader *bufio.Reader) {
	fmt.Print("Enter user ID to update: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Enter new user name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new user email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	userData := fmt.Sprintf(`{"name": "%s", "email": "%s"}`, name, email)
	response, err := cmd.Put("/users/"+id, userData, authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(response)
}

func deleteUser(reader *bufio.Reader) {
	fmt.Print("Enter user ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	response, err := cmd.Delete("/users/"+id, authToken)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(response)
}
