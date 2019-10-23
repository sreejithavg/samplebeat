// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package panw

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "panw", asset.ModuleFieldsPri, AssetPanw); err != nil {
		panic(err)
	}
}

// AssetPanw returns asset data.
// This is the base64 encoded gzipped contents of module/panw.
func AssetPanw() string {
	return "eJzMmM9u4zYQxu95irm5BeLcesmhQNpFgABpauw26HFBUyOJNcXRDkdxtU+/IC3bWpu2tLbXiG6STH6/+fNRpKewwPYeauWWNwBixOLmLkOv2dRiyN3D7zcAAH9R1liEnBhmyhI8WCF4QVkSLzz8Mnt4mf796dcbgNygzfx9HDQFp6rttOGStsZ7KJiaunuSEAvXY5wHcqYKpMQ4B1SR4q77UV8KdvTIbx+nVI9Jf6dPHOUTQa9iBkuFv+uP3cPqkXFj0aN8L9XhLbBdEmc7744xAsCLqhAoj4xhcpBSCVRKdIkZSGk8ePTekLtLAnlqWGOSZy9dwzRd0oQA/xd0WcQSqqcW39B2YkDz/1DL3c7oVNr6pF/J7XIO5W4Ecbg+rbCCQFfvQ1nr8xgnyLnaS95loTYqP0Dm1G6DwdGijiSakZfpy8M/6zKqLGP0/hZMvn4U3hoPNXJOXGG2z3i4zr3MpgDXARx4OYJ/P4KnWQpws4oQp/K4BrHkisuhBLEtTNKqGXoxToV5r+TXnuK7M+2HHtv7cm6f7D3at1/Vvof7z88z8qCVB8w80kMHYkq7eoSvB519HtcIi7vV3uJK9u7UTrR2rVX92aQ8cgEPzZReoIBWtTSM8PQh+keBlIzqUBbhTBONaWhNVdU4I2069DHhj0xBuP5cq8UMWFpOS+XLzaY0dNhvU2nq7Z74QGPlxl5rkxekTmypENxp/fKHcYpb2GRn3SkrGo9OAu8cQTll26+Yrsu8jbH8a2z2aDiM4zejMbWcpIuczH3D9kqpb9iemHmtBAvi9ue4+TH2a6zH68fn8LWRiY/srx+fN9rpVTuMXRfkNo55Q86MFiAXb2OFlcvA+OQEaKREhkmlrNGGGj+5hUnBql0qxsktEMNkjs4UbjJkIkvLfdufcXh7CrsDpyy4pkI2GkyGTkxukGMXo9Ll/n4hfY7DLw06jZ9dU82Rk4yJ79oA4DMVgE647ZPFE6bxYJxmrNAJZp28GGVtoo6vznxpcBuSpSIiDcTULfaMRw6pJ+U9dB3xqnOC1JiPyw7UJdsg8d/CTiNEm/8AXbj5qXybrPXJvgUAAP//KvmK+g=="
}
