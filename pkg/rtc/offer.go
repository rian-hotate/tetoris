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
