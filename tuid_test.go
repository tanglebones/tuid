package tuid

import (
  `testing`
)

var tp TuidProvider = NewTuidProvider(DefaultResolver)

func TestEquals(t *testing.T) {
  if !Zero.Equals(Zero) {
    t.Errorf(`Zero != Zero`)
  }
  a := tp.New()
  if !a.Equals(a) {
    as := a.String()
    t.Errorf(`%v != %v`, as, as)
  }
}

func TestZero(t *testing.T) {
  zs := Zero.String()
  ns := tp.New().String()
  t.Logf(`zero is: %v `, zs)
  t.Logf(`new is: %v`, ns)
  const expectedLen = 32
  if len(zs) != expectedLen {
    t.Errorf(`tuid.Zero is not length %v`, expectedLen)
  }
  if len(ns) != expectedLen {
    t.Errorf(`tuid from New() is not length %v`, expectedLen)
  }
  if ns == zs {
    t.Errorf(`tuid from New() == tuid.Zero`)
  }
}

func TestParse(t *testing.T) {
  a := tp.New()
  as := a.String()
  b, err := Parse(as)
  a.dump(t)
  b.dump(t)
  if err != nil {
    t.Errorf(`string %v could not be parsed`, as)
  }
  if !a.Equals(b) {
    t.Errorf(`tuid %v != %v`, as, b.String())
  }
  _, err = Parse("invalid tuid")
  if err == nil {
    t.Errorf(`expected Prase to fail on invalid input`)
  }
}

func TestBytesFailure(t *testing.T) {
  bytes19 := make([]byte,19)
  _, err := FromBytes(bytes19)
  if err == nil {
    t.Errorf(`expected FromBytes() to fail when given 19 bytes`)
  }
  bytes21 := make([]byte,21)
  _, err = FromBytes(bytes21)
  if err == nil {
    t.Errorf(`expected FromBytes() to fail when given 21 bytes`)
  }
}

func (tuid Tuid) dump(t *testing.T) {
  t.Logf(`%#v`, tuid)
}

type mockTimeProvider struct{}
type mockTuidResolver struct{}

var mockTime uint32 = 0

func (_ mockTimeProvider) Seconds() uint32 {
  mockTime++
  return mockTime
}

func (_ mockTuidResolver) TimeProvider() TimeProvider {
  return mockTimeProvider{}
}

func TestTuidsAreIncreasingOverTime(t *testing.T) {
  mtp := NewTuidProvider(mockTuidResolver{})
  a := mtp.New()
  b := mtp.New()
  a.dump(t)
  b.dump(t)
  if a.After(b) {
    t.Errorf(`tuids not in ascending order: a.After(b) `)
  }
  if !(b.After(a)) {
    t.Errorf(`tuids not in ascending order: !b.After(a)`)
  }
  if b.Before(a) {
    t.Errorf(`tuids not in ascending order: b.Before(a)`)
  }
  if !(a.Before(b)) {
    t.Errorf(`tuids not in ascending order: !a.Before(b)`)
  }
}

func TestTuidsAreDifferent(t *testing.T) {
  tuids := map[Tuid]bool{}
  prev := Zero

  for howMany := 100; howMany >= 0; howMany -= 1 {
    tuid := tp.New()
    if _, exists := tuids[tuid]; exists {
      t.Errorf(`tuid duplicate found.`)
    }
    tuids[tuid] = true
    if tuid.t < prev.t {
      t.Errorf(`newer tuid %#v has t less previous tuid %#v`, tuid, prev)
    }
    prev = tuid
  }
}

func TestInterferencePatterns(t *testing.T) {
  allOnes := byte(0xFF)
  altOdd := byte(0x55)  // 01010101b
  altEven := byte(0xAA) // 10101010b

  for _, src := range []byte{allOnes, altOdd, altEven} {
    bytes := make([]byte, 20)
    for i := 0; i < 20; i++ {
      bytes[i] = src
    }
    t1, err1 := FromBytes(bytes)
    s1 := t1.String()
    t2, err2 := Parse(s1)
    s2 := t2.String()
    t1.dump(t)
    t2.dump(t)
    t.Logf(`%v %v`, s1, s2)
    if s1 != s2 {
      t.Errorf(`Encoding failed`)
    }
    if err1 != nil {
      t.Errorf(`err1 = %v`, err1)
    }
    if err2 != nil {
      t.Errorf(`err2 = %v`, err1)
    }
  }
}
