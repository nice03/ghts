package lib

import (
	"fmt"
	"strings"
)

func init() {
	나를_위한_문구()
	//메모()
}

const (
	P참  bool = true
	P거짓 bool = false

	P시점_형식 string = "2006-01-02 15:04:05.000000000 (MST) Mon -0700"
	P일자_형식 string = "2006-01-02"

	P매개변수_안전성_검사_건너뛰기 = P거짓
)

const (
	KRW P통화종류 = iota
	USD
	CNY
	EUR

	INVALID_CURRENCY_TYPE P통화종류 = P통화종류(-1)
)

var (
	테스트_모드      V참거짓 = NV참거짓(false)
	문자열_출력_일시정시 V참거짓 = NV참거짓(false)

	c참  C참거짓 = &sC참거짓{&s참거짓{true}}
	c거짓 C참거짓 = &sC참거짓{&s참거짓{false}}

	문자열_후보값_모음 []string = strings.Split(
		"1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"+
			"~!@#$%^&*()_+|';:/?.,<>`가나다라마바사하자차카타파하", "")

	// Exponential Back-off. 재시도 하기 전에 기다리는 대기시간 (단위는 나노초).
	대기시간_한도 = [...]int64{1, 3, 7, 15, 31, 63, 127, 255, 511,
		1023, 2047, 4095, 8191, 16383, 32767,
		65535, 131071, 262143, 524287, 1048575,
		2097151, 4194303, 8388607, 16777215, 33554431,
		67108863, 100000000}
)

func 나를_위한_문구() {
	fmt.Println("")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("	쉽고 간단하게, 테스트로 검증해 가면서 마음 편하게.")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("")

}

func 메모() {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("                  메            모")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("")
	fmt.Println("락을 피할 수 있는 10가지 방법.")
	fmt.Println("")
	fmt.Println("1. 코드가 아닌 데이터를 보호하라.")
	fmt.Println("2. 락을 사용한 부분에서 비싼 계산을 하지 말아라.")
	fmt.Println("3. 락을 분리하라.")
	fmt.Println("4. interlocked이나 atomic 작업을 사용하라.")
	fmt.Println("5. 동기화 된 데이터 구조를 사용하라. ex) lock-free 메시지 큐")
	fmt.Println("6. 가능하다면 읽기-쓰기 락을 사용하라.")
	fmt.Println("7. 가능하다면 읽기 전용 immutable 데이터를 사용하라.")
	fmt.Println("8. 객체 풀링을 피하라. (Go언어에서는 명시적으로 안전하다고 해 놨던 데...)")
	fmt.Println("9. 지역 변수나 개별 쓰레드에 국한된 로컬 변수를 사용하라.")
	fmt.Println("10. 핫스팟(자주 변경되어야 하는 공용 리소스)을 피하라.")
	fmt.Println("")
	fmt.Println("TODO 개요")
	fmt.Println("	- 기본 자료형.")
	fmt.Println("		: immutable 내지 persistent 리스트와 맵 작성할 것.")
	fmt.Println("		  그리고, 이러한 자료형을 이용해서 매개변수 집합체 작성.")
	fmt.Println("	- 기타 공용 자료형")
	fmt.Println("	- 가격정보 취득 tool")
	fmt.Println("	- 포트폴리오 관리.")
	fmt.Println("			Mutex? Channel?")
	fmt.Println("			기본 자료형에는 Mutex가 낫지만,")
	fmt.Println("			여러 연산이 포함되는 포트폴리오 관리에는 Channel이 나을 듯 함.")
	fmt.Println("	- 위험관리")
	fmt.Println("	- 전략, 서버,  가격정보 취득")
	fmt.Println("	- TA-Lib 바인딩 및 테스트 케이스 작성.")
	fmt.Println("")
	fmt.Println("TODO : lib.문자열_후보값_모음 에 한자도 포함시킬 것.")
	fmt.Println("TODO : lib.문자열_후보값_모음, lib.대기시간_한도 상수형 슬라이스로 대체.")
	fmt.Println("")
	fmt.Println("TODO : 자주 사용되는 함수 중 panic 가능성이 높은 함수에 recover() 추가.")
	fmt.Println("")
	fmt.Println("TODO : 재귀적으로 모든 내부값의 상태를 기록하는 메소드.")
	fmt.Println("		F값_일치(), F공유해도_안전함()에서 사용할 예정.")
	fmt.Println("		1. 모든 내부 필드에 대해서 값을 복사해서 식별정보와 함께 슬라이스에 보관.")
	fmt.Println("			a. 복사할 수 있는 형식은 값을 복사해서 식별정보와 함께 보관.")
	fmt.Println("				식별정보는 ")
	fmt.Println("				- 필드 인덱스(reflect.Value.FieldByIndex().")
	fmt.Println("				- 형식")
	fmt.Println("				- 포인터나 인터페이스등 참조형은 그것이 가리키는 값을 기록.")
	fmt.Println("			b. 복사할 수 없는 형식은 복사할 수 있는 형식에 도달할 때까지 ")
	fmt.Println("				각 멤버에 대해서 재귀적으로 호출.")
	fmt.Println("				재귀호출할 때 식별정보도 같이 넘겨줘야 함.")
	fmt.Println("			c. 참조 순환루프으로 인한 문제를 방지하기 위해서,")
	fmt.Println("				이미 복사를 시도한 값과 같은 주소값를 가지고 있는 지 확인.")
	fmt.Println("				무한루프라고 판단되면 에러발생 및 안전하지 않다고 판단.")
	fmt.Println("			d. 이 과정을 거치면 복사 완료 된다고 판단됨.")
	fmt.Println("		2. 값 복사를 마쳤으면, 모든 메소드를 실행.")
	fmt.Println("		4. 복사된 값과 메소드 실행 후의 값을 비교.")
	fmt.Println("			복사값과 원래 값을 불러올 때 위에서 언급한 식별정보를 이용함.")
	fmt.Println("		5. 이러한 과정을 내부 멤버필드에 대해서도 재귀적으로 반복.")
	fmt.Println("			단, complex64, complex128을 제외한 모든 primitive들은 ")
	fmt.Println("			공유해도 안전하다고 판단함.")
	fmt.Println("			sC정밀수로 감싼 big.Rat은 ")
	fmt.Println("			비록, 내부적으로 값을 변경하는 메소드가 존재하더라도 ")
	fmt.Println("			공유해도 안전하다고 판단해도 됨.)")
	fmt.Println("			complex64, complex128에 대해서는 복사에 독립적인 지 확인필요.")
	fmt.Println("")
	fmt.Println("TODO : F값_같음() : Real Value DeepEqual 구현 계획.")
	fmt.Println("		기본 계획")
	fmt.Println("		- 두 값의 내부 필드값을 기록한 슬라이스 확보.")
	fmt.Println("		- 두 슬라이스에 대해서 비교.")
	fmt.Println("")
	fmt.Println("TODO : F공유해도_안전함() 개선 계획.")
	fmt.Println("		기술적으로는 어렵지만 immutability를 확인하는 가장 확실한 방법.")
	fmt.Println("		- 외부에 공개된 필드는 없어야 한다.")
	fmt.Println("		- 내부 멤버 필드들도 모두 immutable 해야 한다.")
	fmt.Println("		  (일부분만 변경되어도 더이상 immutable이 아님.)")
	fmt.Println("		- 모든 메소드를 실행해도 내부값이 변하지 않아야 한다.")
	fmt.Println("			(내부 멤버 구조체의 메소드도 포함.)")
	fmt.Println("			a. 원래값을 기록.")
	fmt.Println("			b. 모든 메소드를 실행.")
	fmt.Println("			c. 값이 변경되었는 지 확인.")
	fmt.Println("			d. 내부 멤버 필드에 대해서도 재귀적으로 반복 수행.")
	fmt.Println("		  		만약 안전하지 않은 멤버 필드가 하나라도 있으면 false")
	fmt.Println("		- primitive value 형식은 Go언어의 CallByValue에 의해서 안전함.")
	fmt.Println("		  *sC큰정수, *sC정밀수등 안전하다고 알려진 형식에 대해서는 ")
	fmt.Println("		  내부값인 *big.Int, *big.Rat이 안전하지 않아도 안전하다고 판단함.")
	fmt.Println("")
	fmt.Println("TODO : I상수형 자동 테스트.")
	fmt.Println("		1. 패키지에 I상수형을 구현하는 자료형 파악하기. pkgreflect를 이용.")
	fmt.Println("		2. 테스트용 인스턴스 확보.")
	fmt.Println("		 	I상수형.상수형임()으로 테스트용 인스턴스 확보한 후,")
	fmt.Println("			테스트용 인스턴스와 원래 주어진 값이 같은 형식인지 확인할 것.")
	fmt.Println("		3. F공유해도_안전함()으로 검사.")
	fmt.Println("")
	fmt.Println("TODO : 메모리 profiling 이후 상수형 구조체로 인해서 메모리 낭비가")
	fmt.Println("		너무 심하다고 판단되면 가장 자주 사용되는 상수형 값들을 ")
	fmt.Println("		캐쉬해서 생성자에서 캐쉬된 게 있는 지 확인해서 중복생성을 줄여주자.")
	fmt.Println("		참거짓은 이미 그런 식으로 되어 있으며,")
	fmt.Println("		숫자도 일정 범위까지는 값과 인덱스가 일치하도로 슬라이스에 저장하면 됨.")
	fmt.Println("		종목도 마찬가지.")
	fmt.Println("		이렇게 하면 상수형을 많이 사용해도 GC의 필요성이 줄어든다.")
	fmt.Println("")
	fmt.Println("TODO : 나눗셈의 나머지값 처리방식을 지정해야 할 경우 I정밀수 내부 구현을 ")
	fmt.Println("		speter.decimal로 변경할 것.")
	fmt.Println("")
	fmt.Println("TODO : V통화 연산함수에서 통화종류가 다를 경우 어떻게 처리할 것인지 생각해볼 것.")
	fmt.Println("		일단, 연산 불가능하도록 하고, 서로 다른 통화끼리 연산은 환율로 환산을")
	fmt.Println("		거친 후에만 가능하도록 할 것. 환율 환산 함수는 별도로 만들어야 함.")
	fmt.Println("")
	fmt.Println("TODO : TestS종목별_포트폴리오(), TestS포트폴리오_통합관리(), TestS종목별_포트폴리오_통합관리().")
	fmt.Println("TODO : I포트폴리오 구현체.")
	fmt.Println("TODO : I종목별_포트폴리오_통합관리. G단가() 기준.")
	fmt.Println("TODO : C포트폴리오내역구성원, C포트폴리오내역구성원 수정.")
	fmt.Println("       0203 C포트폴리오내역구성원 : G현재단가(), G매입금액().")
	fmt.Println("       0203 C포트폴리오내역 : GSharpe비율(), G연평균수익률().")
	fmt.Println("TODO : I위험관리 구현체. VAR, 절대 수치, 자본 대비 비율.")
	fmt.Println("TODO : I전략그룹과 I전략을 하나로 통일.")
	fmt.Println("TODO : 0900_테스트용_샘플데이터_test.go")
	fmt.Println("		테스트의 기반이 되는 샘플데이터에 테스트가 없으면 테스트 자체가 취약해 짐.")
	fmt.Println("TODO : go test ghts/kr/common 병목지점 해결.")
	fmt.Println("")
	fmt.Println("TO_RESEARCH : Concurrency 문제의 원인은 mutable 데이터에 동시 접근임.")
	fmt.Println("			   이에 대한 근본적인 해결책을 연구해 볼 것.")
	fmt.Println("			1. Erlang : 완벽한 메모리 공간 분리. immutable만 사용.")
	fmt.Println("						이제서야 함수형 프로그래밍을 배운다는 게 가능한가?")
	fmt.Println("						개발 효율도 확연히 떨어질 것이다.")
	fmt.Println("			2. Rust : 아직 미완성이다. 불확실성이 높고, 컴파일 속도가 심히 의심됨")
	fmt.Println("			3. Go : 근본적으로 mutable 데이터에 대한 동시접근을 차단할 수 없다.")
	fmt.Println("					코드 규칙과 race detector가 해결책이 될까?")
	fmt.Println("			- immutable 및 자동복사 되는 데이터만 주고 받기.")
	fmt.Println("			- 함수나 메소드는 외부 mutable 데이터를 직접 건드리지 않기.")
	fmt.Println("			- 모든 mutable 데이터의 변경은 중앙 저장소를 통해서만 하기.")
	fmt.Println("			  (메모리DB 사용도 고려. STM과 비슷하다.)")
	fmt.Println("			- 운영체제 프로세스를 이용해서 메모리 공간을 근본적으로 분리.")
	fmt.Println("			  : mutable한 데이터에 대한 접근은 메시지 큐 혹은 RPC를 이용.")
	fmt.Println("			    메모리 낭비 및 태스트 스위칭으로 인한 성능 저하가 있겠지만,")
	fmt.Println("			    mutable 데이터 공유로 인한 골치아픈 문제를 겪는 것보다 ")
	fmt.Println("			    CPU코어와 메모리를 추가하는 것이 훨씬 간단하고, 확실하고,")
	fmt.Println("			    장기적으로 비용도 적게 들 것이다.")
	fmt.Println("				개별 goroutine이 외부 mutable 데이터에 동시 접근하는 것을")
	fmt.Println("				자동으로 검사할 수 있으면 좋을 텐데. Rust가 답이지만 미완성.")
	fmt.Println("")
	fmt.Println("1. 위험관리 : VAR 방식의 잠재적 손실폭 제한.")
	fmt.Println("			잠재적 최대손실이 현재 자본금의 비율을 추가 매수를 중단.")
	fmt.Println("2. 위험관리 : VAR을 벗어나는 경제위기 상황에 대비한 절대적 손실폭 제한.")
	fmt.Println("3. 위험관리 : 경제위기 상황에서 손실율이 감당할 수 있는 범위 이내인지")
	fmt.Println("				확인하는 테스트 케이스 작성할 것.")
	fmt.Println("4. 종목 선정 : 우량주 위주로 일평균 거래량, 주당 가격을 고려하여 선정.")
	fmt.Println("				4 묶음으로 나누어서, 1/4만 전략 개발용으로 사용.")
	fmt.Println("				검증용으로 1/4씩만 사용하여 3중 검증을 거칠 것.")
	fmt.Println("				일단 초기 연구 대상 종목부터 선정할 것.")
	fmt.Println("")
	fmt.Println("PLAN : 만약 사용자 UI를 만들게 된다면 HTML5 기반으로 한다.")
	fmt.Println("		a. GopherJS : Javascript에 적응하는 어려움을 덜어줄 가능성이 있음.")
	fmt.Println("		b. AngularJS : DOM을 직접 조작해야 하는 어려움을 덜어줄 가능성이 있음.")
	fmt.Println("			 			GopherJS용 바인딩도 존재함.")
}
