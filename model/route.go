package model

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

type Version struct {
	Routes  Routes
	Version string
}

type Versions []Version
