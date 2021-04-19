package pointerRel

import "fmt"

/*

*[]Users 是一个指向切片的指针Users
[]*Users 是一个切片,且每个元素指向Users

*/

type Users struct {
	ID   int
	Name string
}

func PointerTest1() {
	var (
		user1 Users
		user2 Users
	)
	//Make a couple Users:
	user1 = Users{ID: 43215, Name: "Billy"}
	user2 = Users{ID: 84632, Name: "Bobby"}

	//Then make a list of pointers to those Users:
	var userList []*Users = []*Users{&user1, &user2}

	for i, userPrt := range userList {
		fmt.Printf("i: %d; userPrt: %v; user: %v\n", i, userPrt, *userPrt)
	}

	//Now you can change an individual Users in that list.
	//This changes the variable user2:
	*userList[1] = Users{ID: 1337, Name: "Larry"}

	fmt.Println(user1) // Outputs: {43215 Billy}
	fmt.Println(user2) // Outputs: {1337 Larry}
}

func PointerTest2() {
	var (
		userList []Users
	)
	//Make the slice of Users
	userList = []Users{Users{ID: 43215, Name: "Billy"}}

	//Then pass the slice as a reference to some function
	myFunc(&userList)

	fmt.Println(userList) // Outputs: [{1337 Bobby}]
}

//Now the function gets a pointer *[]Users that when changed, will affect the global variable "userList"
func myFunc(input *[]Users) {
	*input = []Users{Users{ID: 1337, Name: "Bobby"}}
	for i, user := range *input {
		fmt.Printf("i: %d; user: %v\n", i, user)
	}
}
