package model

import "gonum.org/v1/gonum/stat/distuv"

type Connection struct {
	UUID string
	Mean float64
	Std  float64
	Data StreamData
}

type StreamData struct {
	SessionID  string
	NormalDist *distuv.Normal
	Frequency  float64
	Timestamp  int64
}
