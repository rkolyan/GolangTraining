package main

import (
	"strings"
	bg "math/big"
)

func find_number_of_operation(A, B, S, K_str string) int64 {
	A = strings.ToLower(A);
	B = strings.ToLower(B);
	S = strings.ToLower(S);
	
	//1 Ищем количество совпадений с S в A
	A_count := strings.Count(A, S);
	
	//2 Ищем количество совпадений с S в B
	B_count := strings.Count(B, S);
	
	//3 Ищем количество совпадений с S в AB, BA, BB (AA - невозможно)
	AB_count := strings.Count(strings.Join([]string{A, B}, ""), S) - (A_count + B_count);
	BA_count := strings.Count(strings.Join([]string{B, A}, ""), S) - (A_count + B_count);
	BB_count := strings.Count(strings.Join([]string{B, B}, ""), S) - (B_count + B_count);
	if !A_count && !B_count && !AB_count && !BA_count && !BB_count {
		return 0;
	}
	bad_result.SetInt64(-1);

	K := bg.NewInt(0);
	count := bg.NewInt(0);
	fib1 := bg.NewInt(1);
	fib2 := bg.NewInt(1);
	tmp := bg.NewInt(0);

	fib1.SetInt64(1);
	fib2.SetInt64(1);

	K, success := K.SetString(K_str, 0);
	if !success {
		return bad_result;
	}

	flag := false;
	for count.Cmp(K) < 0 {
		tmp.Set(fib2);
		fib2.Add(fib2, fib1);
		fib1.Set(tmp);
		count.Add()
	}
	if count.Cmp(K) == 0 {
		return count;
	}
	return bad_result;
}
