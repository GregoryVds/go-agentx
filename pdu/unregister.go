package pdu

import (
	"github.com/posteo/go-agentx/marshaler"
	"gopkg.in/errgo.v1"
)

// Unregister defines the pdu unregister packet.
type Unregister struct {
	Timeout Timeout
	Subtree ObjectIdentifier
}

// Type returns the pdu packet type.
func (u *Unregister) Type() Type {
	return TypeUnregister
}

// MarshalBinary returns the pdu packet as a slice of bytes.
func (u *Unregister) MarshalBinary() ([]byte, error) {
	combined := marshaler.NewMulti(&u.Timeout, &u.Subtree)

	combinedBytes, err := combined.MarshalBinary()
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return combinedBytes, nil
}

// UnmarshalBinary sets the packet structure from the provided slice of bytes.
func (u *Unregister) UnmarshalBinary(data []byte) error {
	return nil
}