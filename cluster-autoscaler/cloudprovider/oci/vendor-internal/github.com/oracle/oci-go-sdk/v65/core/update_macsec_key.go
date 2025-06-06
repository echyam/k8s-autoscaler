// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMacsecKey An object defining the OCID of the Secret held in Vault that represent the MACsec key.
type UpdateMacsecKey struct {

	// Secret OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity Association Key Name (CKN) of this MACsec key.
	ConnectivityAssociationNameSecretId *string `mandatory:"true" json:"connectivityAssociationNameSecretId"`

	// The secret version of the connectivity association name secret in Vault.
	ConnectivityAssociationNameSecretVersion *int64 `mandatory:"true" json:"connectivityAssociationNameSecretVersion"`

	// Secret OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) containing the Connectivity Association Key (CAK) of this MACsec key.
	ConnectivityAssociationKeySecretId *string `mandatory:"true" json:"connectivityAssociationKeySecretId"`

	// The secret version of the connectivityAssociationKey secret in Vault.
	ConnectivityAssociationKeySecretVersion *int64 `mandatory:"true" json:"connectivityAssociationKeySecretVersion"`
}

func (m UpdateMacsecKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMacsecKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
