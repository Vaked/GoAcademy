package stringset

import (
	"fmt"
	"strings"
	"reflect"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New(strings ...string) Set {
	set := make(map[string]struct{}, len(strings))
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func NewFromSlice(l []string) Set {
	return New(l...)
}

func (s Set) String() string {
	val := make([]string, 0, len(s))
	if len(s) == 0{
		return "{}"
	}
	for item := range s{
		val = append(val, "\"" +string(item) + "\"")
	}
	return fmt.Sprintf("{" + "%s", strings.Join(val, ", ") + "}")
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2){
		return false
	}

	for key1, val1 := range s1{
		if val2, found := s2[key1]; !found || val2 != val1 {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	if len(s1) < len(s2) || len(s2) < len(s1) {
		return true
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	for k1 := range s1{
		for k2 := range s2{
			if k1 == k2{
				return false
			}
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2){
		return false
	}

	if len(s1) == 0 && len(s2) == 0{
		return true
	}

	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	set := New()

	if len(s1) == 0 || len(s2) == 0 {
		return set
	}

	for k, v := range s2{
		_, ok := s1[k]
		if ok {
			set[k] = v
		}
	}

	return set
}

func Difference(s1, s2 Set) Set {
	set := New()

	if len(s1) == 0 && len(s2) == 0{
		return set
	} else if len(s1) == 0 && len(s2) > 0{
		return set
	} else if len(s1) > 0 && len(s2) == 0{
		return s1
	}

	for k1, v := range s1{
		for k2 := range s2{
			if k1 == k2{
				set[k1] = v
			}
		}
	}

	result := New()
	for k, v := range s1{
		for k2 := range set{
			if k != k2{
				result[k] = v
			}
		}
	}
	return result
	
}

func Union(s1, s2 Set) Set {
	set := New()

	for key, val := range s1{
		set[key] = val
	}

	for key, val := range s2{
		set[key] = val
	}

	return set
}