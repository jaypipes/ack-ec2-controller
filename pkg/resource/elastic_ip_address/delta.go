// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package elastic_ip_address

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.Address, b.ko.Spec.Address) {
		delta.Add("Spec.Address", a.ko.Spec.Address, b.ko.Spec.Address)
	} else if a.ko.Spec.Address != nil && b.ko.Spec.Address != nil {
		if *a.ko.Spec.Address != *b.ko.Spec.Address {
			delta.Add("Spec.Address", a.ko.Spec.Address, b.ko.Spec.Address)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CustomerOwnedIPv4Pool, b.ko.Spec.CustomerOwnedIPv4Pool) {
		delta.Add("Spec.CustomerOwnedIPv4Pool", a.ko.Spec.CustomerOwnedIPv4Pool, b.ko.Spec.CustomerOwnedIPv4Pool)
	} else if a.ko.Spec.CustomerOwnedIPv4Pool != nil && b.ko.Spec.CustomerOwnedIPv4Pool != nil {
		if *a.ko.Spec.CustomerOwnedIPv4Pool != *b.ko.Spec.CustomerOwnedIPv4Pool {
			delta.Add("Spec.CustomerOwnedIPv4Pool", a.ko.Spec.CustomerOwnedIPv4Pool, b.ko.Spec.CustomerOwnedIPv4Pool)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.NetworkBorderGroup, b.ko.Spec.NetworkBorderGroup) {
		delta.Add("Spec.NetworkBorderGroup", a.ko.Spec.NetworkBorderGroup, b.ko.Spec.NetworkBorderGroup)
	} else if a.ko.Spec.NetworkBorderGroup != nil && b.ko.Spec.NetworkBorderGroup != nil {
		if *a.ko.Spec.NetworkBorderGroup != *b.ko.Spec.NetworkBorderGroup {
			delta.Add("Spec.NetworkBorderGroup", a.ko.Spec.NetworkBorderGroup, b.ko.Spec.NetworkBorderGroup)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PublicIPv4Pool, b.ko.Spec.PublicIPv4Pool) {
		delta.Add("Spec.PublicIPv4Pool", a.ko.Spec.PublicIPv4Pool, b.ko.Spec.PublicIPv4Pool)
	} else if a.ko.Spec.PublicIPv4Pool != nil && b.ko.Spec.PublicIPv4Pool != nil {
		if *a.ko.Spec.PublicIPv4Pool != *b.ko.Spec.PublicIPv4Pool {
			delta.Add("Spec.PublicIPv4Pool", a.ko.Spec.PublicIPv4Pool, b.ko.Spec.PublicIPv4Pool)
		}
	}

	return delta
}
