// Auto-generated by avdl-compiler v1.4.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/debugging.avdl

package keybase1

type FirstStepResult struct {
	ValPlusTwo int `codec:"valPlusTwo" json:"valPlusTwo"`
}

func (o FirstStepResult) DeepCopy() FirstStepResult {
	return FirstStepResult{
		ValPlusTwo: o.ValPlusTwo,
	}
}
