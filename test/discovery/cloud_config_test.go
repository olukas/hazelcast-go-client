// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discovery

import (
	"testing"

	"os"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/config"
	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/test/assert"
)

var testDiscoveryToken = "testDiscoveryToken"

func TestCloudConfigDefaults(t *testing.T) {
	cfg := hazelcast.NewConfig()
	cloudConfig := cfg.NetworkConfig().CloudConfig()
	assert.Equalf(t, nil, false, cloudConfig.IsEnabled(), "Default cloud config should be disabled.")
	assert.Equalf(t, nil, "", cloudConfig.DiscoveryToken(), "Default cloud config discovery token"+
		" should be empty.")
}

func TestCloudConfig(t *testing.T) {
	cfg := hazelcast.NewConfig()
	cloudConfig := config.NewClientCloud()
	cloudConfig.SetEnabled(true)
	cloudConfig.SetDiscoveryToken(testDiscoveryToken)
	cfg.NetworkConfig().SetCloudConfig(cloudConfig)
	returnedCloudCfg := cfg.NetworkConfig().CloudConfig()
	assert.Equalf(t, nil, true, returnedCloudCfg.IsEnabled(), "Cloud discovery should be enabled.")
	assert.Equalf(t, nil, testDiscoveryToken, returnedCloudCfg.DiscoveryToken(), "Cloud discovery token "+
		"should be set.")
}

func TestCloudConfigWithPropertySet(t *testing.T) {
	cloudConfig := config.NewClientCloud()
	cloudConfig.SetEnabled(true)
	os.Setenv("hazelcast.client.cloud.discovery.token", testDiscoveryToken)
	cfg := hazelcast.NewConfig()
	cfg.NetworkConfig().SetCloudConfig(cloudConfig)
	_, err := hazelcast.NewClientWithConfig(cfg)
	if _, ok := err.(*core.HazelcastIllegalStateError); !ok {
		t.Error("Cloud discovery should have returned an error for both property and client configuration based" +
			" setup")
	}
}