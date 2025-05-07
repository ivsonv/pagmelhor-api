package Healthz

import "time"

type GetHealthzResponseDto struct {
	UpdatedAt time.Time      `json:"updated_at"`
	Database  DatabaseStatus `json:"database"`
	Caching   CachingStatus  `json:"caching,omitzero"`
}

type DatabaseStatus struct {
	AvailableConnections int    `json:"available_connections"`
	OpenConnections      int    `json:"open_connections"`
	MaxConnections       int    `json:"max_connections"`
	Version              string `json:"version"`
}

type CachingStatus struct {
	OpenConnections string           `json:"open_connections"`
	Version         string           `json:"version"`
	Stats           CacheStats       `json:"stats"`
	Memory          CacheMemory      `json:"memory"`
	Replication     CacheReplication `json:"replication"`
}

type CacheStats struct {
	InstantaneousOpsPerSec string `json:"instantaneous_ops_per_sec"`
	RejectedConnections    string `json:"rejected_connections"`
}

type CacheMemory struct {
	MaxmemoryPolicy    string `json:"maxmemory_policy"`
	UsedMemoryRssHuman string `json:"used_memory_rss_human"`
	MemoryUsed         string `json:"memory_used"`
	Maxmemory          string `json:"maxmemory"`
}

type CacheReplication struct {
	Role string `json:"role"`
}
