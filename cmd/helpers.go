// Copyright © 2018 Dennis273 <dennic695@gmail.com>
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

package cmd

import (
	"fmt"
)

func ValidateUsername(username string) bool {
	if username == "" {
		fmt.Println("Username cannot be empty.")
		return false
	}
	// check if username is taken
	return true
}
func ValidatePassword(password string) bool {
	if password == "" {
		fmt.Println("Password cannot be empty.")
		return false
	}
	return true
}

func ValidateEmail(email string) bool {
	if email == "" {
		fmt.Println("Email cannot be empty.")
		return false
	}
	return true
}

func ValidatePhoneNumber(phone string) bool {
	if phone == "" {
		fmt.Println("Phone number cannot be empty.")
		return false
	}
	return true
}
