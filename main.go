package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

// 여러 개의 상수를 묶어서 지정하고자 사용
var (
	framework string
	//선언과 동시에 할당할 수 있다
	// framework string = "fiber"
	port = flag.Int("p", 8888, "서버가 Listen할 port 번호를 입력해주세요.")
)

func init() {
	//flag.Parse()
	//fmt.Println(flag.Args())
	//if len(flag.Args()) != 1 {
	//	log.Fatal("하나의 인자를 전달해 framework를 정의해주세요. (e.g. http, echo, fiber)")
	//}
	//framework = flag.Arg(0)
	framework = "fiber"
}
func main() {
	switch framework {
	case "http":
		RunNewHttpServer()
	case "fiber":
		RunNewFiberServer()
	default:
		log.Fatal("지원하지 않는 framework 입니다.")
	}
}

func RunNewHttpServer() {
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server is listening %d", *port)
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		if _, err := w.Write([]byte("PingPong by net/http\n")); err != nil {
			log.Print(err)
		}
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Print(err)
	}
}

func RunNewFiberServer() {
	addr := fmt.Sprintf(":%d", *port)
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})
	log.Printf("Server is listening %d", *port)
	if err := app.Listen(addr); err != nil {
		log.Print(err)
	}
}

/*
*
복수 return 이 가능
*/
func sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}
