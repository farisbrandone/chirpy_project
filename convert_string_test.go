package main

import (
	"testing"
)

func Mytest(t *testing.T){
cases:=[]struct{
	input Bol
	expected string
}{
	{
		input: Bol{
			Body: "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
			Extras: "this should be ignored",
		  },
		  expected:"I hear Mastodon is better than Chirpy. **** I need to migrate",
	},
	{
		input: Bol{
			Body: "I really need a kerfuffle to go to bed sooner, Fornax !",
		  },
		expected:"I really need a **** to go to bed sooner, **** !" ,
	},
}

for _,cse:=range cases{
	actual:=joinTolower(cse.input)
	if actual!=cse.expected {
		t.Errorf("the length are not equal: %s vs %s", 
		actual, 
		cse.expected,
	)
	
		continue
	}
	
}

}