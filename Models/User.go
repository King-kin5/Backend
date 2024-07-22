package models

type User struct{
	ID          int     `json:"id,omitempty" bson:"_id"`
    Name        string  `json:"name" bson:"name"`
    Email       string  `json:"email"`
   
    Password    string  `json:"password" bson:"password"`
    Credit      float64 `json:"credit" bson:"credit"`
    Area        int     `json:"area" bson:"area"`
    Address     string  `json:"address" bson:"address"`
}
func NewUser(res *User) *User {
    r := new(User)
    r.ID = res.ID
    r.Name = res.Name
    r.Email=res.Email
   
    r.Password = res.Password
    r.Credit = res.Credit
    r.Area = res.Area
    r.Address = res.Address
    return r
}
