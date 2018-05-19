package rtorrent

import (
	"encoding/xml"
	"errors"
)

func NewMethodCall(method string, id *string) MethodCall{
	mc := MethodCall{
		MethodName:method,
	}

	if id != nil{
		mc.Cparams=&Cparams{
			Cparam:&Cparam{
				Cvalue:[]*Cvalue{
					{ Cstring: &Cstring {String: *id, }},
				},
			},
		}
	}
	return mc
}

///////////////////////////
/// structs - to unmarshal from the xmlrpc json
///////////////////////////

type MethodCall struct {
	XMLName xml.Name `xml:"methodCall,omitempty" json:"methodCall,omitempty"`
	MethodName string `xml:"methodName"`
	Cparams *Cparams `xml:"params,omitempty" json:"params,omitempty"`
}

type Carray struct {
	XMLName xml.Name `xml:"array,omitempty" json:"array,omitempty"`
	Cdata *Cdata `xml:"data,omitempty" json:"data,omitempty"`
}

type Cdata struct {
	XMLName xml.Name `xml:"data,omitempty" json:"data,omitempty"`
	Cvalue []*Cvalue `xml:"value,omitempty" json:"value,omitempty"`
}

type CmethodResponse struct {
	XMLName xml.Name `xml:"methodResponse,omitempty" json:"methodResponse,omitempty"`
	Cparams *Cparams `xml:"params,omitempty" json:"params,omitempty"`
	Cfault *Cfault `xml:"fault,omitempty" json:"fault,omitempty"`
}

type Cparam struct {
	XMLName xml.Name `xml:"param,omitempty" json:"param,omitempty"`
	Cvalue []*Cvalue `xml:"value,omitempty" json:"value,omitempty"`
}

type Cparams struct {
	XMLName xml.Name `xml:"params,omitempty" json:"params,omitempty"`
	Cparam *Cparam `xml:"param,omitempty" json:"param,omitempty"`
}

type Cstring struct {
	XMLName xml.Name `xml:"string,omitempty" json:"string,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type Ci4 struct{
	XMLName xml.Name `xml:"i4,omitempty" json:"i4,omitempty"`
	Int int `xml:",chardata" json:",omitempty"`
}

type Ci8 struct{
	XMLName xml.Name `xml:"i8,omitempty" json:"i8,omitempty"`
	Int int `xml:",chardata" json:",omitempty"`
}

type Cvalue struct {
	XMLName xml.Name `xml:"value,omitempty" json:"value,omitempty"`
	Carray *Carray `xml:"array,omitempty" json:"array,omitempty"`
	Cstring *Cstring `xml:"string,omitempty" json:"string,omitempty"`
	Ci4 *Ci4 `xml:"i4,omitempty" json:"i4,omitempty"`
	Ci8 *Ci8 `xml:"i8,omitempty" json:"i8,omitempty"`
	Cstruct *Cstruct `xml:"struct,omitempty" json:"struct,omitempty"`
}

///---------////

type Cfault struct {
	XMLName xml.Name `xml:"fault,omitempty" json:"fault,omitempty"`
	Cvalue *Cvalue `xml:"value,omitempty" json:"value,omitempty"`
}

type Cstruct struct {
	XMLName xml.Name `xml:"struct,omitempty" json:"struct,omitempty"`
	Cmember []*Cmember `xml:"member,omitempty" json:"member,omitempty"`
}

type Cmember struct {
	XMLName xml.Name `xml:"member,omitempty" json:"member,omitempty"`
	Cname *Cname `xml:"name,omitempty" json:"name,omitempty"`
	Cvalue *Cvalue `xml:"value,omitempty" json:"value,omitempty"`
}

type Cname struct {
	XMLName xml.Name `xml:"name,omitempty" json:"name,omitempty"`
	string string `xml:",chardata" json:",omitempty"`
}

func (r *CmethodResponse) ErrIfFault() error{
	if r.Cfault != nil {
		return errors.New("fault")
	}
	return nil
}

func (r *CmethodResponse) GetArray() ([]*Cvalue, error) {

	err := r.ErrIfFault()
	if err != nil{
		return nil, err
	}

	if len(r.Cparams.Cparam.Cvalue) == 0 ||
		r.Cparams.Cparam.Cvalue[0].Carray == nil ||
		r.Cparams.Cparam.Cvalue[0].Carray.Cdata == nil {
		return nil, errors.New("no data")
	}

	return r.Cparams.Cparam.Cvalue[0].Carray.Cdata.Cvalue, nil
}

func (r *CmethodResponse) GetStrings() ([]string, error){
	data, err := r.GetArray()
	if err != nil{
		return nil, err
	}

	result := make([]string, 0)
	for _, x := range data{
		if x.Cstring != nil{
			result = append(result, x.Cstring.String)
		}
	}

	return result, nil
}

func (r *CmethodResponse) GetString() (string, error){
	err := r.ErrIfFault()
	if err != nil {
		return "", err
	}

	if len(r.Cparams.Cparam.Cvalue) == 0 {
		return "", errors.New("no data")
	}

	if r.Cparams.Cparam.Cvalue[0].Cstring != nil{
		return r.Cparams.Cparam.Cvalue[0].Cstring.String, nil
	}

	return "", errors.New("no data")
}

func (r *CmethodResponse) GetInt() (int, error){

	err := r.ErrIfFault()
	if err != nil{
		return 0, err
	}

	if len(r.Cparams.Cparam.Cvalue) == 0 {
		return 0, errors.New("no data")
	}

	if r.Cparams.Cparam.Cvalue[0].Ci8 != nil{
		return r.Cparams.Cparam.Cvalue[0].Ci8.Int, nil
	}

	if r.Cparams.Cparam.Cvalue[0].Ci4 != nil{
		return r.Cparams.Cparam.Cvalue[0].Ci4.Int, nil
	}

	return 0, errors.New("no data")
}

func (r *CmethodResponse) GetBool() (bool, error) {
	i, err := r.GetInt()
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, errors.New("no data")
	}
}

///////////////////////////


