// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstallAddonDetails The properties that define to install/enable addon on a cluster
type InstallAddonDetails struct {

	// The name of the addon.
	AddonName *string `mandatory:"true" json:"addonName"`

	// The version of addon to be installed.
	Version *string `mandatory:"false" json:"version"`

	// Addon configuration details.
	Configurations []AddonConfiguration `mandatory:"false" json:"configurations"`

	// Whether or not to override an existing addon installation. Defaults to false. If set to true, any existing addon installation would be overridden as per new installation details.
	IsOverrideExisting *bool `mandatory:"false" json:"isOverrideExisting"`
}

func (m InstallAddonDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstallAddonDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
