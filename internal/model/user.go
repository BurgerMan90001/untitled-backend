package model


type User struct {
    Id string `json:"id"`
    Username  string `json:"name"`
    Email string `json:"email"`
}


/* Create a user example
var user = User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}
*/


