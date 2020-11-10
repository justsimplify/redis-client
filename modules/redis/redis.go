package redis

type Client struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}
