package goqty

import "fmt"

func (q *Qty) Eq(other Qty) bool {
	return q.CompareTo(other) == 0
}
func (q *Qty) Lt(other Qty) bool {
	return q.CompareTo(other) == -1
}
func (q *Qty) Lte(other Qty) bool {
	return q.Eq(other) || q.Lt(other)
}
func (q *Qty) Gt(other Qty) bool {
	return q.CompareTo(other) == 1
}
func (q *Qty) Gte(other Qty) bool {
	return q.Eq(other) || q.Gt(other)
}

// Compare two Qty objects. Throws an exception if they are not of compatible types.
// Comparisons are done based on the value of the quantity in base SI units.
//
// NOTE: We cannot compare inverses as that breaks the general compareTo contract:
//
//	if a.compareTo(b) < 0 then b.compareTo(a) > 0
//	if a.compareTo(b) == 0 then b.compareTo(a) == 0
//
//	Since "10S" == ".1ohm" (10 > .1) and "10ohm" == ".1S" (10 > .1)
//	  Qty("10S").inverse().compareTo("10ohm") == -1
//	  Qty("10ohm").inverse().compareTo("10S") == -1
//
//	If including inverses in the sort is needed, I suggest writing: Qty.sort(qtyArray,units)
func (q *Qty) CompareTo(other Qty) (int, error) {
	if !q.IsCompatible(other) {
		return 0, fmt.Errorf("Incompatible Units %v %v", q.units, other.units)
	}
	if q.baseScalar < other.baseScalar {
		return -1, nil
	} else if q.baseScalar == other.baseScalar {
		return 0, nil
	} else if q.baseScalar > other.baseScalar {
		return 1, nil
	}
}

// Return true if quantities and units match
// Unit("100 cm").same(Unit("100 cm"))  # => true
// Unit("100 cm").same(Unit("1 m"))     # => false
func (q *Qty) Same(other Qty) bool {
	return (q.scalar == other.scalar) && (q.units == other.units)
}
