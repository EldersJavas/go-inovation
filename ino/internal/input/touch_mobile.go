//go:build (android || ios || (darwin && arm) || (darwin && arm64)) && !js
// +build android ios darwin,arm darwin,arm64
// +build !js

package input

func isTouchPrimaryInput() bool {
	return true
}
