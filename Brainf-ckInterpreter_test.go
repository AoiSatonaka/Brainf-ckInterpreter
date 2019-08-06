package Brainf_ckInterpreter

import "testing"

func TestInterpretCacheSuccess(t *testing.T){
	err := Interpret([]rune("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.+.+."),nil)
	if err != nil {
		t.Fatalf("failed test(cache) %#v",err)
	}
}

func TestInterpretLoopSuccess(t *testing.T){
	err := Interpret([]rune("++++++++[>++++++++<-]>+.+.+."),nil)
	if err != nil {
		t.Fatalf("failed test(loop) %#v",err)
	}
}

func TestInterpretInputSuccess(t *testing.T) {
	println("inputCheck")
	err := Interpret([]rune(",.,,,."),[]rune("abcde"))
	if err != nil {
		t.Fatalf("failed test %#v",err)
	}
}

//func TestInterpretFailed(t *testing.T) {
//	err := Interpret([]rune("+++++++++++\n\r"))
//	if err != nil {
//		t.Fatalf("failed test %#v",err)
//	}
//}
