// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workload

import (
	"time"

	"github.com/banzaicloud/allspark/internal/platform/log"
	"github.com/spf13/viper"
)

const EchoWorkloadName = "Echo"

type EchoWorkload struct {
	name string
	str  string

	logger log.Logger
}

func NewEchoWorkload(str string, logger log.Logger) Workload {
	return &EchoWorkload{
		name: EchoWorkloadName,
		str:  str,

		logger: logger,
	}
}

func (w *EchoWorkload) GetName() string {
	return w.name
}

func (w *EchoWorkload) Execute() (string, string, error) {
	delay := viper.GetDuration("ECHO_DELAY")
	if delay != 0 {
		w.logger.Info("DELAY", delay)
		time.Sleep(delay)
	}
	return viper.GetString("ECHO_STR"), "text/plain", nil
	//return w.str, "text/plain", nil
}
