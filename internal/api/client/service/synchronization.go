package service

func Sync(userId int64) ([][]string, [][]string) {
	dataTblText := [][]string{{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"},
		{"NAME_1", "DATA", "DESCRIPTION", "01/01/2000", "01/01/2000"}}
	dataTblCard := [][]string{{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"},
		{"NAME_1", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "01/01/2000", "01/01/2000", "01/01/2000"}}

	return dataTblText, dataTblCard
}
