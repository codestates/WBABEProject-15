# WBABEProject-15


<About The Project>

#### <2023/01/06 2차 결과물 제출>
> 수정사항
  1. 메뉴등록 : 중복 로직 추가 구현
  2. 메뉴주문 : 주문 시 primary key로 작동될 수 있는 '주문번호' 속성 추가 구현하여 전체적인 주문 시스템 관리
  3. 주문상태 : const를 활용한 열거형 방식으로 주문 state 관리
  4. 평점리뷰 : min, max 설정 로직 구현
  5. 설정파일 : password 암호화
  

### 프로젝트 설명

> 언택트 시대에 급증하고 있는 온라인 주문 시스템을 이해, 경험하고 각 단계별 프로세스를 이해하여 구현함으로써 서비스 구축에 경험을 쌓고, golang의 이해를 돕는 프로젝트.
  1. 주문자/피주문자의 역할에서 필수적인 기능을 도출, 구현
  2. 주문자와 피주문자 입장에서 필요한 기능을 도출하여, 기능을 확장하고 주문 서비스를 원할하게 지원할수 있는 기능을 구현
  3. 주문자는 신뢰있는 주문과 배달까지를 원합니다. 피주문자는 주문내역을 관리하고 원할한 서비스를 제공

### 프로젝트 구성
  
![image](https://user-images.githubusercontent.com/65848709/209455977-7bea30a7-e193-4790-a20e-6d1112d96c8d.png)


   
### 기술 스택
  
 ![image](https://user-images.githubusercontent.com/65848709/209456025-dc86f5c1-191d-4f56-a520-2263ff1f4e96.png)


   
   
### 기능 설명
   
   
  - 메뉴 리스트 출력 조회 - 주문자
    **API |** 메뉴 리스트 조회 및 정렬(추천/평점/주문수/최신)
      - 각 카테고리별  sort 리스트 출력(ex. order by 추천, 평점, 재주문수, 최신)
      - 결과 5~10여개 임의 생성 출력, sorting 여부 확인

  - 메뉴별 평점 및 리뷰 조회 - 주문자
     **API |** 개별 메뉴별 평점 및 리뷰 보기
      - UI에서 메뉴 리스트에서 상기 리스트 출력에 따라 개별 메뉴를 선택했다고 가정
      - 해당 메뉴 선택시 메뉴에 따른 평점 및 리뷰 데이터 리턴

  - 메뉴별 평점 작성 - 주문자
     **API |** 과거 주문 내역 중, 평점 및 리뷰 작성
      - 해당 주문내역을 기준, 평점 정보, 리뷰 스트링을 입력받아 과거 주문내역 업데이트 저장
      - 성공 여부 리턴

  - 주문 - 주문자
     **API |** UI에서 메뉴 리스트에서 해당 메뉴 선택, 주문 요청 및 초기상태 저장
      - 주문정보를 입력받아 주문 저장(ex. 선택 메뉴 정보, 전화번호, 주소등 정보를 입력받아 DB 저장)
      - 주문 내역 초기상태 저장
      - 금일 주문 받은 일련번호-주문번호 리턴

  - 주문 변경 - 주문자
    **API |** 메뉴 변경 및 추가
      - 메뉴 추가시 상태조회 후 `배달중`일 경우 실패 알림
      - 성공 실패 알림, 실패시 신규주문으로 전환
      - 메뉴 변경시 상태가 `조리중`, `배달중`일 경우 확인
      - 성공 실패 알림

  - 주문 내역 조회 - 주문자
    **API |** 주문내역 조회
      - 현재 주문내역 리스트 및 상태 조회 - 하기 **주문 상태 조회**에서도 사용
      - ex. 접수중/조리중/배달중 etc
      - 없으면 null 리턴
      - 과거 주문내역 리스트 최신순으로 출력
      - 없으면 null 리턴

  - 주문 상태 조회 - 피주문자
    **API |** 현재 주문내역 리스트 조회
    **API |** 각 메뉴별 상태 변경
      - ex. 상태 : 접수중/접수취소/추가접수/접수-조리중/배달중/배달완료 등을 이용 상태 저장
      - 각 단계별 사업장에서 상태 업데이트
      - **접수중 → 접수** or **접수취소 → 조리중** or **추가주문 → 배달중**
      - 성공여부 리턴
        
  - 메뉴 신규 등록  - 피주문자
    **API |** 신규 메뉴 등록
     - 사업장에서 신규 메뉴 관련 정보를 등록하는 과정(ex. 메뉴 이름, 주문가능여부, 한정수량,  원산지, 가격, 맵기정도, etc)
     - 성공 여부를 리턴

  - 메뉴 수정 / 삭제 - 피주문자
    **API |** 기존 메뉴 수정/삭제
      - 사업장에서 기존의 메뉴 정보 변경기능(ex. 가격변경, 원산지 변경, soldout)
      - 메뉴 삭제시, 실제 데이터 백업이나 뷰플래그를 이용한 안보임 처리
      - 금일 추천 메뉴 설정 변경, 리스트 출력
      - 성공 여부를 리턴
    
  

  
 ### 데이터베이스 

  1. menu-list : 피주문자가 등록한 메뉴판

  2.order-info : 주문자들의 주문정보

  3. menu-review : 메뉴에 대한 리뷰
 
  
  
  
  ### 사용 방법

  ![스크린샷 2022-12-25 12-26-14](https://user-images.githubusercontent.com/65848709/209456086-32b8a03c-12d7-4182-98af-2008de4f31fc.png)

  > API Tester 에서 기능에 맞게 해당 함수 실행
  
  ex)
  
  ![스크린샷 2022-12-28 15-22-52](https://user-images.githubusercontent.com/65848709/209767943-9b3df2c9-61da-42c2-b9f6-895ac504a923.png)

  
  ![스크린샷 2022-12-25 12-28-39](https://user-images.githubusercontent.com/65848709/209456117-51f6a8db-8c72-4ae4-b968-943a8ec3fad7.png)
