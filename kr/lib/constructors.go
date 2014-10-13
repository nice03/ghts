package lib

import (
	"math/big"
	"time"
)

func NC정수(값 int64) C정수 { return &sC정수64{&s정수64{값}} }
func NV정수(값 int64) V정수 { return &sV정수64{s정수64: &s정수64{값}} }

func NC부호없는_정수(값 uint64) C부호없는_정수 {
	return &sC부호없는_정수64{&s부호없는_정수64{값}}
}
func NV부호없는_정수(값 uint64) V부호없는_정수 {
	return &sV부호없는_정수64{s부호없는_정수64: &s부호없는_정수64{값}}
}

func NC실수(값 float64) C실수 { return &sC실수64{&s실수64{값}} }
func NV실수(값 float64) V실수 { return &sV실수64{s실수64: &s실수64{값}} }

func NC참거짓(값 bool) C참거짓 {
	if 값 {
		return c참
	} else {
		return c거짓
	}
}
func NV참거짓(값 bool) V참거짓 { return &sV참거짓{s참거짓: &s참거짓{값}} }

func NC문자열(값 string) C문자열 { return &sC문자열{값} }

func NC시점(값 time.Time) C시점 { return &sC시점{&s시점{값}} }
func NC시점_문자열(값 string) C시점 {
	시점, 에러 := F문자열2시점(값)

	if 에러 != nil {
		return nil
	}

	return NC시점(시점)
}

func NV시점(값 time.Time) V시점 {
	return &sV시점{s시점: &s시점{값}}
}

func NV시점_문자열(값 string) V시점 {
	시점, 에러 := F문자열2시점(값)

	if 에러 != nil {
		return nil
	}

	return NV시점(시점)
}

func NC정밀수(값 interface{}) C정밀수 {
	if !F숫자형식임(값) && !F문자열형식임(값) {
		return nil
	}

	var 정밀수 *big.Rat

	switch 값.(type) {
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))
	case *sC정밀수:
		return 값.(*sC정밀수) // 상수형은 굳이 새로운 인스턴스를 생성할 필요가 없다.
	case *sV정밀수:
		정밀수 = 값.(*sV정밀수).GRat()
	default:
		F매개변수_안정성_검사(값)

		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))

		if !성공 {
			return nil
		}
	}

	return &sC정밀수{&s정밀수{정밀수}}
}

func NV정밀수(값 interface{}) V정밀수 {
	if !F숫자형식임(값) && !F문자열형식임(값) {
		return nil
	}

	var 정밀수 *big.Rat

	switch 값.(type) {
	case *big.Rat:
		정밀수 = new(big.Rat).Set(값.(*big.Rat))
	case I정밀수:
		정밀수 = 값.(I정밀수).GRat()
	default:
		F매개변수_안정성_검사(값)

		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))

		if !성공 {
			F문자열_출력("common.NV정밀수() : 입력값이 숫자가 아님. %v", 값)

			return nil
		}
	}

	return &sV정밀수{s정밀수: &s정밀수{정밀수}}
}

// 통화
func NC원화(금액 interface{}) C통화  { return NC통화(KRW, 금액) }
func NC달러(금액 interface{}) C통화  { return NC통화(USD, 금액) }
func NC위안화(금액 interface{}) C통화 { return NC통화(CNY, 금액) }
func NC유로화(금액 interface{}) C통화 { return NC통화(EUR, 금액) }

func NC통화(종류 P통화종류, 금액 interface{}) C통화 {
	if !F숫자형식임(금액) && !F문자열형식임(금액) {
		return nil
	}

	v금액 := NV정밀수(금액)

	if v금액 == nil {
		F문자열_출력("예상치 못한 경우. %v", 금액)
		return nil
	}

	c금액 := v금액.S반올림(F통화종류별_정밀도(종류)).G상수형()

	return &sC통화{종류, c금액}
}

// 변수형 생성자
func NV원화(금액 interface{}) V통화  { return NV통화(KRW, 금액) }
func NV달러(금액 interface{}) V통화  { return NV통화(USD, 금액) }
func NV위안화(금액 interface{}) V통화 { return NV통화(CNY, 금액) }
func NV유로화(금액 interface{}) V통화 { return NV통화(EUR, 금액) }

func NV통화(종류 P통화종류, 금액 interface{}) V통화 {
	if !F숫자형식임(금액) && !F문자열형식임(금액) {
		return nil
	}

	v금액 := NV정밀수(금액)

	if v금액 == nil {
		F문자열_출력("NC통화() : 금액이 숫자가 아님. %v", 금액)
		return nil
	}

	v금액 = v금액.S반올림(F통화종류별_정밀도(종류))

	return &sV통화{종류: 종류, 금액: v금액}
}

func NC매개변수(이름 string, 값 interface{}) C매개변수 {
	var 상수형 I상수형 = F상수형(값)

	if 상수형 == nil {
		return nil
	}

	return &sC매개변수{이름, 상수형}
}