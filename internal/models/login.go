package models

type UserLogin struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Lastname   string `json:"lastname,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type UpdateUserLogin struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Lastname    string `json:"lastname,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
}
