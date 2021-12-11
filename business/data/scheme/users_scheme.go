package scheme

type DepUsers struct {
	DepGroup []DepGroup `json:"depGroup"`
}

type DepGroup struct {
	AppKey    string  `json:"appKey"`
	SecretKey string  `json:"secretKey"`
	GroupName string  `json:"groupName"`
	Sender    string  `json:"sender"`
	Users     []Users `json:"users"`
}

type Users struct {
	Name    string `json:"name,omitempty"`
	PhoneNo string `json:"phoneNo"`
	Email   string `json:"email,omitempty"`
}
