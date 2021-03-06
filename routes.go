package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Header      string
	ContentType string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		"",
		"",
		Index,
	},
	Route{
		"IndexFile",
		"GET",
		"/{file}",
		"",
		"",
		Index,
	},
	Route{
		"JobIndex",
		"GET",
		"/api/v1/jobs",
		"",
		"",
		JobIndex,
	},
	Route{
		"JobSummary",
		"GET",
		"/api/v1/jobs/summary",
		"",
		"",
		JobSummary,
	},
	Route{
		"JobShow",
		"GET",
		"/api/v1/jobs/{jobID}",
		"",
		"",
		JobShow,
	},
	Route{
		"JobUpdateJSON",
		"POST",
		"/api/v1/jobs/{jobID}",
		"",
		"",
		JobUpdateJSON,
	},
	Route{
		"JobCreateURLEnc",
		"POST",
		"/api/v1/jobs",
		"Content-Type",
		//"application/x-www-form-urlencoded",
		"application/x-www-form-urlencoded.*",
		JobCreateURLEnc,
	},
	Route{
		"JobCreateJSON",
		"POST",
		"/api/v1/jobs",
		"",
		"",
		JobCreateJSON,
	},
	Route{
		"WorkerIndex",
		"GET",
		"/api/v1/workers",
		"",
		"",
		WorkerIndex,
	},
	Route{
		"WorkerSummary",
		"GET",
		"/api/v1/workers/summary",
		"",
		"",
		WorkerSummary,
	},
	Route{
		"WorkerShow",
		"GET",
		"/api/v1/workers/{workerID}",
		"",
		"",
		WorkerShow,
	},
	Route{
		"WorkerUpdateJSON",
		"POST",
		"/api/v1/workers/{workerID}",
		"",
		"",
		WorkerUpdateJSON,
	},
	Route{
		"WorkerCreateURLEnc",
		"POST",
		"/api/v1/workers",
		"Content-Type",
		//"application/x-www-form-urlencoded",
		"application/x-www-form-urlencoded.*",
		WorkerCreateURLEnc,
	},
	Route{
		"WorkerCreateJSON",
		"POST",
		"/api/v1/workers",
		"",
		"",
		WorkerCreateJSON,
	},
}
