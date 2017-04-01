package main

import (
	"fmt"
	"github.com/mcuadros/go-version"
)

func MyTest(in []string, out []string) bool {
	if len(in) != len(out) {
		return false
	}
	for i, _ := range in {
		if in[i] != out[i] {
			return false
		}
	}
	return true
}
func main() {
	//in := []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
	//out := []string{"1.0.2", "1.0.3", "1.1", "1.1.1", "1.1.2", "1.2rc1", "1.2rc2", "1.2rc3", "1.2rc4", "1.2rc5", "1.2", "1.2.1", "1.2.2", "1.3beta1", "1.3beta2", "1.3rc1", "1.3rc2", "1.3", "1.3.1", "1.3.2", "1.3.3", "1.4beta1", "1.4rc1", "1.4rc2", "1.4", "1.4.1", "1.4.2", "1.5alpha1", "1.5beta1", "1.5beta2", "1.5rc1", "1.5", "1.5.1", "1.5.2", "1.6beta1"}
	in := []string{"11.1.2", "12.2.3", "9.5.4"}
	version.Sort(in)
	fmt.Println(in)
	//fmt.Println(MyTest(in, out))
}
