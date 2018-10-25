function directData(requestDetails) {
	console.log("$$$$$$$$$$$$$$$$$$$$$$$$");

	var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance
	xmlhttp.open("POST", "http://localhost:9299/data");
	xmlhttp.setRequestHeader("Content-Type", "text/plain");

	xmlhttp.onload = function () {
		console.log("##################");
	};

	xmlhttp.send(JSON.stringify(requestDetails));

	console.log(requestDetails);
	console.log("$$$$$$$$$$$$$$$$$$$$$$$$");
	// "localhost:9299/data"
}


browser.webRequest.onBeforeRequest.addListener(
  directData,
  {urls: ["<all_urls>"]},
);

// "localhost:9299/data"]},
