// cvtLib.go
// author pr azul software
// date: 9. June 2022
// update:
// copyright prr, azul softwre
// license: see github
//

package cvtLib

import (
    "fmt"
    "os"
    )

const (
	json = iota
	toml
	yaml
)

type cvtObj struct {
	inpFil *os.File
	goFil *os.File
	typ string
}

func CreGoFil(inpfilNam, goFilNam string)(err error) {

	var cvt cvtObj

	err = cvt.ParseInpFil(inpfilNam)
	if err != nil {return fmt.Errorf("error ParseInpFil: %v", err)}

	err = cvt.ParseGoFil(goFilNam)
	if err != nil {return fmt.Errorf("error ParseGoFil: %v", err)}

	switch cvt.typ {
	case "json":
		err = cvt.CvtJson()
	case "toml":
		err = cvt.CvtToml()
	case "yaml":
		err = cvt.CvtYaml()
	default:
		return fmt.Errorf("cvt.typ not valid!")
	}
	if err != nil {return fmt.Errorf("error Cvt: %v", err)}

	return nil
}

func (cvt *cvtObj) ParseGoFil(gofilNam string)(err error) {

	extSt := -1
	for i:=0; i<len(gofilNam); i++ {
		if gofilNam[i] == '.' {
			extSt = i+1
		}
	}
	if extSt < 0 {
		gofilNam += ".go"
	} else {
		extStr := string(gofilNam[extSt+1:])
		if extStr != "go" {return fmt.Errorf("invalid gofilNam extension!")}
	}

	gofil, err := os.Create(gofilNam)
	if err != nil {return fmt.Errorf("os.Create: %v", err)}
	cvt.goFil = gofil
	return nil
}

func (cvt *cvtObj) ParseInpFil(infilNam string)(err error) {

	extSt := -1
	for i:=0; i<len(infilNam); i++ {
		if infilNam[i] == '.' {
			extSt = i+1
		}
	}
	if extSt < 0 {return fmt.Errorf("no extension!")}


	extStr := string(infilNam[extSt:])

	fmt.Println("extension: ", extStr)

	switch extStr {

	case "json":
		cvt.typ = "json"
	case "yaml":
		cvt.typ = "yaml"
	case "toml":
		cvt.typ = "toml"
	default:
		return fmt.Errorf("invalid extension: %s", extStr)
	}

	inpfil, err := os.Open(infilNam)
	if err != nil { return fmt.Errorf("os.Open: %v", err) }
	cvt.inpFil = inpfil


	return nil
}

func (cvt *cvtObj) CvtJson()(err error) {

	if cvt.inpFil == nil {return fmt.Errorf("no input file!")}
	if cvt.goFil == nil {return fmt.Errorf("no go file!")}


	return nil
}

func (cvt *cvtObj) CvtToml()(err error) {

	if cvt.inpFil == nil {return fmt.Errorf("no input file!")}
	if cvt.goFil == nil {return fmt.Errorf("no go file!")}


	return nil
}

func (cvt *cvtObj) CvtYaml()(err error) {

	if cvt.inpFil == nil {return fmt.Errorf("no input file!")}
	if cvt.goFil == nil {return fmt.Errorf("no go file!")}


	return nil
}
