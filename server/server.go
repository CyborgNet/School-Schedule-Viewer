package server

import "../config"

func Init() {
	r := newRouter()
	r.Run(config.GetServerAddress())
}