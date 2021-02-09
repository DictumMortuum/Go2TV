package soapcalls

import (
	"log"
	"net/url"

	"github.com/huin/goupnp/soap"
)

type playStopRequest struct {
	InstanceID string
	Speed      string
}

type playStopResponse struct {
}

type setAVTransportRequest struct {
	InstanceID         string
	CurrentURI         string
	CurrentURIMetaData string
}

type setAVTransportResponse struct {
}

// TVPayload - we need this populated in order
type TVPayload struct {
	TransportURL string
	VideoURL     string
	SubtitlesURL string
}

func setAVTransportSoapCall(videoURL, transporturl string) {
	setAVTransportRequest := &setAVTransportRequest{InstanceID: "0", CurrentURI: videoURL}
	setAVTransportResponse := &setAVTransportResponse{}

	parsedURL, err := url.Parse(transporturl)
	if err != nil {
		log.Fatal(err)
	}

	newSetAVTransportURIcall := soap.NewSOAPClient(*parsedURL)
	if err := newSetAVTransportURIcall.PerformAction("urn:schemas-upnp-org:service:AVTransport:1",
		"SetAVTransportURI", setAVTransportRequest, setAVTransportResponse); err != nil {
		log.Fatal(err)
	}

}

// PlayStopSoapCall - Build and call the play soap call
func playStopSoapCall(action, transporturl string) {

	psRequest := &playStopRequest{InstanceID: "0", Speed: "1"}
	psResponse := &playStopResponse{}

	parsedURL, err := url.Parse(transporturl)
	if err != nil {
		log.Fatal(err)
	}

	newPlaycall := soap.NewSOAPClient(*parsedURL)
	if err := newPlaycall.PerformAction("urn:schemas-upnp-org:service:AVTransport:1",
		action, psRequest, psResponse); err != nil {
		log.Fatal(err)
	}
}

// SendtoTV - Send to tv
func (p *TVPayload) SendtoTV(action string) {
	if action == "Play" {
		setAVTransportSoapCall(p.VideoURL, p.TransportURL)
	}
	playStopSoapCall(action, p.TransportURL)

}