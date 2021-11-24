package controllers

import (
	"fmt"
	"strings"

	// "log"
	// "github.com/fuadsuleyman/go-couriers/database"
	// "github.com/fuadsuleyman/go-couriers/helper"
	// "github.com/fuadsuleyman/go-couriers/models"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"bytes"
	"io/ioutil"
	"github.com/google/uuid"
)

// I don't need this api
func MyProxy(c *fiber.Ctx) error {

	fmt.Println("port: ", c.Port())
	fmt.Println("BaseURL: ", c.BaseURL())
	fmt.Println("Hostname: ", c.Hostname())
	fmt.Println("Path: ", c.Path())

	origin, _ := url.Parse("http://192.168.31.74:8005/api/v1.0/proxy")
	fmt.Println("origin:", origin)
	baseUrl := c.BaseURL()
	myLen := len(baseUrl) - 4
	port := baseUrl[myLen:]
	path := c.Path()
	
	fmt.Println("custom port: ", port)

	parts := strings.Split(strings.TrimPrefix(path, "/"), "/")

	fmt.Println("parts", parts)

	targetHost := fmt.Sprintf("svc-%s", parts[1])
	targetNamespace := fmt.Sprintf("svc-%s", parts[2])

	targetAddr := fmt.Sprintf(
		"http://%s.%s:%d/api/%s",
		targetHost, targetNamespace, 10000, strings.Join(parts[3:], "/"),
	)

	targetUrl, err := url.Parse(targetAddr)

	// "net/http/httputil" part
	p := httputil.NewSingleHostReverseProxy(targetUrl)

	fmt.Println("p from httputil:", *p)

	p.Director = func(request *http.Request) {
		request.Host = targetUrl.Host
		request.URL.Scheme = targetUrl.Scheme
		request.URL.Host = targetUrl.Host
		request.URL.Path = targetUrl.Path
	}
	p.ModifyResponse = func(response *http.Response) error {
		if response.StatusCode == http.StatusInternalServerError {
			u, s := readBody(response)
			logrus.Errorf("%s ,req %s ,with error %d, body:%s", u.String(), targetUrl, response.StatusCode, s)
			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(fmt.Sprintf("error %s", u.String()))))
		} else if response.StatusCode > 300 {
			_, s := readBody(response)
			logrus.Errorf("req %s ,with error %d, body:%s", targetUrl, response.StatusCode, s)
			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(s)))
		}
		return nil
	}



	fmt.Println("targetHost:", targetHost)
	fmt.Println("targetNamespace:", targetNamespace)
	fmt.Println("targetAddr:", targetAddr)
	fmt.Println("err:", err)
	fmt.Println("targetUrl:", targetUrl)
	fmt.Println("p after mofications:", *p)




	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "look at terminal",
	})
}

func readBody(response *http.Response) (uuid.UUID, string) {
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	u := uuid.New()
	var s string
	if len(all) > 0 {
		s = string(all)
	}
	return u, s
}