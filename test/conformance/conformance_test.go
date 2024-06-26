// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

//go:build conformance
// +build conformance

package conformance

import (
	"flag"
	"testing"

	"sigs.k8s.io/gateway-api/conformance"
	"sigs.k8s.io/gateway-api/conformance/tests"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
	"sigs.k8s.io/gateway-api/pkg/features"
)

func TestGatewayAPIConformance(t *testing.T) {
	flag.Parse()

	opts := conformance.DefaultOptions(t)
	opts.SkipTests = []string{
		tests.GatewayStaticAddresses.ShortName,
		tests.GatewayHTTPListenerIsolation.ShortName,          // https://github.com/kubernetes-sigs/gateway-api/issues/3049
		tests.HTTPRouteBackendRequestHeaderModifier.ShortName, // https://github.com/envoyproxy/gateway/issues/3338
	}
	opts.SupportedFeatures = features.AllFeatures
	opts.ExemptFeatures = features.MeshCoreFeatures

	cSuite, err := suite.NewConformanceTestSuite(opts)
	if err != nil {
		t.Fatalf("Error creating conformance test suite: %v", err)
	}
	cSuite.Setup(t, tests.ConformanceTests)
	if err := cSuite.Run(t, tests.ConformanceTests); err != nil {
		t.Fatalf("Error running conformance tests: %v", err)
	}
}
