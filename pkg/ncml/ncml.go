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
	"encoding/xml"
	"fmt"
)

// Aggregation is a NcML aggretation element. See https://www.unidata.ucar.edu/software/netcdf-java/v4.6/ncml/AnnotatedSchema4.html for details.
type Aggregation struct {
	XMLName    xml.Name `xml:"aggregation,omitempty"`
	DimName    string   `xml:"dimName,attr,omitempty"`
	Type       string   `xml:"type,attr,omitempty"`
	NetCDFList []NetCDF
}

// NetCDF is the root tag of the NcML instance document, and is said to define a NetCDF dataset. See https://www.unidata.ucar.edu/software/netcdf-java/v4.6/ncml/AnnotatedSchema4.html for details.
type NetCDF struct {
	XMLName     xml.Name `xml:"netcdf"`
	XMLNS       string   `xml:"xmlns,attr,omitempty"`
	Aggregation *Aggregation
	Location    string `xml:"location,attr,omitempty"`
}

// CreateNcMLWithAggregation returns a ncml file body as a string aggregating provided paths into one ncml file
func CreateNcMLWithAggregation(locations []string) (*string, error) {

	v := &NetCDF{XMLNS: "http://www.unidata.ucar.edu/namespaces/netcdf/ncml-2.2", Aggregation: &Aggregation{DimName: "time", Type: "joinExisting", NetCDFList: []NetCDF{}}}

	for _, s := range locations {
		v.Aggregation.NetCDFList = append(v.Aggregation.NetCDFList, NetCDF{Aggregation: nil, Location: s})
	}

	myString, err := xml.MarshalIndent(&v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling: %v", err)
	}
	myString = []byte(xml.Header + string(myString))

	s := fmt.Sprintf("%s\n", myString)

	return &s, nil
}
