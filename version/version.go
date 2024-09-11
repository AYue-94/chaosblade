/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

import "os"

var (
	Ver         = "1.7.4"
	Env         = "unknown"
	BuildTime   = "unknown"
	UserHome, _ = os.UserHomeDir()
	Home        = UserHome + "/github/chaosblade/target/chaosblade-1.7.4"
	YamlHome    = Home + "/yaml"
	SandboxHome = Home + "/lib/sandbox"
)
