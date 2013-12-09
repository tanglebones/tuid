PACKAGE DOCUMENTATION

package tuid
    import "github.com/tanglebones/tuid"

    Package tuid implements a time prefixed 20 byte uuid. Time prefixed uids
    have an advantage over random uuids when used in systems as idenitifiers
    for entities that are indexed for look up. Because the time prefix
    groups entities created near each other in time the updates to the
    indexed will cluster into a set of 'hot' nodes reducing the number of
    touched nodes in the index.

    Example:
	ta := tuid.Zero // placeholder tuid
	fmt.Printf("%v", ta)
	// Output:
	// AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

VARIABLES

var DefaultResolver defaultResolver = defaultResolver{}
    A DefaultResolver is provided because all of the dependencies can be
    resolved using built in system defaults

var DefaultTimeProvider defaultTimeProvider = defaultTimeProvider{}
    A DefaultTimeProvider that uses the built in time package to determine
    seconds from january 1, 2000 00:00:00 GMT (arbitrarily chosen epoc)

var Zero = Tuid{}
    Zero is a reference tuid for use similar to nil or null and meant to be
    used to indicate a no-value state


TYPES

type Resolver interface {
    TimeProvider() TimeProvider
}
    Resolver is used resolve get all the dependencies needed for
    TuidProvider note: no dependency on randomness is given, the use of
    crypto/rand is forced because the naturn of uid generation is tightly
    coupled to randomness



type TimeProvider interface {
    Seconds() uint32
}
    TimeProvider interface for getting Seconds from epoc from an external
    source



type Tuid struct {
    // contains filtered or unexported fields
}


func FromBytes(b []byte) (Tuid, error)
    FromBytes converts an array of 20 bytes into a tuid or returns an error
    if the array is not exactly 20 bytes in length


func Parse(encoded string) (Tuid, error)
    Parse converts the string represented tuid into a tuid or returns an
    error if the string does not represent a tuid use String() to convert a
    tuid into its string representation

    Example:
	tp := tuid.NewTuidProvider(tuid.DefaultResolver)
	tas := tp.New().String()
	ta, err := tuid.Parse(tas)
	if err != nil {
	    fmt.Printf("%v parsed to %v", tas, ta)
	} else {
	    fmt.Printf("%v did not parse", tas)
	}

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


type TuidProvider interface {
    New() Tuid
}
    TuidProvider is used to create tuids

    Example:
	tp := tuid.NewTuidProvider(tuid.DefaultResolver)
	ta := tp.New()
	tb := tp.New() // will be different than ta but have a similar prefix when converted to string
	fmt.Printf("%v %v\n", ta, tb)

func NewTuidProvider(resolver Resolver) TuidProvider
    Constuctor for creating a TuidProvider given a Resolver to resolve the
    dependencies needed.




