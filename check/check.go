package check

type Check struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	Url            string            `json:"url"`
	Protocol       string            `json:"protocol"` //HTTP, HTTPS, TCP
	Path           string            `json:"path"`
	Port           string            `json:"port"`
	WebhookUrl     string            `json:"webhook_url"`
	Timeout        int32             `json:"time_out"` //millisecond
	Interval       int32             `json:"time_interval"`
	Threshold      int               `json:"threshold"`
	Authentication string            `json:"authentication"`
	HttpHeaders    map[string]string `json:"httpHeaders"`
	Assert         []string          `json:"Assert"`
	Tags           []string          `json:"Tags"`
	IgnoreSSL      bool              `json:"IgnoreSSL"` //True--> ignore
}

type User struct {
	Id int `json:"id"`
}

type DeleteCheck struct {
	CheckId int `json:"check_id"`
}

var arrayCheck map[int][]Check

func GetCheckByID(userId int) []Check {
	return arrayCheck[userId]
}
