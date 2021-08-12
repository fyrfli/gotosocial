// Code generated by astool. DO NOT EDIT.

package propertymediatype

import (
	"fmt"
	rfc2045 "github.com/go-fed/activity/streams/values/rfc2045"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsMediaTypeProperty is the functional property "mediaType". It is
// permitted to be a single default-valued value type.
type ActivityStreamsMediaTypeProperty struct {
	rfcRfc2045Member string
	hasRfc2045Member bool
	unknown          interface{}
	iri              *url.URL
	alias            string
}

// DeserializeMediaTypeProperty creates a "mediaType" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializeMediaTypeProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsMediaTypeProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/ns/activitystreams"]; ok {
		alias = a
	}
	propName := "mediaType"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "mediaType")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsMediaTypeProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if v, err := rfc2045.DeserializeRfc2045(i); err == nil {
			this := &ActivityStreamsMediaTypeProperty{
				alias:            alias,
				hasRfc2045Member: true,
				rfcRfc2045Member: v,
			}
			return this, nil
		}
		this := &ActivityStreamsMediaTypeProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsMediaTypeProperty creates a new mediaType property.
func NewActivityStreamsMediaTypeProperty() *ActivityStreamsMediaTypeProperty {
	return &ActivityStreamsMediaTypeProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling IsRFCRfc2045 afterwards
// will return false.
func (this *ActivityStreamsMediaTypeProperty) Clear() {
	this.unknown = nil
	this.iri = nil
	this.hasRfc2045Member = false
}

// Get returns the value of this property. When IsRFCRfc2045 returns false, Get
// will return any arbitrary value.
func (this ActivityStreamsMediaTypeProperty) Get() string {
	return this.rfcRfc2045Member
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ActivityStreamsMediaTypeProperty) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ActivityStreamsMediaTypeProperty) HasAny() bool {
	return this.IsRFCRfc2045() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ActivityStreamsMediaTypeProperty) IsIRI() bool {
	return this.iri != nil
}

// IsRFCRfc2045 returns true if this property is set and not an IRI.
func (this ActivityStreamsMediaTypeProperty) IsRFCRfc2045() bool {
	return this.hasRfc2045Member
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsMediaTypeProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/ns/activitystreams": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsMediaTypeProperty) KindIndex() int {
	if this.IsRFCRfc2045() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsMediaTypeProperty) LessThan(o vocab.ActivityStreamsMediaTypeProperty) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsRFCRfc2045() && !o.IsRFCRfc2045() {
		// Both are unknowns.
		return false
	} else if this.IsRFCRfc2045() && !o.IsRFCRfc2045() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsRFCRfc2045() && o.IsRFCRfc2045() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return rfc2045.LessRfc2045(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "mediaType".
func (this ActivityStreamsMediaTypeProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "mediaType"
	} else {
		return "mediaType"
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsMediaTypeProperty) Serialize() (interface{}, error) {
	if this.IsRFCRfc2045() {
		return rfc2045.SerializeRfc2045(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// Set sets the value of this property. Calling IsRFCRfc2045 afterwards will
// return true.
func (this *ActivityStreamsMediaTypeProperty) Set(v string) {
	this.Clear()
	this.rfcRfc2045Member = v
	this.hasRfc2045Member = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ActivityStreamsMediaTypeProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}