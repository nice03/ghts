exp
===

실험적인 소스코드.

개발 초기 단계이며 앞으로 많은 변경이 예상되므로 아직 실제로 사용하는 것은 권장되지 않음.

Experimental

All of the code in this (including subdirectories) are experimental and will change a lot.

It is NOT RECOMMENDED for any real use.

====
ghts
====

********************************************************************

주의 : 지금 개발 초기 단계입니다. 

		개발하는 동안 많은 변동이 있을 예정입니다.
	 
		API가 안정화 된 이후에야 사용하는 것을 권장합니다.

Warning : It is very early stage of development.

		And there will be a lot of API change during development.
		
		It is NOT suggested to use for any real use at this stage.

********************************************************************

GHTS : GH Trading System. GH 매매 시스템.

자동으로 주식를 거래하는 프로그램을 만들어 보고 싶어서 진행 중인 개인 프로젝트임.
'프로그램 매매'라는 단어를 많이 들어봤을 텐데, 이 프로젝트의 목표가 
그런 프로그램을 개발하는 것임.

프로그램 매매에 관심있는 사람에게 유용할 수 있을 것 같아서 소스코드를 공개함.
소스코드는 LGPL V3 라이센스를 따름.

LGPL V3 라이센스를 간단히 설명하면,
- 마음대로 고쳐쓰던, 
- 자기만 쓰던, 
- 다른 사람에게 공짜로 주던, 
- 다른 사람에게 돈 받고 팔던 
모두 다 자유임.

단,
- GHTS 자체에 문제점을 발견하고 수정하거나 개선한 경우에는 변경된 GHTS 소스코드를 공개해야 함.
- 외부에서 GHTS의 기능을 불러다 사용한 소스 코드는 원하지 않는다면 공개할 필요가 없음.
- 외부에 공개하고 싶지 않은 소스코드는 별도의 폴더(내지 디렉토리)에 위치시켜서 GHTS와 구분 요망.
- 마지막으로, GHTS에 대해서 특허소송을 제기하는 사람 혹은 법인은 GHTS를 사용할 권리를 잃게 됨.

LGPL V3 라이센스를 채택한 이유는 기본적인 기능은 소스코드를 공개하여서,
사용 상의 문제점도 함께 찾고, 고치고, 개선해 나가되,
매매 전략, 위험관리 전략등 각자 자신만의 노하우 및 경쟁우위 요소는 외부에 유출해야 되는
상황이 생기지 않도록 하고자 하는 의도임.

아무리 컴퓨터로 주식매매를 하더라도 가장 어려운 것은 컴퓨터 프로그래밍 기술이 아니라,
위험관리 원칙을 지키고, 항상 평정심을 유지하는 것임.
과도한 욕심을 부리거나, 원칙을 어기면 아무리 컴퓨터가 도와줘도 별 도움이 안 됨.
어차피, 컴퓨터는 사람이 시키는 대로 할 뿐임.

그 외, 추가적인 정보는 doc 디렉토리의 파일을 참고할 것.

**************************************************************************

명명규칙

처음에 오는 첫 1~2글자로 형식을  구분함. 
이러한 명명방식을 헝가리안 명명법이라고 하며, 예전에 많이 사용되었으나 요즘은 별로 사용되지 않음.
이 프로그램에서 굳이 예전 헝가리안 명명법을 사용하는 이유는 

1. 고정값 형식과 변경가능값 형식을 간편하게 구분할 수 있음.
2. 인터페이스와 구조체를 간편하게 구분할 수 있음.

고정값 형식은 여러 모듈이 동시에 작동하면서 자료를 공유할 때 발생할 수 있는 온갖 복잡한 문제를
원천적으로 제거하는 가장 확실한 방법임.

Scala언어에서는 컴파일러 차원에서 고정값과 변경가능값을 val, var로 구분해 주지만,
Go언어는 컴파일 속도를 높이기 위해서 그런 기능을 넣지 않았음. (대신, 컴파일 속도는 환상적으로 빠름.)
값을 상수형으로 만들려면 내부값을 변경하는 메소드가 없는 구조체로 값을 감싸면 됨.
약간 귀찮지만, Go언어의 컴파일 속도를 높이기 위해서 꼭 필요한 기능만 넣다보니 생긴 결과이며,
Go언어의 환상적인 컴파일 속도와 그로 인한 높은 개발 생산성를 경험해 보면 이 모든 게 용서됨.

- I : 일반 인터페이스. Interface의 줄임말.
- C : 고정값 자료형 인터페이스. Constant의 줄임말.
- V : 변경가능 자료형 인터페이스. Variable의 줄임말.


- F : 일반적인 함수. function의 줄임말.
- N : 생성자 함수. New의 줄임말.
- NC : 고정값 자료형 생성자 함수. 'New Constant'의 줄임말.
- NV : 변경가능 자료형 생성자. 'New Variable'의 줄임말.


- S : 구조체. Structure의 줄임말. (외부에 공개하지 않는 때는 첫 글자를 소문자로 함.)
- SC : 고정값 자료형 구조체. 'Structure Constant'의 줄임말.
- SV : 변경가능 자료형 구조체. 'Structure Variable'의 줄임말.


- G : 구조체의 값을 변경하지 않고 읽기만 하는 메소드. 
       Get의 줄임말. 대개의 경우 반환값이 있음.
- S : 구조체의 값을 변경 및 설정하는 메소드.
     	Set의 줄임말. 반환값이 없는 경우가 많음.

Go언어는 컴파일러가 생성자 기능을 지원하지 않음.
Go언어에서 구조체를 생성하는 new(<형식>)만으로는 적절한 초기화가 어려움.

이를 해결하기 위해서 다음 패턴을 사용함.
1. 구조체는 외부에 공개하지 않음. (첫 글자는 소문자 s로 시작함. 상수형은 sC, 변수형은 sV.)
2. 생성자 함수를 통해서만 구조체를 생성할 수 있음. 이 때, 생성자 함수에서 적절한 초기화가 가능함.
3. 생성한 구조체는 외부에 공개된 인터페이스를 통해서만 사용함.
   (구조체는 생성된 후에도 공개되지 않은 관계로 직접 사용할 수 없음.)

Java 클래스의 생성자 기능을 구조체, 생성자 함수, 인터페이스 3가지를 조합해서 구현했음.

**************************************************************************

간단한 기본 구조.

1. 가격정보를 받아서 배포하는 모듈
2. 가격정보를 토대로 확률을 계산하는 모듈. 
	(금리 + 알파의 수익을 지향한다는 의미에서 흔히들 알파 모듈이라고 함.)
3. 위험관리 및 자금관리 모듈
4. 거래 주문을 증권사에 보내는 모듈

각 모듈은 goroutine을 이용해서 동시에 실행됨.

즉, 가격정보를 받아서 확률을 계산하고, 위험관리 룰에 위배되는 지 확인한 후 주문을 내는 동작이
순차적으로 진행되는 것이 아니라, 각 모듈이 동시에 독자적으로 실행됨.
순차적으로 수행된다면 알파 모듈이 계산을 수행하는 동안에 들어오는 새로운 가격정보는 못 받게 될 수 있음.
그러나, 알파 모듈과 가격정보 수집 모듈이 독자적으로 동시에 작동하면 그런 문제가 없음.

예전에는 이러한 동시처리를 스레드나 이벤트를 이용해서 구현하면서 온갖 복잡한 문제가 생겼었지만,
Go언어는 자체적으로 간편하게 동시처리를 수행하는 goroutine 기능을 내장하고 있어서 간단함.

**************************************************************************

Go 언어 맛보기.
http://go-tour-kr.appspot.com/

Go언어 관련 번역문서를 모아놓은 곳.
https://code.google.com/p/golang-korea/


Go언어 소개.

1. 누가? : 처음에는 Rob Pike, Ken Thomson이 시작했음.
			현대 컴퓨터 및 인터넷의 주요 기반인 UNIX와 UTF8를 만든 인물이 Ken Thompson이며,
			Rob Pike는 Ken Thompson과 함께 다양한 시스템 소프트웨어를 만든 인물임.
			
2. 언제? : 2000년대 후반.

3. 어디서? : 미국 구글(Google Inc.)

4. 무엇을? : 개발 생산성도 높고, 하드웨어 효율성도 높은 컴퓨터 프로그래밍 언어 Go. 
			흔히, golang으로 칭함. (Go Language의 줄임말.)
			
5. 왜? : C++은 온갖 기능이 너무 많아서 
		- 언어 문법이 복잡하고,
		- 컴파일 속도가 느리고, (컴파일러가 해야할 일이 너무 많음.)
		- 멀티코어 CPU를 효율적을 사용하기 위한 동시처리 및 병렬처리 기능이 없음.
		구글의 대규모 환경에서 개발 생산성도 높고, 멀티코어 CPU를 잘 활용하기 하기 위한
		새로운 프로그래밍 언어가 필요하다고 생각했음.
		Ken Thomson도 여기에 동의하면서 전설적인 인물 2명이 일을 저지르기 시작함.
		목표는 필수적인 기능만을 가진 간단하고 빠르고 효율적인 프로그래밍 언어임.
		(그래서, 가끔 기능이 부족하다거나 불편하다고 느껴질 때도 있음.)

6. 어떻게? : C언어를 기반으로 개발자의 잦은 실수를 유발하던 문제점을 수정한 후,
			간단한 객체지향, (상속 대신 임베딩)
			자동 메모리 관리, (가비지 컬렉션)
			간편한 동시처리 및 병렬처리 기능을 추가하고, (goroutine)
			꼭 필요한 핵심 기능만 넣고, 최대한 기능을 줄여서 빠른 컴파일 속도를 추구함.

인터넷 최대 기업인 Google이 최고의 인재를 모아서 개발한 언어가 Go언어임.
현재는 주로 대규모 인터넷 기업에서 사용됨.

Go언어의 가장 큰 장점은 
1. 단순한 문법.
2. 빠른 컴파일 속도.
3. 간편한 동시성 처리.
4. 자동 메모리 관리.
5. 빠른 실행 성능.

1~4는 개발 생산성에 관련된 항목이고, 
5는 개발 후 실행속도에 관련된 항목.

애초에 필수적인 기능만 포함하고 대신 빠르고 효율적인 언어를 지향했기 때문에,
Scala처럼 기능이 풍부하고 확장성이 높은 언어를 좋아하는 사람에게는 안 맞을 수도 있음.

1. 가장 큰 단점은 IDE(통합 개발환경)이 없다는 것임.
	Java의 Eclipse, C#의 Visual Studio에 익숙한 사람에게는 아주 불편하게 느껴짐.
	그런데, 실제로 써보면 별로 안 불편함.
	Linux 개발자인 Linus Torvalds도 VI라고 하는 텍스트 에디터로 Linux를 개발했음.
	즉, IDE를 사용하지 않고 그토록 거대한 소프트웨어를 개발했음.
	 
2. 한글로 된 문서가 부족함.
	Java, C#, PHP, Python, Ruby등 주류 언어에 비하면 한글 기술 문서가 부족한 편임.
	영어로 된 각종 팁이나 문서는 인터넷에 넘쳐남
	
3. Java의 final, Scala의 val, Ruby의 freeze처럼 
	값을 변경할 수 없도록 고정하는 기능이 없음.
	더 이상 값을 고정시킬 수 있으면, 동시처리에서 자료공유로 인한 온갖 문제를 
	원천적으로 제거할 수 있는 데 언어 자체에서 그런 기능을 지원하지는 않음.
	대신, 값을 구조체로 감싸고 값을 변경할 수 있는 공개된 메소드가 없으면,
	값을 고정시키는 효과를 낼 수 있음.
	그 외 제너릭도 없음. 형변환 하기 귀찮을 때 많음.
	
4. 변수 선언의 순서가 반대라서 처음에 어색함. (int a가 아니고 a int임. 델파이, 파스칼 형식.)
	이것은 컴파일 속도를 빠르게 하기 위한 목적임.
	요즘 새로 나오는 언어들이 많이 채택하는 방식임.

5. if문으로 에러처리를 하다보니 소스코드가 지저분해 짐
	도우미 함수로 어느 정도 해결이 가능하며, 
	if문으로 에러값을 확인하는 과정을 아예 생략하고, 
	함수 전체 차원에 defer recover()를 사용하면 
	Java의 try, catch, finally와 비슷한 효과를 냄.

6. 객체(데이터와 그 데이터를 처리하는 메소드를 함께 가지고 있는 자료형)를 다루는 방식이 
	Java, C#, C++등 일반적인 OOP언어와 다름.
	상속이 아니라 인터페이스 조합, 구조체 조합으로 비슷한 효과를 냄.

7. 메소드 오버로딩이 없음
	각각의 입력 파라메터 형식에 대해서 각각 이름이 다른 별도의 함수를 만들거나,
	하나의 함수에서 타입 switch문으로 입력값 형식을 판별하고 각각 다르게 처리해 주면 됨.
	본인은 타입 switch문 방식을 선호함.
	
	Java나 C#에서는 상속받아서 만든 새로운 타입에 맞는 오버로딩 메소드를 만들지 않아도,
	컴파일 에러가 안 나는 경우가 많아서 깜빡 잊고 그냥 넘어가는 경우가 있지만,
	Go언어의 방식은 타입 switch문의 default문에서 에러 발생하니까, 
	잊어버리고 넘어가는 일이 적어서 오히려 안전한 면도 있음.

이 모든 단점에도 불구하고, 모든 게 다 용서되는 Go언어의 특징은 
환상적으로 빠른 컴파일 속도에 있다.

인간은 대부분 생각의 흐름이 5초 이상 끊기면 집중력이 떨어지기 시작하며,
15초 이상 끊기면 대부분의 경우 딴 생각을 하기 시작함.
Go 언어로 개발하면 컴파일 때문에 집중력이 떨어지는 일은 없음.
따라서, 개발 생산성이 높음.

요약하면, Go언어는 개발속도도 빠르고, 동시처리와 병렬처리도 간편하고, 실행속도도 빠름.

Go 언어 맛보기.
http://go-tour-kr.appspot.com/

Go언어 관련 번역문서를 모아놓은 곳.
https://code.google.com/p/golang-korea/

**************************************************************************

소프트웨어 개발자(들)의 법적 책임 면제.

'ghts', GHTS, 'GH Trading System', 'GH 매매 시스템', 'GH 거래 시스템', 
'GH 자동 매매 시스템', 'GH 자동 거래 시스템' 및  일부분 혹은 전부를 띄워쓰기 없이 붙여쓴 것은
모두 같은 소스코드 및 그 패키지를 지칭합니다. 

GHTS 소스코드는 "as is" 현재 상태 그대로 제공됩니다.

GHTS의 개발자(들) 혹은 저자(들)은 
명확하게 표현되었거나, 묵시적으로 암시되었거나에 상관없이
모든 보증에 대한 책임을 법에 의해서 허용되는 최대한도까지 면제 받습니다.

다음 사항들은 모두 사용자의 책임이며, 
GHTS의 개발자(들)은 다음 사항에 대해서 어떠한 책임도 지지 않는다는 전제 하에,
GHTS의 다운로드, 컴파일, 설치, 수정, 배포 및 사용을 허가합니다.
- 데이터 백업을 하지 않아서 발생한 모든 손상, 손실, 피해.
- GHTS를 운용하면서 발생한 관련 하드웨어나 소프트웨어에 대한 손상, 손실, 피해.
- 기타 GHTS를 사용함으로 생긴 모든 금전적, 물질적, 정신적 손상, 손실, 피해

다음 경로를 통해서 얻은 구두 및 문서 형태의 조언, 정보에 대해서 어떠한 보증도 성립하지 않습니다.
- GHTS의 개발자(들)
- GHTS의 개발자(들)이 운영하는 웹사이트, 메일, 블로그, SNS, 메신저, 채팅, 기타 인터넷 매체. 
- GHTS 소스코드
- GHTS 관련 문서

위에서 언급한 사항 이외에도 명백히 표현되었던지, 묵시적으로 암시되었던 상관없이 
GHTS의 개발자(들)은 
- 금전적 이득을 가져다 줄 것이라던 지, 
- 특정 목적에 적합하다던지,
- 사용자의 요구사항과 기대를 만족시킨다던지,
- 버그, 에러, 바이러스 및 기타 결함이 없다던지,
- 소스코드가 생성한 결과물, 출력물, 데이터가 정확하던가, 
    최신의 것이라던가, 완전하다거나, 신뢰할 수 있다던지,
- 소스코드가 다른 소프트웨어와 호환된다던지,
- 에러가 수정될 것이라던 지,
등을 포함한 그 어떠한 보증도 하지 않습니다.

GHTS를 사용하다가 어떠한 이유에서건 금전적, 정신적 손실이 발생하더라도
GHTS개발자(들)은 책임지지 않습니다.
설사, 그 이유가 이 프로그램의 오류, 잘못된 설계등으로 인해 발생한 것이어도 책임지지 않습니다.
심지어, 그런 오류, 잘못된 설계가 개발자가 고의로 한 것이더라도 책임지지 않습니다.

**************************************************************************

Disclaimer of Warranties.

All of the terms, 'ghts', 'GHTS', 'GH Trading System', means 
same source code, or package of source code.

Source code available in GHTS are provided "as is" 
  without warranty of any kind, 
  either expressed or implied 
  and such software is to be used at your own risk.

Authors(or developers) of GHTS disclaims to the fullest extent 
  authorized by law any and all other warranties, 
  whether express or implied, 
  including, without limitation, any implied warranties 
  of merchantability or fitness for a particular purpose. 
  
The use of GHTS is done at your own discretion 
  and risk and with agreement 
  that you will be solely responsible for any damage or loss 
  to you and your computer hardware and software.

You are solely responsible for adequate protection and backup of the data 
  and equipment used in any of the software related to GHTS. 
  and we will not be liable for any damages 
  that you may suffer in connection with downloading, installing, using, 
  modifying or distributing GHTS. 

No advice or information, whether oral or written, 
obtained by you from authors(or developers) of GHTS 
or from websites, 
or from source code,
or related documents
shall create any warranty for the software.
  
Without limitation of the foregoing, 
authors(or developers) of GHTS expressly does not warrant that:

1. the software will meet your requirements or expectations.
2. the software or the software content will be free of bugs, errors, 
     viruses or other defects.
3. any results, output, or data provided through or generated 
     by the software will be accurate, up-to-date, complete or reliable.
4. the software will be compatible with third party software.
5. any errors in the software will be corrected.

**************************************************************************