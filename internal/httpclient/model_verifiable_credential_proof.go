/*
Ory Hydra API

Documentation for all of Ory Hydra's APIs.

API version:
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VerifiableCredentialProof struct for VerifiableCredentialProof
type VerifiableCredentialProof struct {
	Jwt       *string `json:"jwt,omitempty"`
	ProofType *string `json:"proof_type,omitempty"`
}

// NewVerifiableCredentialProof instantiates a new VerifiableCredentialProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiableCredentialProof() *VerifiableCredentialProof {
	this := VerifiableCredentialProof{}
	return &this
}

// NewVerifiableCredentialProofWithDefaults instantiates a new VerifiableCredentialProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiableCredentialProofWithDefaults() *VerifiableCredentialProof {
	this := VerifiableCredentialProof{}
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *VerifiableCredentialProof) GetJwt() string {
	if o == nil || o.Jwt == nil {
		var ret string
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialProof) GetJwtOk() (*string, bool) {
	if o == nil || o.Jwt == nil {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *VerifiableCredentialProof) HasJwt() bool {
	if o != nil && o.Jwt != nil {
		return true
	}

	return false
}

// SetJwt gets a reference to the given string and assigns it to the Jwt field.
func (o *VerifiableCredentialProof) SetJwt(v string) {
	o.Jwt = &v
}

// GetProofType returns the ProofType field value if set, zero value otherwise.
func (o *VerifiableCredentialProof) GetProofType() string {
	if o == nil || o.ProofType == nil {
		var ret string
		return ret
	}
	return *o.ProofType
}

// GetProofTypeOk returns a tuple with the ProofType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiableCredentialProof) GetProofTypeOk() (*string, bool) {
	if o == nil || o.ProofType == nil {
		return nil, false
	}
	return o.ProofType, true
}

// HasProofType returns a boolean if a field has been set.
func (o *VerifiableCredentialProof) HasProofType() bool {
	if o != nil && o.ProofType != nil {
		return true
	}

	return false
}

// SetProofType gets a reference to the given string and assigns it to the ProofType field.
func (o *VerifiableCredentialProof) SetProofType(v string) {
	o.ProofType = &v
}

func (o VerifiableCredentialProof) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jwt != nil {
		toSerialize["jwt"] = o.Jwt
	}
	if o.ProofType != nil {
		toSerialize["proof_type"] = o.ProofType
	}
	return json.Marshal(toSerialize)
}

type NullableVerifiableCredentialProof struct {
	value *VerifiableCredentialProof
	isSet bool
}

func (v NullableVerifiableCredentialProof) Get() *VerifiableCredentialProof {
	return v.value
}

func (v *NullableVerifiableCredentialProof) Set(val *VerifiableCredentialProof) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiableCredentialProof) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiableCredentialProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiableCredentialProof(val *VerifiableCredentialProof) *NullableVerifiableCredentialProof {
	return &NullableVerifiableCredentialProof{value: val, isSet: true}
}

func (v NullableVerifiableCredentialProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiableCredentialProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
