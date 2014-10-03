package common

import (
	"fmt"
)

func F나를_위한_문구() {
	fmt.Println("")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("	쉽고 간단하게, 테스트로 검증해 가면서 마음 편하게.")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("")
}
func F메모() {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("                  메            모")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("")
	fmt.Println("TODO : sC통화, sV통화에 String() 3자리마다 콤마 추가. 테스트 케이스 추가.")
	fmt.Println("TODO : C복합_상수형 이전 및 테스트 케이스 추가.")
	fmt.Println("TODO : 모든 V구조체에 sync.RWMutext 도입.")
	fmt.Println("TODO : I정밀수 내부 구현을 speter.decimal로 변경.")
	fmt.Println("		나눗셈의 나머지값 처리방식을 원하는 방식으로 지정 가능함.")
	fmt.Println("TODO : common.공용_자료형 및 테스트")
	fmt.Println("TODO : common.기타 자료형 및 펑션")
	fmt.Println("TODO : tools.* 바뀐 API에 맞게 수정.")
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
	fmt.Println("TODO : F값_일치 개선 계획.")
	fmt.Println("		기본 계획")
	fmt.Println("		- 두 값의 내부 필드값을 기록한 슬라이스 확보.")
	fmt.Println("		- 두 슬라이스에 대해서 비교.")
	fmt.Println("")
	fmt.Println("TODO : F공유해도_안전함() 개선 계획.")
	fmt.Println("		내부값을 기록하는 함수가 없어도 어느 정도는 가능할 듯 함.")
	fmt.Println("		모든 멤버 필드 트리의 끝부분(leaf)이 이미 Immutable하다고 ")
	fmt.Println("		알려진 상수형으로만 이루어져 있다면 비록 Immutable하지 않더라도,")
	fmt.Println("		race condition은 발생하지 않을 듯 함.")
	fmt.Println("		슬라이스나 맵은 문제의 소지가 있기는 한데, 그것도 구조체로 감싸버릴까?")
	fmt.Println("		물론, 내부값을 기록한 후 메소드를 실행하고 비교하는 것이 가장 확실하긴 함.")
	fmt.Println("		혹은, 모든 매개변수는 매개변수_맵을 통해서만 주고 받고,")
	fmt.Println("		매개변수 맵에서는 원소를 추가할 때 안전하다고 알려진 상수형만 받아들이는 ")
	fmt.Println("		방법도 있음.")
	fmt.Println("		이 방법이 가장 확실한 것 같다.")	
	fmt.Println("")
	fmt.Println("TODO : F공유해도_안전함() 에서 사용한 방법을 이용하면 F값_일치()에서")
	fmt.Println("		reflect.DeepEqual()에 의존하기 전에,")
	fmt.Println("		2개 값의 복사 기록본을 비교하면 될 듯함.")
	fmt.Println("		그러면, F값_일치()가 제대로 DeepEqual을 판별할 수 있을 것임")
	fmt.Println("")
	fmt.Println("		기술적으로 어렵지만 immutability를 확인하는 가장 확실한 방법.")
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
	fmt.Println("TODO : 자주 사용되는 혹은 자주 사용될 것으로 예상되는 상수형 값들을 ")
	fmt.Println("		미리 생성해서 캐쉬해 놓고 생성자에서는 포인터만 복사해서 주도록 하자.")
	fmt.Println("		참거짓은 이미 그런 식으로 되어 있으며,")
	fmt.Println("		숫자도 일정 범위까지는 값과 인덱스가 일치하도로 슬라이스에 저장하면 됨.")
	fmt.Println("		종목도 마찬가지.")
	fmt.Println("		이렇게 하면 상수형을 많이 사용해도 GC의 필요성이 줄어든다.")
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
	fmt.Println("		a. GopherJS, CoffeeScript : Javascript에 적응하는 어려움을 덜어줄 가능성이 있음.")
	fmt.Println("		b. AngularJS : DOM을 직접 조작해야 하는 어려움을 덜어줄 가능성이 있음.")
	fmt.Println("			 			GopherJS용 바인딩도 존재함.")
}
