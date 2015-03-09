package main

import (
	"net/http"
	"strings"

	"github.com/ant0ine/go-json-rest/rest"
)

func HandlerAllServers(res rest.ResponseWriter, req *rest.Request) {
	tags := req.URL.Query()

	servers, err := AllServers()
	if err != nil {
		rest.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tags) > 0 {
		for tag, _ := range tags {
			for i, server := range servers {
				if server.Tags[tag] == tags.Get(tag) {
					continue
				}

				servers = append(servers[:i], servers[i+1:]...)
			}
		}
	}

	res.WriteJson(servers)
}

func HandlerGetServer(res rest.ResponseWriter, req *rest.Request) {
	addr := getAddr(req)
	only := req.URL.Query().Get("only")

	server, err := GetServer(addr)
	if server == nil {
		if err == nil {
			rest.NotFound(res, req)
			return
		} else {
			rest.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	switch {
	case only == "":
		res.WriteJson(server)
	case only == "name":
		res.WriteJson(server.Name)
	case only == "addr":
		res.WriteJson(server.Addr)
	case only == "tags":
		res.WriteJson(server.Tags)
	case strings.Contains(only, "tags."):
		tag := strings.Split(only, ".")[1]
		if value, ok := server.Tags[tag]; ok {
			res.WriteJson(value)
		} else {
			rest.Error(res, "tag \""+tag+"\" isn't found", http.StatusBadRequest)
		}
	default:
		rest.Error(res, "\""+only+"\" isn't allowed", http.StatusBadRequest)
	}

}

func HandlerDeleteServer(res rest.ResponseWriter, req *rest.Request) {
	ok, err := DeleteServer(getAddr(req))
	if ok == false {
		if err == nil {
			rest.NotFound(res, req)
			return
		} else {
			rest.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	res.WriteHeader(http.StatusOK)
}

func getAddr(req *rest.Request) string {
	addr := req.PathParam("addr")
	if addr == "my" {
		addr = strings.Split(req.RemoteAddr, ":")[0]
	}

	return addr
}
