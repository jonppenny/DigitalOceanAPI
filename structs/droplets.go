package structs

import "time"

type DropletsData struct {
	Droplets []Droplets `json:"droplets"`
	Links    Links      `json:"links"`
	Meta     Meta       `json:"meta"`
}
type NextBackupWindow struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
type Image struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Distribution  string    `json:"distribution"`
	Slug          any       `json:"slug"`
	Public        bool      `json:"public"`
	Regions       []any     `json:"regions"`
	CreatedAt     time.Time `json:"created_at"`
	MinDiskSize   int       `json:"min_disk_size"`
	Type          string    `json:"type"`
	SizeGigabytes float64   `json:"size_gigabytes"`
	Description   string    `json:"description"`
	Tags          []any     `json:"tags"`
	Status        string    `json:"status"`
}
type Size struct {
	Slug                string   `json:"slug"`
	Memory              int      `json:"memory"`
	Vcpus               int      `json:"vcpus"`
	Disk                int      `json:"disk"`
	Transfer            float64  `json:"transfer"`
	PriceMonthly        float64  `json:"price_monthly"`
	PriceHourly         float64  `json:"price_hourly"`
	Regions             []string `json:"regions"`
	Available           bool     `json:"available"`
	Description         string   `json:"description"`
	NetworkingThrougput int      `json:"networking_througput"`
}
type V4 struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}
type Networks struct {
	V4 []V4  `json:"v4"`
	V6 []any `json:"v6"`
}
type Region struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Features  []string `json:"features"`
	Available bool     `json:"available"`
	Sizes     []string `json:"sizes"`
}
type Droplets struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Memory           int              `json:"memory"`
	Vcpus            int              `json:"vcpus"`
	Disk             int              `json:"disk"`
	Locked           bool             `json:"locked"`
	Status           string           `json:"status"`
	Kernel           any              `json:"kernel"`
	CreatedAt        time.Time        `json:"created_at"`
	Features         []string         `json:"features"`
	BackupIds        []int            `json:"backup_ids"`
	NextBackupWindow NextBackupWindow `json:"next_backup_window"`
	SnapshotIds      []any            `json:"snapshot_ids"`
	Image            Image            `json:"image"`
	VolumeIds        []any            `json:"volume_ids"`
	Size             Size             `json:"size"`
	SizeSlug         string           `json:"size_slug"`
	Networks         Networks         `json:"networks"`
	Region           Region           `json:"region"`
	Tags             []any            `json:"tags"`
}
type Links struct {
}
type Meta struct {
	Total int `json:"total"`
}
