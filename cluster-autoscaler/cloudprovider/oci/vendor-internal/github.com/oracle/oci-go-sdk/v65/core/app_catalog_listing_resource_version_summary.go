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

// AppCatalogListingResourceVersionSummary Listing Resource Version summary
type AppCatalogListingResourceVersionSummary struct {

	// The OCID of the listing this resource version belongs to.
	ListingId *string `mandatory:"false" json:"listingId"`

	// Date and time the listing resource version was published, in RFC3339 (https://tools.ietf.org/html/rfc3339) format.
	// Example: `2018-03-20T12:32:53.532Z`
	TimePublished *common.SDKTime `mandatory:"false" json:"timePublished"`

	// OCID of the listing resource.
	ListingResourceId *string `mandatory:"false" json:"listingResourceId"`

	// Resource Version.
	ListingResourceVersion *string `mandatory:"false" json:"listingResourceVersion"`
}

func (m AppCatalogListingResourceVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AppCatalogListingResourceVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
