package model

type Anomaly struct {
	SessionID    string  `gorm:"column:session_id"`
	Frequency    float64 `gorm:"column:frequency"`
	TimestampUTC int64   `gorm:"column:timestamp_utc"`
}
