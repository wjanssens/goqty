package qty

import "fmt"

func (q *Qty) Eq(other interface{}) (bool, error) {
	if i, err := q.CompareTo(other); err != nil {
		return false, err
	} else if i == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func (q *Qty) Lt(other interface{}) (bool, error) {
	if i, err := q.CompareTo(other); err != nil {
		return false, err
	} else if i == -1 {
		return true, nil
	} else {
		return false, nil
	}
}
func (q *Qty) Lte(other interface{}) (bool, error) {
	if eq, err := q.Eq(other); err != nil {
		return false, err
	} else if lt, err := q.Lt(other); err != nil {
		return false, err
	} else {
		return eq || lt, nil
	}
}
func (q *Qty) Gt(other interface{}) (bool, error) {
	if i, err := q.CompareTo(other); err != nil {
		return false, err
	} else if i == 1 {
		return true, nil
	} else {
		return false, nil
	}
}
func (q *Qty) Gte(other interface{}) (bool, error) {
	if eq, err := q.Eq(other); err != nil {
		return false, err
	} else if gt, err := q.Gt(other); err != nil {
		return false, err
	} else {
		return eq || gt, nil
	}
}

// Compare two Qty objects. Throws an exception if they are not of compatible types.
// Comparisons are done based on the value of the quantity in base SI units.
//
// NOTE: Cannot compare inverses as that breaks the general compareTo contract:
//
//	if a.CompareTo(b) < 0 then b.CompareTo(a) > 0
//	if a.CompareTo(b) == 0 then b.CompareTo(a) == 0
//
//	Since "10S" == ".1ohm" (10 > .1) and "10ohm" == ".1S" (10 > .1)
//	  Qty("10S").Inverse().CompareTo(Parse("10ohm")) == -1
//	  Qty("10ohm").Inverse().CompareTo(Parse("10S")) == -1
func (q *Qty) CompareTo(other interface{}) (int, error) {
	var o *Qty
	var err error
	switch t := other.(type) {
	case *Qty:
		o = other.(*Qty)
	case string:
		if o, err = Parse(other.(string)); err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("expecting string or *Qty, got %T", t)
	}

	if !q.IsCompatible(o) {
		return 0, fmt.Errorf("incompatible units: %v and %v", q.Units(), o.Units())
	}
	if q.baseScalar < o.baseScalar {
		return -1, nil
	} else if q.baseScalar > o.baseScalar {
		return 1, nil
	} else {
		return 0, nil
	}
}

// Return true if quantities and units match
// Unit("100 cm").Same(Unit("100 cm"))  # => true
// Unit("100 cm").Same(Unit("1 m"))     # => false
func (q *Qty) Same(other interface{}) (bool, error) {
	var o *Qty
	var err error
	switch t := other.(type) {
	case *Qty:
		o = other.(*Qty)
	case string:
		if o, err = Parse(other.(string)); err != nil {
			return false, err
		}
	default:
		return false, fmt.Errorf("expecting string or *Qty, got %T", t)
	}
	return (q.scalar == o.scalar) && (q.units == o.units), nil
}
