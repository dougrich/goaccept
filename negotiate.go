package goaccept

import (
	"sort"
	"strconv"
	"strings"
)

// Negotiate breaks up a header, sorts it by quality, and comapres it with acceptable to determine the most suitable content type
func Negotiate(header string, acceptable ...string) (string, error) {
	if header == "" {
		return "", ErrorNotAcceptable{[]RequestedType{}, acceptable}
	}
	var requested requestedSet
	parts := strings.Split(header, ",")
	for _, p := range parts {
		p = strings.Trim(p, " ")
		if p == "" {
			return "", ErrorBadAccept{header}
		}

		quality := 1.0

		subp := strings.SplitN(p, ";", 2)
		if len(subp) == 2 {
			p = subp[0]
			if subp[1][:2] != "q=" {
				return "", ErrorBadAccept{header}
			}
			q, err := strconv.ParseFloat(subp[1][2:], 64)
			if err != nil {
				return "", ErrorBadAccept{header}
			}
			quality = q
		}

		requested = append(requested, RequestedType{quality, p})
	}
	sort.Sort(requested)
	for _, r := range requested {
		for _, a := range acceptable {
			if match(r, a) {
				return a, nil
			}
		}
	}
	return "", ErrorNotAcceptable{requested, acceptable}
}

// This is a requested type after it has been parsed
type RequestedType struct {
	// The quality of this type. The higher, the more preferred this type is.
	Quality float64
	// The mime type; note that this might be a pattern like */*
	MimeType string
}

type requestedSet []RequestedType

func (s requestedSet) Len() int {
	return len(s)
}

func (s requestedSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s requestedSet) Less(i, j int) bool {
	return s[i].Quality > s[j].Quality
}

func match(a RequestedType, b string) bool {
	majorA, minorA := mimeType(a.MimeType)
	majorB, minorB := mimeType(b)
	return (majorA == majorB || majorA == "*" || majorB == "*") && (minorA == minorB || minorA == "*" || minorB == "*")
}

func mimeType(source string) (string, string) {
	parts := strings.SplitN(source, "/", 2)
	return parts[0], parts[1]
}
