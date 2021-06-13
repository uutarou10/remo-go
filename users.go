package remo

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

func (c *Client) GetUser() (User, error) {
	user := User{}
	if err := c.getApi("/1/users/me", &user); err != nil {
		return User{}, err
	}
	return user, nil
}
