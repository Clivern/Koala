// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/clivern/hippo"
	"github.com/gin-gonic/gin"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	host, _ := os.Hostname()
	release := fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)

	driver := hippo.NewRedisDriver(
		fmt.Sprintf("%s:%s", os.Getenv("KOALA_REDIS_HOST"), os.Getenv("KOALA_REDIS_PORT")),
		os.Getenv("KOALA_REDIS_PASSWORD"),
		0,
	)

	ok, err := driver.Connect()

	if !ok || err != nil {
		panic("Unable to connect to redis")
	}

	ok, err = driver.Ping()

	if !ok || err != nil {
		panic("Unable to connect to redis")
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stdout

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		LogRequest("/", host, release)
		c.JSON(http.StatusOK, gin.H{
			"TIME":     time.Now().Format("Mon Jan 2 15:04:05 2006"),
			"HOSTNAME": host,
			"RELEASE":  release,
		})
	})
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/_health", func(c *gin.Context) {
		LogRequest("/_health", host, release)
		hostHealth, _ := driver.Get(fmt.Sprintf("koala_host_health__%s", host))
		kindHealth, _ := driver.Get("koala_kind_health")

		if hostHealth == "DOWN" || kindHealth == "DOWN" {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"STATUS": "DOWN",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"STATUS": "UP",
		})
	})

	r.GET("/_change", func(c *gin.Context) {
		LogRequest("/_change", host, release)
		value, _ := driver.Get("koala_state")

		state := 1

		if value != "" {
			ivalue, _ := strconv.Atoi(value)
			state = ivalue + 1
		}

		driver.Set("koala_state", strconv.Itoa(state), 0)

		c.JSON(http.StatusOK, gin.H{
			"TIME":     time.Now().Format("Mon Jan 2 15:04:05 2006"),
			"HOSTNAME": host,
			"RELEASE":  release,
			"STATE":    state,
		})
	})

	r.GET("/_state", func(c *gin.Context) {
		LogRequest("/_state", host, release)
		value, _ := driver.Get("koala_state")

		state, _ := strconv.Atoi(value)

		c.JSON(http.StatusOK, gin.H{
			"TIME":     time.Now().Format("Mon Jan 2 15:04:05 2006"),
			"HOSTNAME": host,
			"RELEASE":  release,
			"STATE":    state,
		})
	})

	r.GET("/_hostup", func(c *gin.Context) {
		LogRequest("/_hostup", host, release)
		driver.Set(fmt.Sprintf("koala_host_health__%s", host), "UP", 0)

		c.JSON(http.StatusOK, gin.H{
			"STATUS": "DONE",
		})
	})

	r.GET("/_hostdown", func(c *gin.Context) {
		LogRequest("/_hostdown", host, release)
		driver.Set(fmt.Sprintf("koala_host_health__%s", host), "DOWN", 0)

		c.JSON(http.StatusOK, gin.H{
			"STATUS": "DONE",
		})
	})

	r.GET("/_kindup", func(c *gin.Context) {
		LogRequest("/_kindup", host, release)
		driver.Set("koala_kind_health", "UP", 0)

		c.JSON(http.StatusOK, gin.H{
			"STATUS": "DONE",
		})
	})

	r.GET("/_kinddown", func(c *gin.Context) {
		LogRequest("/_kinddown", host, release)
		driver.Set("koala_kind_health", "DOWN", 0)

		c.JSON(http.StatusOK, gin.H{
			"STATUS": "DONE",
		})
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("KOALA_PORT")))
}

// LogRequest logs some data
func LogRequest(path, host, release string) {
	log.Printf(`{"Path":"%s", "Time":"%s", "Hostname":"%s", "Version":"%s"}`,
		path,
		time.Now().Format("Mon Jan 2 15:04:05 2006"),
		host,
		release,
	)
}
