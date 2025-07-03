package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func saveUser(user User, filename string) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadUser(filename string) (User, error) {
	var load User
	data, err := os.ReadFile(filename)
	if err != nil {
		return load, err
	}
	err = json.Unmarshal(data, &load)
	return load, err

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter the email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	user := User{Name: name, Email: email}

	filename := "user.json"
	if err := saveUser(user, filename); err != nil {
		fmt.Println(err)
	}

	fmt.Println("User saved to...")

	loadUser, err := loadUser(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n loaded user:")
	fmt.Printf("Name: %s\nEmail: %s\n", loadUser.Name, loadUser.Email)
}
