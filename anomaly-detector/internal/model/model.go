package model

type DetectorStats struct {
	Id                 string
	Mean               float64
	Std                float64
	ValuesProcessed    int
	AnomalyCoefficient float64
	AnomalyDetected    bool
	SumSquares         float64
	DetectionMode      bool
}
