package scheme

// DepUsers Department/Organization Group
type DepUsers struct {
	DepGroup []DepGroup `json:"depGroup"`
}

// DepGroup for DepUsers
type DepGroup struct {
	AppKey    string  `json:"appKey"`
	SecretKey string  `json:"secretKey"`
	GroupName string  `json:"groupName"`
	Sender    string  `json:"sender"`
	Users     []Users `json:"users"`
}

// Users for DepGroup
type Users struct {
	Name    string `json:"name,omitempty"`
	PhoneNo string `json:"phoneNo"`
	Email   string `json:"email,omitempty"`
}
