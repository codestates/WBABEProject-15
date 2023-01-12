package router

//router.go : api 전체 인입에 대한 관리 및 구성을 담당하는 파일
import (
	"fmt"
	ctl "lecture/oos/controller"
	"lecture/oos/docs" //swagger에 의해 자동 생성된 package
	"lecture/oos/logger"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당
	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {
	e := gin.New() // gin선언

	e.Use(logger.GinLogger())       // gin 내부 log, logger 미들웨어 사용 선언
	e.Use(logger.GinRecovery(true)) //gin 내부 recover, recovery 미들웨어 사용 - 패닉복구
	e.Use(CORS())                   // crossdomain 미들웨어 사용

	logger.Info("start server")
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost" //swagger 정보 등록

	customer := e.Group("/customer", liteAuth())
	{
		fmt.Println(customer)
		customer.GET("/getMenu/:sortOption", p.ct.GetMenu)   //메뉴 리스트 출력 조회
		customer.GET("/getReview/:menuName", p.ct.GetReview) //메뉴별 평점 및 리뷰 조회
		customer.POST("/writeReview", p.ct.WriteReview)      //메뉴별 평점 작성
		customer.POST("orderMenu", p.ct.OrderMenu)           //메뉴 선택 후 주문
		customer.PUT("changeMenu", p.ct.ChangeMenu)          // 메뉴변경
		customer.PUT("addMenu", p.ct.AddMenu)                //메뉴 추가
		customer.GET("getOrderState", p.ct.GetAllOrderList)  //주문 내역(상태) 조회
	}

	seller := e.Group("/seller", liteAuth())
	{
		fmt.Println(seller)
		seller.PUT("/updateMenu", p.ct.UpdateMenu)             //메뉴 수정
		seller.POST("/register", p.ct.RegisterMenu)            //신규메뉴 등록
		seller.PUT("/updateOrderState", p.ct.UpdateOrderState) //주문내역 조회 및 상태 변경
		seller.DELETE("/delete/:menu", p.ct.DeleteMenu)        //메뉴 삭제
	}
	/* [코드리뷰]
	 * Group을 사용하여 API 성격에 따라 request를 관리하는 코드는 매우 좋은 코드입니다.
     * 일반적으로 현업에서도 이와 같은 코드를 자주 사용합니다. 훌륭합니다.
	 *
	 * 코드의 확장성을 고려하였을때, endpoint 관리를 함께 고려한 코드를 개발하는 것도 추천드립니다.
	 * 예를들어 /customer/orderMenu 를 호출하는 클라이언트(Web, App, etc..)들이 실시간으로 들어오고 있을 때,
	 * controller의 OrderMenu function을 변경해야 하는 상황이 발생한다면,
	 * /customer/orderMenu2 로 받아주는 경우가 있을 것이고(/customer/orderMenu 그대로 받아주면서)
	 * 처음부터 /customer/v1/OrderMenu 로 관리되며, 
	 * /customer/v2/orderMenu 리뉴얼 버전에 따라 version up을 시켜
	 * v01 방식의 클라이언트와, v02 방식의 클라이언트를 모두 받아줄 수 있는 
	 * 확장성 있는 코드를 구현해보시는 것을 추천드립니다.
	 */

	return e
}
