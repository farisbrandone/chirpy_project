package main

import(
	"slices"
	"strings"
	"fmt"
)


func joinTolower(result Bol) string{
	tabResult:= strings.Split(result.Body, " ")
	for i, word:=range tabResult{
		lowerWord:=strings.ToLower(word)
		filterDate:=[]string{"kerfuffle","sharbert","fornax"}
		if slices.Contains(filterDate, lowerWord){
			tabResult[i]="****"
		}
		
	}
	    fmt.Println(tabResult)
		finalResults:=strings.Join(tabResult," ")
		return finalResults
}