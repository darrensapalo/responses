package responses

import (
	"strconv"
)

type (
	// JSONError is a struct representation of JSON error as defined in https://jsonapi.org/format/#error-objects
	JSONError struct {
		// ID is a unique identifier for this particular occurrence of the problem
		ID string `json:"id,omitempty"`
		// Links is a string or a Link object that points to more details about the error
		Links map[string]interface{} `json:"links,omitempty"`
		// Status the HTTP status code expressed as a string
		Status string `json:"status,omitempty"`
		// Code is an application specific error code expressed as a string
		Code string `json:"code,omitempty"`
		// Title is a short, human-readable summary of the problem. SHOULD NOT change from occurrence to occurrence of the
		// same problem
		Title string `json:"title,omitempty"`
		// Detail is a human-readable explaination specific to this occurrence of the problem
		Detail string `json:"detail,omitempty"`
		// Source is an object containing references to the source of the error
		Source Source `json:"source,omitempty"`
		// Meta contains non standard meta-information about the error
		Meta map[string]interface{} `json:"meta,omitempty"`
	}

	// Link represents a link object
	Link struct {
		// HREF specifies the URL
		HREF string `json:"href"`
		// Meta contains non standard meta-information about the link
		Meta map[string]interface{} `json:"meta,omitempty"`
	}

	// Source is an object containing reference to the source of the error
	Source struct {
		Pointer   string      `json:"pointer,omitempty"`
		Parameter string      `json:"parameter,omitempty"`
		Other     interface{} `json:",inline,omitempty"`
	}

	// JSONErrorBuilder is a convenience builder for creating a new JSONError
	JSONErrorBuilder struct {
		*JSONError
	}
)

// ErrorBuilder is a convenience function for building new JSONError
func ErrorBuilder() *JSONErrorBuilder {
	return &JSONErrorBuilder{&JSONError{}}
}

// ID sets the id of the JSONError
func (b *JSONErrorBuilder) ID(id string) *JSONErrorBuilder {
	b.JSONError.ID = id
	return b
}

// AddSimpleLink adds a simple link with a string url
func (b *JSONErrorBuilder) AddSimpleLink(key, url string) *JSONErrorBuilder {
	if b.JSONError.Links == nil {
		b.JSONError.Links = make(map[string]interface{})
	}
	b.JSONError.Links[key] = url
	return b
}

// AddLink Adds an Link
func (b *JSONErrorBuilder) AddLink(key, url string, meta map[string]interface{}) *JSONErrorBuilder {
	if b.JSONError.Links == nil {
		b.JSONError.Links = make(map[string]interface{})
	}
	b.JSONError.Links[key] = Link{
		HREF: url,
		Meta: meta,
	}
	return b
}

// Status sets the HTTP status of the error
func (b *JSONErrorBuilder) Status(status int) *JSONErrorBuilder {
	b.JSONError.Status = strconv.Itoa(status)
	return b
}

// Code sets the application specific code
func (b *JSONErrorBuilder) Code(code string) *JSONErrorBuilder {
	b.JSONError.Code = code
	return b
}

// Title sets the summary of the error
func (b *JSONErrorBuilder) Title(title string) *JSONErrorBuilder {
	b.JSONError.Title = title
	return b
}

// Detail sets the detailed error
func (b *JSONErrorBuilder) Detail(detail string) *JSONErrorBuilder {
	b.JSONError.Detail = detail
	return b
}

// SourcePointer sets the pointer of the request body to where the error occured
func (b *JSONErrorBuilder) SourcePointer(pointer string) *JSONErrorBuilder {
	b.JSONError.Source = Source{
		Pointer: pointer,
	}
	return b
}

// SourceParameter sets the name of the header/path/query parameter that caused the error
func (b *JSONErrorBuilder) SourceParameter(parameter string) *JSONErrorBuilder {
	b.JSONError.Source = Source{
		Parameter: parameter,
	}
	return b
}

// Source sets the source of the error
func (b *JSONErrorBuilder) Source(pointer, parameter string, other interface{}) *JSONErrorBuilder {
	b.JSONError.Source = Source{
		Pointer:   pointer,
		Parameter: parameter,
		Other:     other,
	}
	return b
}

// AddMeta adds metadata to the error
func (b *JSONErrorBuilder) AddMeta(key string, value interface{}) *JSONErrorBuilder {
	if b.JSONError.Meta == nil {
		b.JSONError.Meta = map[string]interface{}{}
	}
	b.JSONError.Meta[key] = value
	return b
}

// Build creates the JSONError
func (b *JSONErrorBuilder) Build() JSONError {
	return *b.JSONError
}
