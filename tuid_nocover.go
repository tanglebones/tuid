// Copywrite (c) 2013 Clifford Hammerschmidt
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package tuid

import (
  `crypto/rand`
  `encoding/binary`
)

// not covered because making rand.Read fail would be difficult
// wrapping crypt/rand in an externalized interface would just move the problem there.
func randUint64() uint64 {
  b := make([]byte, 8)
  _, err := rand.Read(b)
  if err != nil {
    panic(err) 
  }
  return binary.BigEndian.Uint64(b)
}
