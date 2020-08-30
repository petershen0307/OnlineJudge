package q1169

import (
	"strconv"
	"strings"
)

//transactions = ["alice,20,800,mtv","alice,50,100,beijing"]

type transaction struct {
	name    string
	minute  string
	account string
	city    string
	v       string
	invalid bool
}

func invalidTransactions(transactions []string) []string {
	// invalid
	//1. account > 1000
	//2. two city <= 60 min
	invalidTrans := make([]string, 0)
	nameMap := make(map[string][]transaction, 0)
	for _, trans := range transactions {
		tSplit := strings.Split(trans, ",")
		nameMap[tSplit[0]] = append(nameMap[tSplit[0]], transaction{
			name:    tSplit[0],
			minute:  tSplit[1],
			account: tSplit[2],
			city:    tSplit[3],
			v:       trans,
			invalid: false,
		})
	}
	accountInvalid := func(trans transaction) bool {
		if i, _ := strconv.Atoi(trans.account); i > 1000 {
			return true
		}
		return false
	}
	for _, vTrans := range nameMap {
		for i := 0; i != len(vTrans); i++ {
			if !vTrans[i].invalid {
				vTrans[i].invalid = accountInvalid(vTrans[i])
			}
			for j := i; j != len(vTrans); j++ {
				if vTrans[i].city != vTrans[j].city {
					iM, _ := strconv.Atoi(vTrans[i].minute)
					jM, _ := strconv.Atoi(vTrans[j].minute)
					if (iM-jM) >= 0 && (iM-jM) <= 60 {
						vTrans[i].invalid = true
						vTrans[j].invalid = true
					}
					if (iM-jM) < 0 && (iM-jM) >= -60 {
						vTrans[i].invalid = true
						vTrans[j].invalid = true
					}
				}
			}
			if vTrans[i].invalid {
				invalidTrans = append(invalidTrans, vTrans[i].v)
			}
		}
	}
	return invalidTrans
}
