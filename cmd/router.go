package cmd

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dashotv/groak/database"
	"github.com/dashotv/groak/scraper"
)

func Router(p *scraper.Processor, db *database.Database, settings *database.Settings) {
	go func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, settings)
		})
		e.POST("/process", func(c echo.Context) error {
			go func() {
				p.Process()
			}()
			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.POST("/scrapers", func(c echo.Context) error {
			n := struct {
				Name string `json:"name"`
			}{}

			err := c.Bind(&n)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			if n.Name == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, errors.New("name is required"))
			}

			settings.AddScraper(n.Name)

			err = db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.DELETE("/scrapers/:name", func(c echo.Context) error {
			name := c.Param("name")
			settings.RemoveScraper(name)

			err := db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.POST("/downloaders", func(c echo.Context) error {
			n := struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			}{}

			err := c.Bind(&n)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			if n.Name == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, errors.New("name is required"))
			}
			if n.URL == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, errors.New("url is required"))
			}

			settings.AddDownloader(n.Name, n.URL)

			err = db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.DELETE("/downloaders/:name", func(c echo.Context) error {
			name := c.Param("name")
			settings.RemoveDownloader(name)

			err := db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.POST("/pages", func(c echo.Context) error {
			n := &database.Page{}

			err := c.Bind(n)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			if n.Name == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, errors.New("name is required"))
			}
			if n.URL == "" {
				return echo.NewHTTPError(http.StatusInternalServerError, errors.New("url is required"))
			}
			if n.Scraper == "" {
				n.Scraper = "myanime"
			}
			if n.Downloader == "" {
				n.Downloader = "metube"
			}

			settings.AddPage(&database.Page{
				Name:       n.Name,
				URL:        n.URL,
				Scraper:    n.Scraper,
				Downloader: n.Downloader,
			})

			err = db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			if err := p.ProcessSingle(n.Name); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.DELETE("/pages/:name", func(c echo.Context) error {
			name := c.Param("name")
			settings.RemovePage(name)

			err := db.SaveSettings(settings)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		g := e.Group("/buckets")
		g.GET("/", func(c echo.Context) error {
			buckets, err := db.ListBuckets()
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return c.JSON(http.StatusOK, buckets)
		})
		g.GET("/:bucket", func(c echo.Context) error {
			bucket := c.Param("bucket")
			data := []*kv{}
			err := eachItem(bucket, func(s string) error {
				val, err := db.Get(bucket, s)
				if err != nil {
					return err
				}

				data = append(data, &kv{s, val})
				return nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return c.JSON(http.StatusOK, data)
		})
		g.DELETE("/:bucket", func(c echo.Context) error {
			bucket := c.Param("bucket")
			err := db.DeleteBucket(bucket)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		g.DELETE("/:bucket/:index", func(c echo.Context) error {
			bucket := c.Param("bucket")
			index := c.Param("index")

			i, err := strconv.Atoi(index)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			data := []*kv{}
			err = eachItem(bucket, func(s string) error {
				val, err := db.Get(bucket, s)
				if err != nil {
					return err
				}

				data = append(data, &kv{s, val})
				return nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}

			kv := data[i]
			err = db.Delete(bucket, kv.Key)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
		e.Logger.Fatal(e.Start(":9003"))
	}()
}
