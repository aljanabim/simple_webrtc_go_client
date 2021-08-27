package lib

type SignalingChannel struct {
	peerId             string
	peerType           string
	signalingServerUrl string
	token              string
	verbose            bool
}

func (channel SignalingChannel) Connect() {

}
