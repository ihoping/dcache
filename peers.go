package dcache

import "github.com/ihoping/dcache/dcachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
type PeerGetter interface {
	Get(in *dcachepb.Request, out *dcachepb.Response) error
}
