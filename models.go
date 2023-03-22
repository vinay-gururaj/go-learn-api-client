package kwclient

type XPProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SQSConnection struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	ReadQueue  string `json:"readQueue"`
	WriteQueue string `json:"writeQueue"`
	ProjectId  string `json:"projectId"`
}
