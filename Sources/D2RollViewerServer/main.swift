import PerfectHTTP
import PerfectHTTPServer

func handler(request: HTTPRequest, response: HTTPResponse) {
	response.setHeader(.contentType, value: "text/html")
	response.appendBody(string: "<html><title>Hello, world!</title><body>Hello, world!</body></html>")
	response.completed()
}

var routes = Routes()
routes.add(method: .get, uri: "/", handler: handler)
try HTTPServer.launch(name: "localhost",
					  port: 8181,
					  routes: routes,
					  responseFilters: [
						(PerfectHTTPServer.HTTPFilter.contentCompression(data: [:]), HTTPFilterPriority.high)])
