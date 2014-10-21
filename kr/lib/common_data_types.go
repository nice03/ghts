// Copyright 2014 UnHa Kim. All rights reserved.
// Use of this source code is governed by a LGPL V3
// license that can be found in the LICENSE file.

package lib

import (
	"math/rand"
	"reflect"
	//"sync"
)

// 종목
type sC종목 struct {
	코드 string
	명칭 string

	출력_문자열 string
}

func (s *sC종목) 상수형임()       {}
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

// 전략 (임시)
type sC전략 struct {
	코드 string
}
func (s *sC전략) 상수형임()       {}
func (s *sC전략) G코드() string { return s.코드 }
func (s *sC전략) String() string { return s.코드 }
func (s *sC전략) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	c := NC문자열("")
	코드 := c.Generate(임의값_생성기, 크기).Interface().(C문자열).G값()

	return reflect.ValueOf(&sC전략{코드: 코드})
}

// 포트폴리오 변동내역
type sC포트폴리오_변동내역 struct {
	키 string
	전략 C전략
	종목 C종목
	롱포지션_변동수량 int64
	숏포지션_변동수량 int64
	시점 time.Time
}

func (s *sC포트폴리오_변동내역) G키() {
	if s.키 == "" {
		s.키 = s.종목.G코드() + "_" + s.전략.G코드()
	}
	
	return s.키
}
func (s *sC포트폴리오_변동내역) G전략() C전략 { return s.전략 }
func (s *sC포트폴리오_변동내역) G종목() C종목 { return s.종목 }
func (s *sC포트폴리오_변동내역) G롱포지션_변동수량() int64 { return s.롱포지션_변동수량 }
func (s *sC포트폴리오_변동내역) G숏포지션_변동수량() int64 { return s.숏포지션_변동수량 }
func (s *sC포트폴리오_변동내역) G시점() time.Time { return s.시점 }
func (s *sC포트폴리오_변동내역) String() {
	문자열 := F포맷된_문자열(
		"전략 %v, 종목 %v, 롱포지션 변동수량 %v, 숏포지션 변동수량 %v, 시점 %v",
		s.G전략().G코드(), s.G종목().G코드(), 
		s.G롱포지션_수량(), s.G숏포지션_수량(),
		s.G시점().Format(P시점_형식))
	
	return 문자열
}

// 포트폴리오 구성요소
type s포트폴리오_구성요소 struct {
	키 string	// 종목코드 + "_" + 전략코드
	전략 C전략
	종목 C종목	
	롱포지션_수량 int64
	숏포지션_수량 int64
}

func (s *s포트폴리오_구성요소) G키() string {
	if s.키 == "" {
		s.키 = s.종목.G코드() + "_" + s.전략.G코드()
	}
	
	return s.키
}
func (s *s포트폴리오_구성요소) G전략() C전략 { return s.전략 }
func (s *s포트폴리오_구성요소) G전략코드() string { return s.전략.G코드() }
func (s *s포트폴리오_구성요소) G종목() C종목 { return s.종목 }
func (s *s포트폴리오_구성요소) G종목코드() string { return s.종목.G코드() }
func (s *s포트폴리오_구성요소) G롱포지션_수량() {
	return atomic.LoadInt64(&s.롱포지션_수량)
}
func (s *s포트폴리오_구성요소) G숏포지션_수량() {
	return atomic.LoadInt64(&s.숏포지션_수량)
}
func (s *s포트폴리오_구성요소) G순_수량() {
	return s.G롱포지션_수량() - s.G숏포지션_수량()
}
func (s *s포트폴리오_구성요소) G총_수량() {
	return s.G롱포지션_수량() + s.G숏포지션_수량()
}
func (s *s포트폴리오_구성요소) String() string {
	문자열 := F포맷된_문자열(
		"전략코드 %v, 종목코드 %v, 롱포지션_수량 %v, 숏포지션_수량 %v",
		s.G전략코드(), s.G종목코드(), 
		s.G롱포지션_수량(), s.G숏포지션_수량())
	
	return 문자열
}

type sC포트폴리오_구성요소 struct { *s포트폴리오_구성요소 }

func (s *sC포트폴리오_구성요소) 상수형임{}
func (s *sC포트폴리오_구성요소) G변수형() V포트폴리오_구성요소 {
	return NV포트폴리오_구성요소(전략 C전략, 종목 C종목, 
				롱포지션_수량, 숏포지션_수량 int64)
}

type sV포트폴리오_구성요소 struct { *s포트폴리오_구성요소 }

func (s *sV포트폴리오_구성요소) 변수형임{}
func (s *sV포트폴리오_구성요소) G상수형() C포트폴리오_구성요소 {
	return &sC포트폴리오_구성요소{
				전략: s.G전략(), 종목: s.G종목(),
				롱포지션_수량: s.G롱포지션_수량(), 
				숏포지션_수량: s.G숏포지션_수량())
}
func (s *sV포트폴리오_구성요소) S변동(변동내역 C포트폴리오_변동내역) error {
	if s.G키() != 변동내역.G키() {	
		에러 := F에러_생성("키 불일치. %s, %s", s.G키(), 변동내역.G키())
		
		F문자열_출력(에러.Error())
		return 에러
	}	
	
	if 변동내역.G롱포지션_변동수량() != 0 {
		atomic.AddInt64(
				&(s.s포트폴리오_구성요소.롱포지션_수량), 
				변동내역.G롱포지션_변동수량())
	}
	
	if 변동내역.G숏포지션_변동수량() != 0 {
		atomic.AddInt64(
				&(s.s포트폴리오_구성요소.숏포지션_수량), 
				변동내역.G숏포지션_변동수량())
	}
}

// 전체 포트폴리오
type s전체_포트폴리오 struct {
	잠금 sync.RWMutex	// V포트폴리오_구성내역 생성할 때 필요함.
	
	전체_저장소 map[string]V포트폴리오_구성내역
	종목별_저장소 map[string](map[string]V포트폴리오_구성내역)	// 키1 : 종목, 키2 : 전략코드
	전략별_저장소 map[string](map[string]V포트폴리오_구성내역)	// 키1 : 전략코드, 키2 : 종목
	
	자본 C통화
}
func (s *s전체_포트폴리오) G구성요소(키 string) V포트폴리오_구성요소 {
	s.잠금.RLock()
	포트폴리오_구성요소, ok = s.전체_저장소[키]
	s.잠금.RUnlock()
	
	if !ok { return nil }
	
	return 포트폴리오_구성요소
}

func (s *s전체_포트폴리오) s추가(전략 C전략, 종목 C종목) (V포트폴리오_구성요소, error) {
	새_구성요소 := NV포트폴리오_구성요소(
					구성요소.G전략코드(), 구성요소.G종목(), 0, 0)
					
	_, ok := s.전체_저장소[새_구성요소.G키()]
	
	if ok {
		에러 := F에러_생성("이미 존재하는 항목입니다. %s", 새_구성요소.G키())
		
		return nil, 에러
	}
	
	s.잠금.Lock()
	defer s.잠금.Unlock()
	
	// 추가하기 전에 최종 확인.
	_, ok := s.전체_저장소[새_구성요소.G키()]
	
	if ok {		
		에러 := F에러_생성("이미 존재하는 항목입니다. %s", 새_구성요소.G키())
		
		return nil, 에러
	} 
	
	// 전체 저장소 설정
	s.전체_저장소[새_구성요소.G키()] = 새_항목
	
	// 종목별 저장소 설정
	종목별_맵, ok := s.종목별_저장소[새_구성요소.G종목코드()]
	if !ok {
		종목별_맵 = make(map[string]V포트폴리오_구성내역)
		s.종목별_저장소[새_구성요소.G종목코드()] = 종목별_맵
	}
	
	종목별_맵[새_구성요소.G전략코드()] = 새_항목
	
	// 전략별 저장소 설정
	전략별_맵, ok := s.전략별_저장소[새_구성요소.G전략코드()]
	if !ok {
		전략별_맵 = make(map[string]V포트폴리오_구성내역)
		s.전략별_저장소[새_구성요소.G전략코드()] = 전략별_맵
	}
	
	전략별_맵[새_구성요소.G종목코드()] = 새_항목
	
	return 새_항목, nil
}


func (s *s전체_포트폴리오) S변동(변동내역 C포트폴리오_변동내역) error {
	F메모("포트폴리오 변동내역 DB에 저장하는 기능 추가할 것.")
	
	var 포트폴리오_구성요소 V포트폴리오_구성요소
	var ok bool
	var 에러 error
	
	// 전체 저장소
	for 반복횟수 := 0 ; 반복횟수 < 100 ; 반복횟수++ {
		포트폴리오_구성요소 = s.G구성요소(변동내역.G키())
		
		if 포트폴리오_구성요소 != nil { break }
	
		포트폴리오_구성요소, 에러 = s.s추가(변동내역.G전략(), 변동내역.G종목())
		
		if 에러 == nil && 
		   포트폴리오_구성요소 != nil { break }
	}
	
	if 포트폴리오_구성요소 == nil {
		에러 := F에러_생성("포트폴리오 구성요소를 읽거나 생성할 수 없습니다.")
		
		return 에러
	}
		
	return 포트폴리오_구성요소.S변동(변동내역)
}

/*
	G종목_모음() []C종목
	G종목별_포트폴리오_모음() []I종목별_포트폴리오
	G종목별_포트폴리오(종목 C종목) I종목별_포트폴리오

	G전략_모음() []C전략
	G전략별_포트폴리오_모음() []I전략별_포트폴리오
	G전략별_포트폴리오(전략코드 string) I전략별_포트폴리오	// 향후 C전략 형식으로 바꾸는 것 고려.

	G롱포지션_금액() C통화
	G숏포지션_금액() C통화
	G순_금액() C통화
	G총_금액() C통화

	G자본() C통화
*/



/*
type s종목별_포트폴리오 struct {
	종목    C종목
	매입수량  int64
	공매도수량 int64
}

func (s *s종목별_포트폴리오) G종목() C종목     { return s.종목 }
func (s *s종목별_포트폴리오) G매입수량() int64 { return atomic.LoadInt64(&(s.매입수량)) }
func (s *s종목별_포트폴리오) G공매도수량() int64 {
	return atomic.LoadInt64(&(s.공매도수량))
}
func (s *s종목별_포트폴리오) G순수량() int64 {
	return s.G매입수량() - s.G공매도수량()
}
func (s *s종목별_포트폴리오) G총수량() int64 {
	return s.G매입수량() + s.G공매도수량()
}
func (s *s종목별_포트폴리오) G매입금액() C통화 {
	F메모("종목별 시세 구하는 'F종목별_시세(종목 C종목) C통화' 함수가 필요함.")

	단가 := NC통화(KRW, 100).G변수형()

	return 단가.S곱하기(s.G매입수량()).G상수형()
}
func (s *s종목별_포트폴리오) G공매도금액() C통화 {
	F메모("종목별 시세 구하는 'F종목별_시세(종목 C종목) C통화' 함수가 필요함.")

	단가 := NC통화(KRW, 100).G변수형()

	return 단가.S곱하기(s.G공매도수량()).G상수형()
}
func (s *s종목별_포트폴리오) G순금액() C통화 {
	F메모("종목별 시세 구하는 'F종목별_시세(종목 C종목) C통화' 함수가 필요함.")

	단가 := NC통화(KRW, 100).G변수형()

	return 단가.S곱하기(s.G순수량()).G상수형()
}
func (s *s종목별_포트폴리오) G총금액() C통화 {
	F메모("종목별 시세 구하는 함수 구현한 후 단가 구하는 소스를 수정할 것.")

	단가 := NC통화(KRW, 100).G변수형() // F종목별_시세(종목 C종목) C통화 이 필요함.

	return 단가.S곱하기(s.G총수량()).G상수형()
}
func (s *s종목별_포트폴리오) String() string {
	매입수량 := s.G매입수량()
	공매도수량 := s.G공매도수량()

	return F포맷된_문자열("%s : 매입수량 %v, 공매도수량 %v, 순수량 %v, 총수량 %v\n",
		s.종목.String(), s.G매입수량(), s.G공매도수량(), s.G순수량(), s.G총수량())
}
func (s *s종목별_포트폴리오) generate(임의값_생성기 *rand.Rand, 크기 int) *s종목별_포트폴리오 {
	종목 := NC종목("", "").Generate(임의값_생성기, 크기).Interface().(C종목)
	매입수량 := NC부호없는_정수(int64(임의값_생성기.Int63()))
	공매도수량 := NC부호없는_정수(int64(임의값_생성기.Int63()))

	return &s종목별_포트폴리오{종목: 종목, 매입수량: 매입수량, 공매도수량: 공매도수량}
}

type sC종목별_포트폴리오 struct{ *s종목별_포트폴리오 }

func (s *sC종목별_포트폴리오) 상수형임() {}
func (s *sC종목별_포트폴리오) G변수형() V종목별_포트폴리오 {
	return NV종목별_포트폴리오(s.종목, s.매입수량.G값(), s.공매도수량.G값())
}
func (s *sC종목별_포트폴리오) Generate(임의값_생성기 *rand.Rand, 크기 int) reflect.Value {
	return reflect.ValueOf(
		&sC종목별_포트폴리오{s종목별_포트폴리오: s.generate(임의값_생성기, 크기)})
}

type sV종목별_포트폴리오 struct {
	//잠금 sync.RWMutex	// sync.atomic을 이용하면 Mutex가 필요없음.
	*s종목별_포트폴리오
}

func (s *sV종목별_포트폴리오) 변수형임() {}
func (s *sV종목별_포트폴리오) G상수형() C종목별_포트폴리오 {
	return &sC종목별_포트폴리오{&s종목별_포트폴리오{
		종목: s.G종목(), 매입수량: s.G매입수량(), 공매도수량: s.G공매도수량()}}
}

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
			에러 := F에러_생성("'매입수량_새로운값'이 0보다 작음. "+
				"매입수량_원래값 %v, 매입수량_변동값 %v, 매입수량_새로운값 %v",
				매입수량_원래값, 매입수량_변동값, 매입수량_새로운값)

			F문자열_출력(에러.Error())

			return 에러
		}

		if 공매도수량_새로운값 < 0 {
			에러 := F에러_생성("'공매도수량_새로운값'이 0보다 작음. "+
				"공매도수량_원래값 %v, 공매도수량_변동값 %v, 공매도수량_새로운값 %v",
				공매도수량_원래값, 공매도수량_변동값, 공매도수량_새로운값)

			F문자열_출력(에러.Error())

			return 에러
		}

		var c매입수량_새로운값, c공매도수량_새로운값 C부호없는_정수

		if 매입수량_변동값 == 0 {
			c매입수량_새로운값 = 원래값.매입수량
		} else {
			c매입수량_새로운값 = NC부호없는_정수(int64(매입수량_새로운값))
		}

		if 공매도수량_변동값 == 0 {
			c공매도수량_새로운값 = 원래값.공매도수량
		} else {
			c공매도수량_새로운값 = NC부호없는_정수(int64(공매도수량_새로운값))
		}

		종목별_포트폴리오_새로운값 := &s종목별_포트폴리오{
			종목:    원래값.종목,
			매입수량:  c매입수량_새로운값,
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

// 모든 데이터의 변경은 channel을 통해서 주고 받음.
// 필요할 경우 내부적으로 channel을 사용해서 데이터를 전달하는 메소드를 추가할 수 있음.
// 쓰기 위주로 짤 것인가? 읽기 위주로 짤 것인가?
// 종목별로 분류할 것인가? 관련 전략별로 분류할 것인가?
// 아니면 종목과 관련 전략 둘 다 기준으로 분류할 것인가?
// 포트폴리오 구성요소에는 둘 다 필요하다.
// 로그와 컴팩션 된 결과물을 분리?

type s전체_포트폴리오 struct {
	저장소 map[string]V종목별_포트폴리오
}

func (s s포트폴리오) S추가(종목별_포트폴리오 []C종목별_포트폴리오) error {
	return nil
}
func (s s포트폴리오) S매입수량_변동(종목 C종목, 변동값 int64) error  { return nil }
func (s s포트폴리오) S공매도수량_변동(종목 C종목, 변동값 int64) error { return nil }

/*
type sC포트폴리오 struct {
	// 'I안전한_맵'의 'string-C종목별_포트폴리오' 형태의 type-safe한 자료형 만드는 것 고려.
	// wrapper가 형변환을 해 주던 지, ps.Map 소스코드를 변경하던 지 둘 중 편한 것으로.
	저장소 I안전한_맵
}

func (s *sC포트폴리오) 상수형임() {}
func (s *sC포트폴리오) G변수형() V포트폴리오 {}

func (s *sC포트폴리오) G종목_모음() []C종목
func (s *sC포트폴리오) G존재함(종목 C종목) bool
func (s *sC포트폴리오) G종목별_포트폴리오(종목 C종목) C종목별_포트폴리오
func (s *sC포트폴리오) G전체_포트폴리오() []C종목별_포트폴리오

func (s *sC포트폴리오) G매입수량(종목 C종목) int64
func (s *sC포트폴리오) G공매도수량(종목 C종목) int64
func (s *sC포트폴리오) G순수량(종목 C종목) int64
func (s *sC포트폴리오) G총수량(종목 C종목) int64

func (s *sC포트폴리오) G매입금액(종목 C종목) C통화
func (s *sC포트폴리오) G공매도금액(종목 C종목) C통화
func (s *sC포트폴리오) G순금액(종목 C종목) C통화
func (s *sC포트폴리오) G총금액(종목 C종목) C통화

func (s *sC포트폴리오) G전체_매입금액() C통화
func (s *sC포트폴리오) G전체_공매도금액() C통화
func (s *sC포트폴리오) G전체_순금액() C통화
func (s *sC포트폴리오) G전체_총금액() C통화
func (s *sC포트폴리오) String() string
func (s *sC포트폴리오) Generate(임의값_생성기 *rand.Rand, 크기 int)


type sV포트폴리오 struct {
	// "V종목별_포트폴리오'항목을 "추가"하는 루틴만 사용하는 뮤텍스.
	// 각 V종목별_포트폴리오 항목을 수정하는 루틴은 V종목별_포트폴리오에 포함된 개별 뮤텍스를 사용할 것.
	잠금 sync.Mutex
	저장소 map[string]V종목별_포트폴리오
}

func (s *sV포트폴리오) 변수형임() {}
func (s *sV포트폴리오) G상수형() C포트폴리오 {}

func (s *sV포트폴리오) G종목_모음() []C종목
func (s *sV포트폴리오) G존재함(종목 C종목) bool
func (s *sV포트폴리오) G종목별_포트폴리오(종목 C종목) C종목별_포트폴리오
func (s *sV포트폴리오) G전체_포트폴리오() []C종목별_포트폴리오

func (s *sV포트폴리오) G매입수량(종목 C종목) int64
func (s *sV포트폴리오) G공매도수량(종목 C종목) int64
func (s *sV포트폴리오) G순수량(종목 C종목) int64
func (s *sV포트폴리오) G총수량(종목 C종목) int64

func (s *sV포트폴리오) G매입금액(종목 C종목) C통화
func (s *sV포트폴리오) G공매도금액(종목 C종목) C통화
func (s *sV포트폴리오) G순금액(종목 C종목) C통화
func (s *sV포트폴리오) G총금액(종목 C종목) C통화

func (s *sV포트폴리오) G전체_매입금액() C통화
func (s *sV포트폴리오) G전체_공매도금액() C통화
func (s *sV포트폴리오) G전체_순금액() C통화
func (s *sV포트폴리오) G전체_총금액() C통화
func (s *sV포트폴리오) String() string
func (s *sV포트폴리오) Generate(임의값_생성기 *rand.Rand, 크기 int)

*/
