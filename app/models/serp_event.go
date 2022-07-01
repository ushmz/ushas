package models

import (
	"errors"
)

//SerpEvent : Kinds of events on SERP.
type SerpEvent string

const (
	// CLICK : The user click and view page on SERP.
	CLICK = SerpEvent("click")

	// HOVER : The user put cursor on page link on SERP.
	HOVER = SerpEvent("hover")

	// PAGINATE : The user go next/previous page on SERP.
	PAGINATE = SerpEvent("paginate")
)

// Valid : Check given serp event value is valid or not.
func (e SerpEvent) Valid() error {
	switch e {
	case CLICK:
		return nil
	case HOVER:
		return nil
	case PAGINATE:
		return nil
	default:
		return errors.New("Invalid SERP event")
	}
}
