// swift-tools-version:4.0

import PackageDescription

let package = Package(
	name: "D2RollViewerServer",
	products: [
		.executable(name: "D2RollViewerServer", targets: ["D2RollViewerServer"])
	],
	dependencies: [
		.package(url: "https://github.com/PerfectlySoft/Perfect-HTTPServer.git", from: "3.0.0"),
	],
	targets: [
		.target(name: "D2RollViewerServer", dependencies: ["PerfectHTTPServer"])
	]
)
