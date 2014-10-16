﻿package lib

import (
	"math/rand"
	"reflect"
	"sync"
)

type sC종목 struct {
	코드 string
	명칭 string
	
	출력_문자열 string
}

func (s *sC종목) 상수형임() {}
func (s *sC종목) G코드() string { return s.코드 }
func (s *sC종목) G명칭() string { return s.명칭 }
func (s *sC종목) String() string {
	if s.출력_문자열 == "" {
		s.출력_문자열 = s.코드 + " " + s.명칭
	}
	
	return s.출력_문자열
}
func (s *sC종목) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	c := NC문자열("")
	코드 := c.Generate(임의값_생성기, 크기).Interface().(C문자열).G값()
	명칭 := c.Generate(임의값_생성기, 크기).Interface().(C문자열).G값()
	
	return reflect.ValueOf(&sC종목{코드: 코드, 명칭: 명칭})
}

type s종목별_포트폴리오 struct {
	종목 C종목
	매입수량 C부호없는_정수
	공매도수량 C부호없는_정수
}
func (s *s종목별_포트폴리오) G종목() C종목 { return s.종목 }
func (s *s종목별_포트폴리오) G매입수량() uint64 { return s.매입수량.G값() }
func (s *s종목별_포트폴리오) G공매도수량() uint64 { return s.공매도수량.G값() }
func (s *s종목별_포트폴리오) G순수량() int64 {
	return int64(s.G매입수량() - s.G공매도수량())
}
func (s *s종목별_포트폴리오) G총수량() uint64 { return s.G매입수량() + s.G공매도수량() }
func (s *s종목별_포트폴리오) G매입금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")
	
	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.
	
	return 단가.S곱하기(s.G매입수량()).G상수형()
}
func (s *s종목별_포트폴리오) G공매도금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.

	return 단가.S곱하기(s.G공매도수량()).G상수형()
}
func (s *s종목별_포트폴리오) G순금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.

	return 단가.S곱하기(s.G매입수량() - s.G공매도수량()).G상수형()
}
func (s *s종목별_포트폴리오) G총금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.

	return 단가.S곱하기(s.G매입수량() + s.G공매도수량()).G상수형()
}
func (s *s종목별_포트폴리오) String() string {
	매입수량 := s.G매입수량()
	공매도수량 := s.G공매도수량()

	return F포맷된_문자열("%s : 매입수량 %v, 공매도수량 %v, 순수량 %v, 총수량 %v\n",
				s.종목.String(), 매입수량, 공매도수량, 
				매입수량 - 공매도수량, 매입수량 + 공매도수량)
}
func (s *s종목별_포트폴리오) generate(임의값_생성기 *rand.Rand, 크기 int) *s종목별_포트폴리오 {
	종목 := NC종목("", "").Generate(임의값_생성기, 크기).Interface().(C종목)
	매입수량 := NC부호없는_정수(uint64(임의값_생성기.Int63()))
	공매도수량 := NC부호없는_정수(uint64(임의값_생성기.Int63()))
	
	return &s종목별_포트폴리오{종목: 종목, 매입수량: 매입수량, 공매도수량: 공매도수량}
}


type sC종목별_포트폴리오 struct { *s종목별_포트폴리오 }

func (s *sC종목별_포트폴리오) 상수형임() {}
func (s *sC종목별_포트폴리오) G변수형() V종목별_포트폴리오 {
	return NV종목별_포트폴리오(s.종목, s.매입수량.G값(), s.공매도수량.G값())
}
func (s *sC종목별_포트폴리오) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(
				&sC종목별_포트폴리오{s종목별_포트폴리오: s.generate(임의값_생성기, 크기)})
}

// 모든 mutation은 s종목별_포트폴리오 전체를 바꾸어서 매입수량과 공매도 수량의 일관성을 유지하도록 한다.
type sV종목별_포트폴리오 struct {
	잠금 sync.RWMutex
	*s종목별_포트폴리오
}

func (s *sV종목별_포트폴리오) 변수형임() {}
func (s *sV종목별_포트폴리오) G상수형() C종목별_포트폴리오 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return &sC종목별_포트폴리오{&s종목별_포트폴리오{
				종목: 값.G종목(), 매입수량: 값.매입수량, 공매도수량: 값.공매도수량}}
}
func (s *sV종목별_포트폴리오) G종목() C종목 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.G종목()
}
func (s *sV종목별_포트폴리오) G매입수량() uint64 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.G매입수량()
}
func (s *sV종목별_포트폴리오) G공매도수량() uint64 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.G공매도수량()
}
func (s *sV종목별_포트폴리오) G순수량() int64 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.G순수량()
}
func (s *sV종목별_포트폴리오) G총수량() uint64 {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.G총수량()
}
func (s *sV종목별_포트폴리오) G매입금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")
	
	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.
	
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 단가.S곱하기(값.G매입수량()).G상수형()
}
func (s *sV종목별_포트폴리오) G공매도금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.
	
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()

	return 단가.S곱하기(값.G공매도수량()).G상수형()
}
func (s *sV종목별_포트폴리오) G순금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.
	
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()

	return 단가.S곱하기(값.G순수량()).G상수형()
}
func (s *sV종목별_포트폴리오) G총금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형()	// F종목별_시세(종목 C종목) C통화 이 필요함.

	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()

	return 단가.S곱하기(값.G총수량()).G상수형()
}
// 내부 종목별_포트폴리오 인스턴스를 통째로 교체하여서,
// 읽기 메소드들이 일부분만 업데이트 된 상태를 읽는 상황을 원천적으로 봉쇄함.
func (s *sV종목별_포트폴리오) s수량_변동(매입수량_변동값, 공매도수량_변동값 int64) error {
	반복횟수 := 0
	
	for {
		s.잠금.RLock()
		원래값 := s.s종목별_포트폴리오
		s.잠금.RUnlock()
		
		매입수량_원래값 := 원래값.G매입수량()
		공매도수량_원래값 := 원래값.G공매도수량()
		
		매입수량_새로운값 := int64(매입수량_원래값) + 매입수량_변동값
		공매도수량_새로운값 := int64(공매도수량_원래값) + 공매도수량_변동값
		
		if 매입수량_새로운값 < 0 {
			에러 := F에러_생성("'매입수량_새로운값'이 0보다 작음. " + 
							"매입수량_원래값 %v, 매입수량_변동값 %v, 매입수량_새로운값 %v", 
							매입수량_원래값, 매입수량_변동값, 매입수량_새로운값)
						
			F문자열_출력(에러.Error())
			
			return 에러
		}
		
		if 공매도수량_새로운값 < 0 {
			에러 := F에러_생성("'공매도수량_새로운값'이 0보다 작음. " + 
							"공매도수량_원래값 %v, 공매도수량_변동값 %v, 공매도수량_새로운값 %v", 
							공매도수량_원래값, 공매도수량_변동값, 공매도수량_새로운값)
						
			F문자열_출력(에러.Error())
			
			return 에러
		}
		
		var c매입수량_새로운값, c공매도수량_새로운값 C부호없는_정수
		
		if 매입수량_변동값 == 0 {
			c매입수량_새로운값 = 원래값.매입수량
		} else {
			c매입수량_새로운값 = NC부호없는_정수(uint64(매입수량_새로운값))
		}
		
		if 공매도수량_변동값 == 0 {
			c공매도수량_새로운값 = 원래값.공매도수량
		} else {
			c공매도수량_새로운값 = NC부호없는_정수(uint64(공매도수량_새로운값))
		}
		
		종목별_포트폴리오_새로운값 := &s종목별_포트폴리오{
									종목: 원래값.종목, 
									매입수량: c매입수량_새로운값, 
									공매도수량: c공매도수량_새로운값}
		
		
		s.잠금.Lock()
		defer s.잠금.Unlock()
		
		if s.G매입수량() == 매입수량_원래값 &&
			s.G공매도수량() == 공매도수량_원래값 {
			// 다른 goroutine이 내부값을 변경하지 않았음.
			// 내부 s종목별_포트폴리오를 통째로 새로운 값으로 교체.
			// 다른 읽기전용 메소드들이 매입수량과 공매도수량의 일관된 값을 읽도록 보장함.
			s.s종목별_포트폴리오 = 종목별_포트폴리오_새로운값
			
			return nil
		}

		// 다른 goroutine에서 값을 이미 변경했음.
		// 새로운 종목별 포트폴리오 값에 대해서 변경 재시도.
		F잠시_대기(반복횟수)
	}
}
func (s *sV종목별_포트폴리오) S매입수량_변동(변동값 int64) error {
	return s.s수량_변동(변동값, 0)
}
func (s *sV종목별_포트폴리오) S공매도수량_변동(변동값 int64) error {
	return s.s수량_변동(0, 변동값)
}
func (s *sV종목별_포트폴리오) String() string {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return 값.String()
}
func (s *sV종목별_포트폴리오) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	var 값 *s종목별_포트폴리오
	s.잠금.RLock()
	값 = s.s종목별_포트폴리오
	s.잠금.RUnlock()
	
	return reflect.ValueOf(
				&sV종목별_포트폴리오{s종목별_포트폴리오: 값.generate(임의값_생성기, 크기)})
}
