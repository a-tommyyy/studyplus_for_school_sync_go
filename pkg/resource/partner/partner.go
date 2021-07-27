package partner

type Partner struct {
	PublicId    string `json:"public_id"`
	CustomerUid string `json:"customer_uid"`
	SchoolName  string `json:"school_name"`
	Name        string `json:"name"`
	TimeZone    string `json:"time_zone"`
}
