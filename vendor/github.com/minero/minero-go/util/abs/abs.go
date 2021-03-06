// Defines absolute to real and viceversa position and look translators.
package abs

// Pos translates real positions to its absolute form.
func Pos(i float64) int32 {
	return int32(i * 32)
}

// Look translates real pitch and yaw to its absolute form.
func Look(l float32) int8 {
	return int8(int(l / 180.0 * 128))
}

// RealPos translates absolute positions to its real form.
func RealPos(i int32) float64 {
	return float64(i << 5)
}

// RealLook translates absolute pitch and yaw to its real form.
func RealLook(l int8) float32 {
	return float32(l) / 128 * 180.0
}
