//go:build !android && !darwin && !js && !windows
// +build !android,!darwin,!js,!windows

package ahsai

import "errors"

// PlayOgg cut leading demo text and play directly
func PlayOgg(u string) error {
	return errors.New("cannot play directly on this platform")
}
