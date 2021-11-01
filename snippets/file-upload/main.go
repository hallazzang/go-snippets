package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/upload", func(c echo.Context) error {
		// Client uploads a file as `file` field in the request.
		// In HTML, there may be an input tag like this:
		//
		// 	<input type="file" name="file" />
		//
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		f, err := file.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		// Now you can do whatever you want with `f`.
		// For example, you could save this file in the file system:
		//
		// 	f2, err := os.Create(file.Filename)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	defer f2.Close()
		//
		// 	if _, err := io.Copy(f2, f); err != nil {
		// 		return err
		// 	}
		//
		// Or you might want to upload it to AWS S3, etc.

		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
