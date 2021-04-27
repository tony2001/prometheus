// Copyright 2020 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
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

// Package install has the side-effect of registering all builtin
// service discovery config types.
package install

import (
	_ "github.com/tony2001/prometheus/v2/discovery/azure"        // register azure
	_ "github.com/tony2001/prometheus/v2/discovery/consul"       // register consul
	_ "github.com/tony2001/prometheus/v2/discovery/digitalocean" // register digitalocean
	_ "github.com/tony2001/prometheus/v2/discovery/dns"          // register dns
	_ "github.com/tony2001/prometheus/v2/discovery/ec2"          // register ec2
	_ "github.com/tony2001/prometheus/v2/discovery/eureka"       // register eureka
	_ "github.com/tony2001/prometheus/v2/discovery/file"         // register file
	_ "github.com/tony2001/prometheus/v2/discovery/gce"          // register gce
	_ "github.com/tony2001/prometheus/v2/discovery/hetzner"      // register hetzner
	_ "github.com/tony2001/prometheus/v2/discovery/kubernetes"   // register kubernetes
	_ "github.com/tony2001/prometheus/v2/discovery/marathon"     // register marathon
	_ "github.com/tony2001/prometheus/v2/discovery/moby"         // register moby
	_ "github.com/tony2001/prometheus/v2/discovery/openstack"    // register openstack
	_ "github.com/tony2001/prometheus/v2/discovery/scaleway"     // register scaleway
	_ "github.com/tony2001/prometheus/v2/discovery/triton"       // register triton
	_ "github.com/tony2001/prometheus/v2/discovery/zookeeper"    // register zookeeper
)
