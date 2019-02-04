package rtc

import (
	"github.com/pions/webrtc"
	"github.com/pions/webrtc/examples/util"
)

// RTCSession is to communicate vie webrtc.
type TetorisRTCSession struct {
	addr           string
	peerConnection *webrtc.RTCPeerConnection
	stop           chan bool
}

func newPeerConnection() *webrtc.RTCPeerConnection {
	config := webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.New(config)
	util.Check(err)
	return peerConnection
}

// NewWebRTC is to create web rtc session.
func NewWebRTC() *TetorisRTCSession {
	return &TetorisRTCSession{}
}

func (s *TetorisRTCSession) Init() {
	s.addr = ":50000"
	s.peerConnection = newPeerConnection()
}

func (s *TetorisRTCSession) StopConnection() {
	s.stop <- true
}
