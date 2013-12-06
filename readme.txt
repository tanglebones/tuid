PACKAGE DOCUMENTATION

package tuid
    import "github.com/tanglebones/tuid"



VARIABLES

var Zero = Tuid{}
    Zero is a reference tuid for use similar to nil or null and meant to be
    used to indicate a no-value state


TYPES

type Tuid struct {
    // contains filtered or unexported fields
}


func FromBytes(b []byte) (Tuid, error)
    FromBytes converts an array of 20 bytes into a tuid or returns an error
    if the array is not exactly 20 bytes in length


func New() Tuid
    New returns a new tuid that is statsically unique and will compare as
    After other guids created more than 1 second previously


func Parse(encoded string) (Tuid, error)
    Parse converts the string represented tuid into a tuid or returns an
    error if the string does not represent a tuid use String() to convert a
    tuid into its string representation


func (a Tuid) After(b Tuid) bool
    After can use used to sort tuids by relative creation time

func (a Tuid) Before(b Tuid) bool
    Before can be used to sort tuids by relative creation time

func (t Tuid) Bytes() []byte
    Bytes returns the 20 bytes representing the tuid with the time first
    (bytes 0 to 3) in msb order.

func (t Tuid) Equals(u Tuid) bool
    Equals determines if two tuids have the same values

func (t Tuid) String() string
    String returns the tuid encoded as a string suitable for passing to
    Parse to later recontruct the tuid. The encoding place the time based
    part of the tuid at the front so the guids will sort in semi-ascending
    order or bucket into temporally related buckets when being indexed.



