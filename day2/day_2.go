package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type policy struct{
	minimum uint
	maximum uint
	letter string
	data []string
	valid bool
}

func (p *policy)validate(){
	var occurance uint
	for _, aCharacter := range p.data{
		if aCharacter == p.letter{
			occurance++
		}
	}

	if (occurance >= p.minimum && occurance <= p.maximum){
		p.valid = true
	}
}



func main() {
	file, err := os.Open("data/day_2_data.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)


	var policies []*policy

	for fileScanner.Scan() {
		aString := fileScanner.Text()


		policy := parsePolicy(aString)
		policies = append(policies, policy)
		policy.validate()

		if err != nil{
			panic(err)
		}

	}

	var validCount uint
	for _, policy := range policies{
		if policy.valid{
			validCount++
		}
	}

	fmt.Printf("Valid Policies: %d", validCount)




	file.Close()

}

func parsePolicy(aString string) *policy{
	policy := &policy{
		minimum: 0,
		maximum: 0,
		letter:  "",
		data:    nil,
	}

	index := strings.Index(aString, " ")
	tokenA := aString[0: index]
	parseA(policy, tokenA)


	remainString := aString[index + 1:]
	index = strings.Index(remainString, ":")
	tokenB := remainString[0: index]

	parseB(policy, tokenB)


	remainString = remainString[index + 1:]
	index = strings.Index(remainString, " ")
	tokenC := remainString[index + 1:]

	parseC(policy, tokenC)
	return policy
}

func parseA(policy *policy, aString string){
	//7-9 l: vslmtglbc

	index := strings.Index(aString, "-")
	tokenA := aString[0: index]
	tokenB := aString[index+1:]

	var err error
	var value int

	value, err = strconv.Atoi(tokenA)
	if err != nil{
		panic(err)
	}
	policy.minimum = uint(value)


	value, err = strconv.Atoi(tokenB)
	if err != nil{
		panic(err)
	}
	policy.maximum = uint(value)


}


func parseB(policy *policy, aString string){
	policy.letter = aString
}




func parseC(policy *policy, aString string){

	stringArray := strings.Split(aString, "")
	policy.data = stringArray

}
