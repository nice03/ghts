package lib

import (
	"github.com/gh-system/ghts/dep/ps"
	"math/big"
	"time"
)

func N반환값(값 I가변형, 에러 error) I반환값 { return &s반환값{ 값: 값, 에러: 에러 } }

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

func NC시점(값 I가변형) C시점 {
	if !F시점형식임(값) && !F문자열형식임(값) {
		return nil
	}
	
	switch 값.(type) {
	case time.Time:
		return &sC시점{&s시점{값.(time.Time)}}
	case *sC시점, *sV시점:
		return &sC시점{&s시점{값.(I시점).G값()}}
	case string:
		시점, 에러 := F문자열2시점(값.(string))

		if 에러 != nil {
			return nil
		}
		
		return &sC시점{&s시점{시점}}
	case *sC문자열:
		시점, 에러 := F문자열2시점(값.(*sC문자열).G값())

		if 에러 != nil {
			return nil
		}
		
		return &sC시점{&s시점{시점}}
	default:
		F문자열_출력("예상치 못한 입력값 형식. %s", F값_확인_문자열(값))
		
		return nil
	}
}

func NV시점(값 I가변형) V시점 {
	if !F시점형식임(값) && !F문자열형식임(값) {
		return nil
	}
	
	switch 값.(type) {
	case time.Time:
		return &sV시점{s시점: &s시점{값.(time.Time)}}
	case *sC시점, *sV시점:
		return &sV시점{s시점: &s시점{값.(I시점).G값()}}
	case string:
		시점, 에러 := F문자열2시점(값.(string))

		if 에러 != nil {
			return nil
		}
		
		return &sV시점{s시점: &s시점{시점}}
	case *sC문자열:
		시점, 에러 := F문자열2시점(값.(*sC문자열).G값())

		if 에러 != nil {
			return nil
		}
		
		return &sV시점{s시점: &s시점{시점}}
	default:
		F문자열_출력("예상치 못한 입력값 형식. %s", F값_확인_문자열(값))
		
		return nil
	}
}

func NC정밀수(값 I가변형) C정밀수 {
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
		F매개변수_안전성_검사(값)

		var 성공 bool
		정밀수, 성공 = new(big.Rat).SetString(F문자열(값))

		if !성공 {
			return nil
		}
	}

	return &sC정밀수{&s정밀수{정밀수}}
}

func NV정밀수(값 I가변형) V정밀수 {
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
		F매개변수_안전성_검사(값)

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
func NC원화(금액 I가변형) C통화  { return NC통화(KRW, 금액) }
func NC달러(금액 I가변형) C통화  { return NC통화(USD, 금액) }
func NC위안화(금액 I가변형) C통화 { return NC통화(CNY, 금액) }
func NC유로화(금액 I가변형) C통화 { return NC통화(EUR, 금액) }

func NC통화(종류 P통화종류, 금액 I가변형) C통화 {
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
func NV원화(금액 I가변형) V통화  { return NV통화(KRW, 금액) }
func NV달러(금액 I가변형) V통화  { return NV통화(USD, 금액) }
func NV위안화(금액 I가변형) V통화 { return NV통화(CNY, 금액) }
func NV유로화(금액 I가변형) V통화 { return NV통화(EUR, 금액) }

func NV통화(종류 P통화종류, 금액 I가변형) V통화 {
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

func NC매개변수(이름 string, 값 I가변형) C매개변수 {
	F매개변수_안전성_검사(값)
	
	if _, ok := 값.([]I가변형); ok {
		값 = F중첩된_외부_슬라이스_제거(값.([]I가변형))
	}
	
	switch 값.(type) {
	case *sC매개변수:
		c매개변수 := 값.(*sC매개변수)
		이름 = c매개변수.G이름()
		값 = c매개변수.G값()
	case []I가변형:
		F중첩된_외부_슬라이스_제거(값.([]I가변형))
	}
	
	return &sC매개변수{이름, 값}
}

func N안전한_배열() I안전한_배열 {
	return &s안전한_배열{ps.NewList()}
}

func N안전한_맵() I안전한_맵 {
	return &s안전한_맵{ps.NewMap()}
}

func NC종목(코드, 명칭 string) C종목 { return &sC종목{코드: 코드, 명칭: 명칭} }

/*
func NC종목별_포트폴리오(종목 C종목, 매입수량, 매도수량 uint64) {
	F메모("NC종목별_포트폴리오가 필요한가?, 그냥 V종목별_포트폴리오.G상수형()으로 충분하지 않을까?")
	&sC종목별_포트폴리오{종목: 종목, 매입수량: 매입수량, 매도수량: 매도수량}
} */

func NV종목별_포트폴리오(종목 C종목, 매입수량, 공매도수량 uint64) V종목별_포트폴리오 {
	c매입수량 := NC부호없는_정수(매입수량)
	c공매도수량 := NC부호없는_정수(공매도수량)
	
	return &sV종목별_포트폴리오{s종목별_포트폴리오:
				&s종목별_포트폴리오{종목: 종목, 매입수량: c매입수량, 공매도수량: c공매도수량}}
}

/*
func NC포트폴리오(종목별_포트폴리오_모음 []종목별_포트폴리오_모음) C포트폴리오 {
	F메모("NC포트폴리오가 필요한가?, 그냥 V포트폴리오.G상수형()으로 충분하지 않을까?")
	return nil
} */

func NV포트폴리오(통화종류 P통화종류) V포트폴리오 {
	F메모("sV포트폴리오 작성")
	return nil
	//return &sV포트폴리오{저장소: make(map[string]V종목별_포트폴리오)}
}

