/*
 *  VITacademics-ErrorPages
 *  Copyright (C) 2015  Aneesh Neelam <neelam.aneesh@gmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package base

import (
  "fmt"
  "encoding/json"
)

type status struct {
  Message string `json:"message"`
  Code int `json:"code"`
}

type obj struct {
  Status status `json:"status"`
}

func GetStatus (key string) []byte {

  maintenance := obj{
    Status: status{
      Message: "Our backend servers are Down for Maintenance, please contact the VITacademics Developers for more Information",
      Code: 98,
    },
  }
  unknown := obj{
    Status: status{
      Message: "An Unforeseen/Unknown/Irrecoverable Error has Occurred",
      Code: 99,
    },
  }

  var response []byte

  if key == "unknown" {
    obj, err := json.Marshal(unknown)
    if err != nil {
      fmt.Println(err)
  		response = nil
   	} else {
      response = obj
    }
  } else if key == "maintenance" {
    obj, err := json.Marshal(maintenance)
    if err != nil {
      fmt.Println(err)
  		response = nil
   	} else {
      response = obj
    }
  }

  return response
}
