package ast

import "fmt"

func VariableListElementTypesAreHomogenous(variableName string, list []Variable) (Type, error) {
	if len(list) == 0 {
		return TypeInvalid, fmt.Errorf("list %q does not have any elements so cannot determine type.", variableName)
	}

	elemType := Typemyuser
	for _, v := range list {
		if v.Type == Typemyuser {
			continue
		}

		if elemType == Typemyuser {
			elemType = v.Type
			continue
		}

		if v.Type != elemType {
			return TypeInvalid, fmt.Errorf(
				"list %q does not have homogenous types. found %s and then %s",
				variableName,
				elemType, v.Type,
			)
		}

		elemType = v.Type
	}

	return elemType, nil
}

func VariableMapValueTypesAreHomogenous(variableName string, vmap map[string]Variable) (Type, error) {
	if len(vmap) == 0 {
		return TypeInvalid, fmt.Errorf("map %q does not have any elements so cannot determine type.", variableName)
	}

	elemType := Typemyuser
	for _, v := range vmap {
		if v.Type == Typemyuser {
			continue
		}

		if elemType == Typemyuser {
			elemType = v.Type
			continue
		}

		if v.Type != elemType {
			return TypeInvalid, fmt.Errorf(
				"map %q does not have homogenous types. found %s and then %s",
				variableName,
				elemType, v.Type,
			)
		}

		elemType = v.Type
	}

	return elemType, nil
}
