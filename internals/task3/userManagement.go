package task3

import "fmt"

func DisplayUserManagement() {
	userManager := NewUserManager()
	user1 := &User{ID: 1, Name: "Farid"}
	user2 := &User{ID: 2, Name: "Rhamadhan"}

	user3 := &User{ID: 1, Name: "Darari"} // kalo ID sama

	if err := userManager.AddUser(user1); err != nil {
		fmt.Println(err)
	}

	if err := userManager.AddUser(user2); err != nil {
		fmt.Println(err)
	}

	if err := userManager.AddUser(user3); err != nil {
		fmt.Println(err)
	}

	user, err := userManager.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User dengan ID: %d ditemukan\n", user.ID)
	}

	user, err = userManager.GetUser(10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User dengan ID: %d ditemukan\n", user.ID)
	}
}
