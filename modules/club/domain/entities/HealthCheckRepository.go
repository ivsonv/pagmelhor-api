package entities

type HealthCheckRepository struct {
	AvailableConnections int    `json:"available_connections"`
	OpenConnections      int    `json:"open_connections"`
	MaxConnections       int    `json:"max_connections"`
	Version              string `json:"version"`
}
