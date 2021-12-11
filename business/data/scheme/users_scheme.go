package scheme

type DepUsers struct {
	AppKey    string  `json:"appKey"`
	SecretKey string  `json:"secretKey"`
	DepGroup  string  `json:"depGroup"`
	Users     []Users `json:"users"`
}

type Users struct {
	Name    string `json:"name,omitempty"`
	PhoneNo string `json:"phoneNo"`
	Email   string `json:"email,omitempty"`
}
