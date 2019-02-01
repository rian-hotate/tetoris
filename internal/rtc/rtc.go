package rtc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pions/webrtc"
	"github.com/pions/webrtc/examples/util"
	"github.com/pions/webrtc/pkg/datachannel"
	"github.com/pions/webrtc/pkg/ice"
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

func (s *TetorisRTCSession) init() {
	s.addr = ":50000"
	s.peerConnection = newPeerConnection()
}

func (s *TetorisRTCSession) offer() {

}
func (s *TetorisRTCSession) stopConnection() {
	s.stop <- true
}
func (s *TetorisRTCSession) answer() {
	s.peerConnection.OnICEConnectionStateChange(func(connectionState ice.ConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})
	s.peerConnection.OnDataChannel(func(d *webrtc.RTCDataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label, d.ID)
		d.OnOpen(func() {
			}
		})

		// Register message handling
		d.OnMessage(func(payload datachannel.Payload) {
		})
	})

	offerChan, answerChan := mustSignalViaHTTP(*s.addr)

	offer := <-offerChan

	err := s.peerConnection.SetRemoteDescription(offer)
	util.Check(err)

	answer, err := s.peerConnection.CreateAnswer(nil)
	util.Check(err)

	answerChan <- answer

	<-s.stop
}

func mustSignalViaHTTP(address string) (offerOut chan webrtc.RTCSessionDescription, answerIn chan webrtc.RTCSessionDescription) {
	offerOut = make(chan webrtc.RTCSessionDescription)
	answerIn = make(chan webrtc.RTCSessionDescription)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var offer webrtc.RTCSessionDescription
		err := json.NewDecoder(r.Body).Decode(&offer)
		util.Check(err)

		offerOut <- offer
		answer := <-answerIn

		err = json.NewEncoder(w).Encode(answer)
		util.Check(err)

	})

	go http.ListenAndServe(address, nil)
	fmt.Println("Listening on", address)

	return
}
