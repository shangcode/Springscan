package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var letters = []string{"/%20/swagger-ui.html", "/actuator", "/actuator/auditevents", "/actuator/beans", "/actuator/conditions", "/actuator/configprops", "/actuator/env",
	"/actuator/health", "/actuator/heapdump", "/actuator/httptrace", "/actuator/hystrix.stream", "/actuator/info", "/actuator/jolokia", "/actuator/logfile", "/actuator/loggers",
	"/actuator/mappings", "/actuator/metrics", "/actuator/scheduledtasks", "/actuator/swagger-ui.html", "/actuator/threaddump", "/actuator/trace", "/api.html", "/api/index.html",
	"/api/swagger-ui.html", "/api/v2/api-docs", "/api-docs", "/auditevents/", "/autoconfig", "/beans", "/caches", "/cloudfoundryapplication", "/conditions", "/configprops",
	"/distv2/index.html", "/docs", "/druid/index.html", "/druid/login.html", "/druid/websession.html", "/dubbo-provider/distv2/index.html", "/dump", "/entity/all", "/env",
	"/env/(name)", "/eureka", "/flyway", "/gateway/actuator", "/gateway/actuator/auditevents", "/gateway/actuator/beans", "/gateway/actuator/conditions", "/gateway/actuator/configprops",
	"/gateway/actuator/env", "/gateway/actuator/health", "/gateway/actuator/heapdump", "/gateway/actuator/httptrace", "/gateway/actuator/hystrix.stream", "/gateway/actuator/info",
	"/gateway/actuator/jolokia", "/gateway/actuator/logfile", "/gateway/actuator/loggers", "/gateway/actuator/mappings", "/gateway/actuator/metrics", "/gateway/actuator/scheduledtasks",
	"/gateway/actuator/swagger-ui.html", "/gateway/actuator/threaddump", "/gateway/actuator/trace", "/health", "/heapdump", "/heapdump.json", "/httptrace", "/hystrix", "/hystrix.stream",
	"/info", "/intergrationgraph", "/jolokia", "/jolokia/list", "/liquibase", "/logfile", "/loggers", "/mappings", "/metrics", "/monitor", "/prometheus", "/refresh", "/scheduledtasks",
	"/sessions", "/shutdown", "/spring-security-oauth-resource/swagger-ui.html", "/spring-security-rest/api/swagger-ui.html", "/static/swagger.json", "/sw/swagger-ui.html", "/swagger",
	"/swagger/codes", "/swagger/index.html", "/swagger/static/index.html", "/swagger/swagger-ui.html", "/swagger-dubbo/api-docs", "/swagger-ui", "/swagger-ui.html", "/swagger-ui/html",
	"/swagger-ui/index.html", "/system/druid/index.html", "/template/swagger-ui.html", "/threaddump", "/trace", "/user/swagger-ui.html", "/v1.1/swagger-ui.html", "/v1.2/swagger-ui.html", "/v1.3/swagger-ui.html", "/users",
	"/v1.4/swagger-ui.html", "/v1.5/swagger-ui.html", "/v1.6/swagger-ui.html", "/v1.7/swagger-ui.html", "/v1.8/swagger-ui.html", "/v1.9/swagger-ui.html", "/v2.0/swagger-ui.html", "/v2.1/swagger-ui.html", "/v2.2/swagger-ui.html", "/v2.3/swagger-ui.html", "/v2/swagger.json", "/webpage/system/druid/index.html",
}

func worker(u string, dirct chan string, wg *sync.WaitGroup) {

	for d := range dirct {
		address := u + d
		req, _ := http.NewRequest("GET", address, nil)
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		resp, err := (&http.Client{}).Do(req)
		if err != nil {
			fmt.Printf(Red("[-] " + address + " InternetError\n"))
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != 404 {
			fmt.Printf(Green("[+] %s status:%d len:%d\n"), address, resp.StatusCode, len(body))
		}

		resp.Body.Close()
		wg.Done()
	}
}

func Envscan(u string, thread int) {

	fmt.Println(Yellow("\nSpring敏感端点路径扫描开始："))
	dirct := make(chan string, thread)
	var wg sync.WaitGroup

	for i := 0; i < cap(dirct); i++ {
		go worker(u, dirct, &wg)

	}

	for _, v := range letters {
		wg.Add(1)
		dirct <- v

	}
	wg.Wait()
	close(dirct)
}
