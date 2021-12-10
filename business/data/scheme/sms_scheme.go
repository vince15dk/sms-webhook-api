package scheme

type RecipientList struct {
	RecipientNo              string `json:"recipientNo"`
	CountryCode              string `json:"countryCode"`
	InternationalRecipientNo string `json:"internationalRecipientNo,omitempty"`
	RecipientGroupingKey     string `json:"recipientGroupingKey,omitempty"`
}

type SMSRequest struct {
	TemplateID        string          `json:"templateId,omitempty"`
	Body              string          `json:"body"`
	SendNo            string          `json:"sendNo"`
	RequestDate       string          `json:"requestDate,omitempty"`
	SenderGroupingKey string          `json:"senderGroupingKey,omitempty"`
	RecipientList     []RecipientList `json:"recipientList"`
	UserID            string          `json:"userId,omitempty"`
	StatsID           string          `json:"statsId,omitempty"`
}
