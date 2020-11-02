/*
Copyright © 2020 iiusky sky@03sec.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package slog

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {

	tests := []struct {
		name  string
		lable string
		err   error
	}{
		{
			name:  "Test Error Log",
			lable: "error",
			err:   errors.New("This is Error msg"),
		},
		{
			name:  "Test Info Log",
			lable: "info",
			err:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.lable == "error" {
				Logger(&SLoggerSetting{
					AppName:    "TestApplication",
					Path:       "./",
					IsDebug:    false,
					CallerSkip: 0,
				}).Error(tt.name, ErrorLog(tt.err))
			} else {
				Logger().Info(tt.name)
			}
		})
	}
}
