package tuid

import (
  `testing`
  `time`
)

func TestEquals(t *testing.T) {
  if !Zero.Equals(Zero) {
    t.Errorf(`Zero != Zero`)
  }
  a := New()
  if !a.Equals(a) {
    as := a.String()
    t.Errorf(`%v != %v`, as, as)
  }
}

func TestZero(t *testing.T) {
  zs := Zero.String()
  ns := New().String()
  t.Logf(`zero is: %v `, zs)
  t.Logf(`new is: %v`, ns)
  const expectedLen = 32
  if len(zs) != expectedLen {
    t.Errorf(`tuid.Zero is not length %v`, expectedLen)
  }
  if len(ns) != expectedLen {
    t.Errorf(`tuid.New() is not length %v`, expectedLen)
  }
  if ns == zs {
    t.Errorf(`tuid.New() == tuid.Zero`)
  }
}

func TestParse(t *testing.T) {
  a := New()
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
}

func (tuid Tuid) dump(t *testing.T) {
  t.Logf(`%#v`, tuid)
}

func TestTuidsAreIncreasingOverTime(t *testing.T) {
  a := New()
  time.Sleep(1100 * time.Millisecond)
  b := New()
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
    tuid := New()
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
