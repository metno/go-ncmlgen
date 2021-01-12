/*
  Copyright 2020 MET Norway

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

package ncml

import (
	"testing"
)

var correctNcmlData = `<?xml version="1.0" encoding="UTF-8"?>
<netcdf xmlns="http://www.unidata.ucar.edu/namespaces/netcdf/ncml-2.2">
  <aggregation dimName="time" type="joinExisting">
    <netcdf location="/data/test1.nc"></netcdf>
    <netcdf location="/data/test2.nc"></netcdf>
  </aggregation>
</netcdf>`

func TestCreateNcML(t *testing.T) {
	testdata := []string{"/data/test1.nc", "/data/test2.nc"}
	output, err := createNcMLWithAggregation(testdata)
	if err != nil {
		t.Errorf("Creation failed: %v", err)
	}

	outputLength := len(*output)
	correctNcmlDataLength := len(correctNcmlData)

	if outputLength != correctNcmlDataLength {
		t.Errorf("Content length mismatch.\n Expected: %d\n Got: %d", correctNcmlDataLength, outputLength)
	}

	if *output != correctNcmlData {
		t.Errorf("Expected %v; Got %v", correctNcmlData, *output)
	}
}

func TestMarshalFail(t *testing.T) {
	testdata := []string{""}
	createNcMLWithAggregation(testdata)
}
