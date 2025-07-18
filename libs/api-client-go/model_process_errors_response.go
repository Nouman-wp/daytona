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

// checks if the ProcessErrorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessErrorsResponse{}

// ProcessErrorsResponse struct for ProcessErrorsResponse
type ProcessErrorsResponse struct {
	// The name of the VNC process whose error logs were retrieved
	ProcessName string `json:"processName"`
	// The error log output from the specified VNC process
	Errors string `json:"errors"`
}

type _ProcessErrorsResponse ProcessErrorsResponse

// NewProcessErrorsResponse instantiates a new ProcessErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessErrorsResponse(processName string, errors string) *ProcessErrorsResponse {
	this := ProcessErrorsResponse{}
	this.ProcessName = processName
	this.Errors = errors
	return &this
}

// NewProcessErrorsResponseWithDefaults instantiates a new ProcessErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessErrorsResponseWithDefaults() *ProcessErrorsResponse {
	this := ProcessErrorsResponse{}
	return &this
}

// GetProcessName returns the ProcessName field value
func (o *ProcessErrorsResponse) GetProcessName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessName
}

// GetProcessNameOk returns a tuple with the ProcessName field value
// and a boolean to check if the value has been set.
func (o *ProcessErrorsResponse) GetProcessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessName, true
}

// SetProcessName sets field value
func (o *ProcessErrorsResponse) SetProcessName(v string) {
	o.ProcessName = v
}

// GetErrors returns the Errors field value
func (o *ProcessErrorsResponse) GetErrors() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ProcessErrorsResponse) GetErrorsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ProcessErrorsResponse) SetErrors(v string) {
	o.Errors = v
}

func (o ProcessErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessErrorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["processName"] = o.ProcessName
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *ProcessErrorsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"processName",
		"errors",
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

	varProcessErrorsResponse := _ProcessErrorsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProcessErrorsResponse)

	if err != nil {
		return err
	}

	*o = ProcessErrorsResponse(varProcessErrorsResponse)

	return err
}

type NullableProcessErrorsResponse struct {
	value *ProcessErrorsResponse
	isSet bool
}

func (v NullableProcessErrorsResponse) Get() *ProcessErrorsResponse {
	return v.value
}

func (v *NullableProcessErrorsResponse) Set(val *ProcessErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessErrorsResponse(val *ProcessErrorsResponse) *NullableProcessErrorsResponse {
	return &NullableProcessErrorsResponse{value: val, isSet: true}
}

func (v NullableProcessErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
