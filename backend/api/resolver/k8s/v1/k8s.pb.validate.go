// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resolver/k8s/v1/k8s.proto

package k8sv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _k_8_s_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PodID with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PodID) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Clientset

	// no validation rules for Namespace

	return nil
}

// PodIDValidationError is the validation error returned by PodID.Validate if
// the designated constraints aren't met.
type PodIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PodIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PodIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PodIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PodIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PodIDValidationError) ErrorName() string { return "PodIDValidationError" }

// Error satisfies the builtin error interface
func (e PodIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPodID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PodIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PodIDValidationError{}

// Validate checks the field values on IPAddress with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *IPAddress) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IpAddress

	return nil
}

// IPAddressValidationError is the validation error returned by
// IPAddress.Validate if the designated constraints aren't met.
type IPAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPAddressValidationError) ErrorName() string { return "IPAddressValidationError" }

// Error satisfies the builtin error interface
func (e IPAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPAddressValidationError{}

// Validate checks the field values on HPAName with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HPAName) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Clientset

	// no validation rules for Namespace

	return nil
}

// HPANameValidationError is the validation error returned by HPAName.Validate
// if the designated constraints aren't met.
type HPANameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPANameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPANameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPANameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPANameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPANameValidationError) ErrorName() string { return "HPANameValidationError" }

// Error satisfies the builtin error interface
func (e HPANameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPAName.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPANameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPANameValidationError{}

// Validate checks the field values on Deployment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Deployment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Clientset

	// no validation rules for Namespace

	return nil
}

// DeploymentValidationError is the validation error returned by
// Deployment.Validate if the designated constraints aren't met.
type DeploymentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentValidationError) ErrorName() string { return "DeploymentValidationError" }

// Error satisfies the builtin error interface
func (e DeploymentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentValidationError{}

// Validate checks the field values on StatefulSet with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StatefulSet) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Clientset

	// no validation rules for Namespace

	return nil
}

// StatefulSetValidationError is the validation error returned by
// StatefulSet.Validate if the designated constraints aren't met.
type StatefulSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatefulSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatefulSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatefulSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatefulSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatefulSetValidationError) ErrorName() string { return "StatefulSetValidationError" }

// Error satisfies the builtin error interface
func (e StatefulSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatefulSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatefulSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatefulSetValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Clientset

	// no validation rules for Namespace

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Facet with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Facet) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Servicename

	// no validation rules for Environment

	// no validation rules for Facet

	// no validation rules for Type

	return nil
}

// FacetValidationError is the validation error returned by Facet.Validate if
// the designated constraints aren't met.
type FacetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FacetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FacetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FacetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FacetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FacetValidationError) ErrorName() string { return "FacetValidationError" }

// Error satisfies the builtin error interface
func (e FacetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFacet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FacetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FacetValidationError{}
