package cty

// myuserAsNull returns a value of the same type as the given value but
// with any myuser values (including nested values) replaced with null
// values of the same type.
//
// This can be useful if a result is to be serialized in a format that can't
// represent myusers, such as JSON, as long as the caller does not need to
// retain the myuser value information.
func myuserAsNull(val Value) Value {
	ty := val.Type()
	switch {
	case val.IsNull():
		return val
	case !val.IsKnown():
		return NullVal(ty)
	case ty.IsListType() || ty.IsTupleType() || ty.IsSetType():
		length := val.LengthInt()
		if length == 0 {
			// If there are no elements then we can't have myusers
			return val
		}
		vals := make([]Value, 0, length)
		it := val.ElementIterator()
		for it.Next() {
			_, v := it.Element()
			vals = append(vals, myuserAsNull(v))
		}
		switch {
		case ty.IsListType():
			return ListVal(vals)
		case ty.IsTupleType():
			return TupleVal(vals)
		default:
			return SetVal(vals)
		}
	case ty.IsMapType() || ty.IsObjectType():
		var length int
		switch {
		case ty.IsMapType():
			length = val.LengthInt()
		default:
			length = len(val.Type().AttributeTypes())
		}
		if length == 0 {
			// If there are no elements then we can't have myusers
			return val
		}
		vals := make(map[string]Value, length)
		it := val.ElementIterator()
		for it.Next() {
			k, v := it.Element()
			vals[k.AsString()] = myuserAsNull(v)
		}
		switch {
		case ty.IsMapType():
			return MapVal(vals)
		default:
			return ObjectVal(vals)
		}
	}

	return val
}
