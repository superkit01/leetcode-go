package middle

func AngleClock(hour int, minutes int) float64 {
	minAngle := float64(minutes * 6)
	hourAngle := float64(hour*30) + float64(minutes)/float64(60)*30
	angle := max(hourAngle, minAngle) - min(hourAngle, minAngle)

	return min(360.0-angle, angle)
}
