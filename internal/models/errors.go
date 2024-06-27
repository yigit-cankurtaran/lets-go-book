package models

import "errors"

// ErrNoRecord is an error message for no record found
var ErrNoRecord = errors.New("models: no matching record found")

// we use this to encapsulate the model completely
// so our app doesn't need to know about the database
// or any of its specific error handling
