package check

// maps user id to list of checks
var checksDB map[int][]Check

type Check struct {
	ID             string // Auto-generated unique ID
	Name           string
	Url            string
	Protocol       string //HTTP, HTTPS, TCP
	Path           string
	Port           string
	WebhookUrl     string
	Timeout        int32 //millisecond
	Interval       int32
	Threshold      int
	Authentication string
	HttpHeaders    map[string]string
	Assert         []string
	Tags           []string
	IgnoreSSL      bool //True--> ignore
	UserId         int  // Foreign key to associate with a User
}

type User struct {
	ID int
}

// get check/s by user id
func GetCheckByID(userId int) []Check {
	return checksDB[userId]
}

func AddCheck(check Check) {
	checksDB[check.UserId] = append(checksDB[check.UserId], check)
}

func DeleteCheck(userId int, checkId string) {
	userChecks := checksDB[userId]

	for i, check := range userChecks {
		if check.ID == checkId {
			userChecks = append(userChecks[:i], userChecks[i+1:]...)
			checksDB[userId] = userChecks
			return

		}

	}
}
