package rtc

import (
	"testing"

	"github.com/rian-hotate/tetoris/internal/rtc"
)

func TestASuccess(t *testing.T) {
	go func() {
		s := rtc.NewWebRTC()
		s.Init()
		s.Answer()
	}()

	go func() {
		c := rtc.NewWebRTC()
		c.Init()
		c.Offer()
	}()

}

func TestAFailed(t *testing.T) {
}
