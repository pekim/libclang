package libclang

// #include <clang-c/Index.h>
import "C"

type AvailabilityKind C.enum_CXAvailabilityKind

func (k AvailabilityKind) String() string {
	return availabilityKindStrings[k]
}

// The entity is available.
const Availability_Available = AvailabilityKind(C.CXAvailability_Available)

// The entity is available, but has been deprecated (and its use is not recommended).
const Availability_Deprecated = AvailabilityKind(C.CXAvailability_Deprecated)

// The entity is not available; any use of it will be an error.
const Availability_NotAvailable = AvailabilityKind(C.CXAvailability_NotAvailable)

// The entity is available, but not accessible; any use of it will be an error.
const Availability_NotAccessible = AvailabilityKind(C.CXAvailability_NotAccessible)

var availabilityKindStrings = map[AvailabilityKind]string{
	Availability_Available:     "Available",
	Availability_Deprecated:    "Deprecated",
	Availability_NotAvailable:  "NotAvailable",
	Availability_NotAccessible: "NotAccessible",
}
