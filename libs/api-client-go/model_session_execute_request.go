/*
Daytona

Daytona AI platform API Docs

API version: 1.0
Contact: support@daytona.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SessionExecuteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionExecuteRequest{}

// SessionExecuteRequest struct for SessionExecuteRequest
type SessionExecuteRequest struct {
	// The command to execute
	Command string `json:"command"`
	// Whether to execute the command asynchronously
	RunAsync *bool `json:"runAsync,omitempty"`
	// Deprecated: Use runAsync instead. Whether to execute the command asynchronously
	// Deprecated
	Async *bool `json:"async,omitempty"`
}

type _SessionExecuteRequest SessionExecuteRequest

// NewSessionExecuteRequest instantiates a new SessionExecuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionExecuteRequest(command string) *SessionExecuteRequest {
	this := SessionExecuteRequest{}
	this.Command = command
	return &this
}

// NewSessionExecuteRequestWithDefaults instantiates a new SessionExecuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionExecuteRequestWithDefaults() *SessionExecuteRequest {
	this := SessionExecuteRequest{}
	return &this
}

// GetCommand returns the Command field value
func (o *SessionExecuteRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *SessionExecuteRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *SessionExecuteRequest) SetCommand(v string) {
	o.Command = v
}

// GetRunAsync returns the RunAsync field value if set, zero value otherwise.
func (o *SessionExecuteRequest) GetRunAsync() bool {
	if o == nil || IsNil(o.RunAsync) {
		var ret bool
		return ret
	}
	return *o.RunAsync
}

// GetRunAsyncOk returns a tuple with the RunAsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionExecuteRequest) GetRunAsyncOk() (*bool, bool) {
	if o == nil || IsNil(o.RunAsync) {
		return nil, false
	}
	return o.RunAsync, true
}

// HasRunAsync returns a boolean if a field has been set.
func (o *SessionExecuteRequest) HasRunAsync() bool {
	if o != nil && !IsNil(o.RunAsync) {
		return true
	}

	return false
}

// SetRunAsync gets a reference to the given bool and assigns it to the RunAsync field.
func (o *SessionExecuteRequest) SetRunAsync(v bool) {
	o.RunAsync = &v
}

// GetAsync returns the Async field value if set, zero value otherwise.
// Deprecated
func (o *SessionExecuteRequest) GetAsync() bool {
	if o == nil || IsNil(o.Async) {
		var ret bool
		return ret
	}
	return *o.Async
}

// GetAsyncOk returns a tuple with the Async field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SessionExecuteRequest) GetAsyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Async) {
		return nil, false
	}
	return o.Async, true
}

// HasAsync returns a boolean if a field has been set.
func (o *SessionExecuteRequest) HasAsync() bool {
	if o != nil && !IsNil(o.Async) {
		return true
	}

	return false
}

// SetAsync gets a reference to the given bool and assigns it to the Async field.
// Deprecated
func (o *SessionExecuteRequest) SetAsync(v bool) {
	o.Async = &v
}

func (o SessionExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionExecuteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["command"] = o.Command
	if !IsNil(o.RunAsync) {
		toSerialize["runAsync"] = o.RunAsync
	}
	if !IsNil(o.Async) {
		toSerialize["async"] = o.Async
	}
	return toSerialize, nil
}

func (o *SessionExecuteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"command",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSessionExecuteRequest := _SessionExecuteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSessionExecuteRequest)

	if err != nil {
		return err
	}

	*o = SessionExecuteRequest(varSessionExecuteRequest)

	return err
}

type NullableSessionExecuteRequest struct {
	value *SessionExecuteRequest
	isSet bool
}

func (v NullableSessionExecuteRequest) Get() *SessionExecuteRequest {
	return v.value
}

func (v *NullableSessionExecuteRequest) Set(val *SessionExecuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionExecuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionExecuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionExecuteRequest(val *SessionExecuteRequest) *NullableSessionExecuteRequest {
	return &NullableSessionExecuteRequest{value: val, isSet: true}
}

func (v NullableSessionExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionExecuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
