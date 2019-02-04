package rtc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pions/webrtc"
	"github.com/pions/webrtc/examples/util"
	"github.com/pions/webrtc/pkg/datachannel"
	"github.com/pions/webrtc/pkg/ice"
)

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
