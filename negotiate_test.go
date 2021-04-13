package goaccept_test

import (
	"fmt"
	"testing"

	accept "github.com/dougrich/goaccept"
	"github.com/stretchr/testify/assert"
)

func ExampleAccept() {
	header := "text/plain, text/html, text/*;q=0.8, */*;q=0.5"

	contentType, err := accept.Negotiate(header, "text/html")
	if err != nil {
		switch err.(type) {
		case accept.ErrorNotAcceptable:
			// return a 406 or set the default content type
			panic(err)
		case accept.ErrorBadAccept:
			// return a 400; they passed an improperly formatted accept header
			panic(err)
		default:
			// this shouldn't happen; return a 500
			panic(err)
		}
	}

	fmt.Println(contentType)
	// Output: text/html
}

func TestAccept(t *testing.T) {
	p := func(name string, contentType string, err error, header string, acceptable ...string) {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			actualContentType, actualError := accept.Negotiate(header, acceptable...)
			assert.Equal(contentType, actualContentType)
			assert.Equal(err, actualError)
		})
	}

	p("Basic", "text/plain", nil, "text/plain", "text/plain")
	p("NotAcceptableMissing", "", accept.ErrorNotAcceptable{[]accept.RequestedType{}, []string{"text/html"}}, "", "text/html")
	p("NotAcceptable", "", accept.ErrorNotAcceptable{[]accept.RequestedType{{1.0, "text/plain"}}, []string{"text/html"}}, "text/plain", "text/html")
	p("BadFormat", "", accept.ErrorBadAccept{"testgarbage,,,"}, "testgarbage,,,", "text/html")
	p("Quality", "text/plain", nil, "text/html;q=0.2, text/plain;q=0.8", "text/html", "text/plain")
}
