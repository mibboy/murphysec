package utils

import (
	"fmt"
	"strings"
)

func InStringSlice(slice []string, s string) bool {
	for _, it := range slice {
		if it == s {
			return true
		}
	}
	return false
}

func JoinStringAny(sep string, a ...fmt.Stringer) string {
	var s []string
	for i := range a {
		s = append(s, a[i].String())
	}
	return strings.Join(s, sep)
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DistinctStringSlice(s []string) []string {
	m := map[string]struct{}{}
	var rs []string
	for _, it := range s {
		if _, ok := m[it]; !ok {
			rs = append(rs, it)
			m[it] = struct{}{}
		}
	}
	return rs
}

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return map[string]struct{}{}
}
func (s StringSet) Put(t string) StringSet {
	s[t] = struct{}{}
	return s
}
func (s StringSet) Contains(t string) bool {
	_, ok := s[t]
	return ok
}
func (s StringSet) ToSlice() []string {
	var r []string
	for i := range s {
		r = append(r, i)
	}
	return r
}