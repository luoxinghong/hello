package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET(
		"/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		},
	)
	fmt.Println("测试合并lxh分支到dev")
	e.Logger.Fatal(e.Start(":6666"))
}
