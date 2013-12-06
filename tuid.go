// Copywrite (c) 2013 Clifford Hammerschmidt
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package tuid implements a time prefixed 20 byte uuid. Time prefixed
// uids have an advantage over random uuids when used in systems as
// idenitifiers for entities that are indexed for look up. Because the
// time prefix groups entities created near each other in time the
// updates to the indexed will cluster into a set of 'hot' nodes
// reducing the number of touched nodes in the index.
package tuid

import (
  `crypto/rand`
  `encoding/base32`
  `encoding/binary`
  `errors`
  `time`
)

type Tuid struct {
  t   uint32
  msb uint64
  lsb uint64
}

func randUint64() uint64 {
  b := make([]byte, 8)
  _, err := rand.Read(b)
  if err != nil {
    panic(err)
  }
  return binary.BigEndian.Uint64(b)
}

// Zero is a reference tuid for use similar to nil or null and meant to be used to indicate a no-value state
var Zero = Tuid{}

var epoc = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()

// New returns a new tuid that is statsically unique and will compare as After other guids created more than 1 second previously
func New() Tuid {
  return Tuid{t: uint32(time.Now().UTC().Unix() - epoc), msb: randUint64(), lsb: randUint64()}
}

// FromBytes converts an array of 20 bytes into a tuid or returns an error if the array is not exactly 20 bytes in length
func FromBytes(b []byte) (Tuid, error) {
  if len(b) != 20 {
    return Zero, errors.New(`tuid not 20 bytes in length`)
  }
  return Tuid{t: binary.BigEndian.Uint32(b[0:4]), msb: binary.BigEndian.Uint64(b[4:12]), lsb: binary.BigEndian.Uint64(b[12:20])}, nil
}

// Parse converts the string represented tuid into a tuid or returns an error if the string does not represent a tuid
// use String() to convert a tuid into its string representation
func Parse(encoded string) (Tuid, error) {
  b, err := base32.StdEncoding.DecodeString(encoded)
  if err != nil {
    return Zero, err
  }
  return FromBytes(b)
}

// String returns the tuid encoded as a string suitable for passing to Parse to later recontruct the tuid. The encoding place
// the time based part of the tuid at the front so the guids will sort in semi-ascending order or bucket into temporally
// related buckets when being indexed.
func (t Tuid) String() string {
  return base32.StdEncoding.EncodeToString(t.Bytes())
}

// Bytes returns the 20 bytes representing the tuid with the time first (bytes 0 to 3) in msb order.
func (t Tuid) Bytes() []byte {
  b := make([]byte, 20)
  binary.BigEndian.PutUint32(b[0:4], t.t)
  binary.BigEndian.PutUint64(b[4:12], t.msb)
  binary.BigEndian.PutUint64(b[12:20], t.lsb)
  return b
}

// Before can be used to sort tuids by relative creation time
func (a Tuid) Before(b Tuid) bool {
  if a.t < b.t {
    return true
  }
  if a.t > b.t {
    return false
  }
  if a.msb < b.msb {
    return true
  }
  if a.msb > b.msb {
    return false
  }
  if a.lsb < b.lsb {
    return true
  }
  return false
}

// After can use used to sort tuids by relative creation time
func (a Tuid) After(b Tuid) bool {
  if a.t > b.t {
    return true
  }
  if a.t < b.t {
    return false
  }
  if a.msb > b.msb {
    return true
  }
  if a.msb < b.msb {
    return false
  }
  if a.lsb > b.lsb {
    return true
  }
  return false
}

// Equals determines if two tuids have the same values
func (t Tuid) Equals(u Tuid) bool {
  return (t.t == u.t) && (t.msb == u.msb) && (t.lsb == u.lsb)
}
