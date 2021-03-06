// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build analytics

package analytics

import (
	"fmt"
	"strings"

	"github.com/google/gapid/core/os/device"
)

func useragent(d *device.Configuration, v AppVersion) string {
	product := fmt.Sprintf("%v/%v.%v.%v", v.Name, v.Major, v.Minor, v.Point)
	info := []string{}
	os := d.GetOS()
	switch os.GetKind() {
	case device.Windows:
		info = append(info, fmt.Sprintf("Windows NT %v.%v", os.Major, os.Minor))
		if d.GetHardware().GetCPU().GetArchitecture().Bitness() == 64 {
			info = append(info, "x64")
		}
	case device.OSX:
		info = append(info, "Macintosh", fmt.Sprintf("Intel Mac OS X %v_%v_%v", os.Major, os.Minor, os.Point))

	case device.Linux:
		info = append(info, "Linux")

	case device.Android:
		info = append(info, "Linux", "U", fmt.Sprintf("Android %v.%v.%v", os.Major, os.Minor, os.Point))
		if os.Build != "" {
			info = append(info, "Build/"+os.Build)
		}
	}
	if len(info) > 0 {
		return fmt.Sprintf("%v (%v)", product, strings.Join(info, "; "))
	}
	return product
}
