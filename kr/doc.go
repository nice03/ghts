/*******************************************************************************

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

********************************************************************************

소프트웨어 개발자(들)의 법적 책임 면제.

'ghts', GHTS, 'GH Trading System', 'GH 매매 시스템', 'GH 거래 시스템',
'GH 자동 매매 시스템', 'GH 자동 거래 시스템' 및  일부분 혹은 전부를 띄워쓰기 없이 붙여쓴 것은
모두 같은 소스코드 및 그 패키지를 지칭함.

GHTS 소스코드는 "as is" 현재 상태 그대로 제공됨.

GHTS의 개발자(들) 혹은 저자(들)은
명확하게 표현되었거나, 묵시적으로 암시되었거나에 상관없이
모든 보증에 대한 책임을 법에 의해서 허용되는 최대한도까지 법적인 책임을 면제 받음.

다음 사항들은 모두 사용자의 책임이며,
GHTS의 개발자(들)은 다음 사항에 대해서 어떠한 책임도 지지 않는다는 전제 하에,
GHTS의 다운로드, 컴파일, 설치, 수정, 배포 및 사용을 허가함.
- 데이터 백업을 하지 않아서 발생한 모든 손상, 손실, 피해.
- GHTS를 운용하면서 발생한 관련 하드웨어나 소프트웨어에 대한 손상, 손실, 피해.
- 기타 GHTS를 사용함으로 생긴 모든 금전적, 물질적, 정신적 손상, 손실, 피해

다음 경로를 통해서 얻은 구두 및 문서 형태의 조언, 정보에 대해서 어떠한 보증도 성립하지 않음.
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
등을 포함한 그 어떠한 보증도 하지 않음.

GHTS를 사용하다가 어떠한 이유에서건 금전적, 정신적 손실이 발생하더라도
GHTS개발자(들)은 책임지지 않음.
설사, 그 이유가 이 프로그램의 오류, 잘못된 설계등으로 인해 발생한 것이어도 책임지지 않음.
심지어, 그런 오류, 잘못된 설계가 개발자가 고의로 한 것이더라도 책임지지 않음.

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
***************************************************************************/
package kr
