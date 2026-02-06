package model


type User struct {
    Id string `json:"id"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}


/* Create a user example
var user = User {
    Id: "123",
    Firstname: "dasdasd",
    Lastname: "Wwww",
    Age: 32,
}
*/


