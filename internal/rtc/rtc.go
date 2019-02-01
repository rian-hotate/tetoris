package rtc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

func (s *TetorisRTCSession) Init() {
	s.addr = ":50000"
	s.peerConnection = newPeerConnection()
}

func (s *TetorisRTCSession) StopConnection() {
	s.stop <- true
}
func (s *TetorisRTCSession) Answer() {
	s.peerConnection.OnICEConnectionStateChange(func(connectionState ice.ConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})
	s.peerConnection.OnDataChannel(func(d *webrtc.RTCDataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label, d.ID)
		d.OnOpen(func() {
		})

		// Register message handling
		d.OnMessage(func(payload datachannel.Payload) {
		})
	})

	offerChan, answerChan := answerSignalViaHTTP(":50000")

	offer := <-offerChan

	err := s.peerConnection.SetRemoteDescription(offer)
	util.Check(err)

	answer, err := s.peerConnection.CreateAnswer(nil)
	util.Check(err)

	answerChan <- answer

	<-s.stop
}

func answerSignalViaHTTP(address string) (offerOut chan webrtc.RTCSessionDescription, answerIn chan webrtc.RTCSessionDescription) {
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

func (s *TetorisRTCSession) Offer() {
	dataChannel, err := s.peerConnection.CreateDataChannel("data", nil)
	util.Check(err)

	s.peerConnection.OnICEConnectionStateChange(func(connectionState ice.ConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	dataChannel.OnOpen(func() {
	})

	dataChannel.OnMessage(func(payload datachannel.Payload) {
	})

	offer, err := s.peerConnection.CreateOffer(nil)
	util.Check(err)

	answer := offerSignalViaHTTP(offer, ":50000")

	err = s.peerConnection.SetRemoteDescription(answer)
	util.Check(err)

	<-s.stop

}

func offerSignalViaHTTP(offer webrtc.RTCSessionDescription, address string) webrtc.RTCSessionDescription {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(offer)
	util.Check(err)

	resp, err := http.Post("http://"+address, "application/json; charset=utf-8", b)
	util.Check(err)
	defer resp.Body.Close()

	var answer webrtc.RTCSessionDescription
	err = json.NewDecoder(resp.Body).Decode(&answer)
	util.Check(err)

	return answer
}
