package model

var dataLayer interface {
	// TODO create
}

// SetDataLayer sets the local data access layer for model logic.
func SetDataLayer(l interface{}) {
	dataLayer = l
}
