package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Server struct {
	Name string            `json:"name"`
	Addr string            `json:"addr"`
	Tags map[string]string `json:"tags"`
}

func AllServers() ([]*Server, error) {
	var servers []*Server

	files, _ := ioutil.ReadDir(DataPath)
	for _, file := range files {
		if path.Ext(file.Name()) == ".json" {
			objectPath := path.Join(DataPath, file.Name())
			data, err := ioutil.ReadFile(objectPath)
			if err != nil {
				log.Println("can't read file"+objectPath+":", err.Error())

				return nil, err
			}

			server := &Server{}
			err = json.Unmarshal(data, server)
			if err != nil {
				log.Println("can't unmarshal content from "+objectPath+":", err.Error())

				return nil, err
			}

			servers = append(servers, server)
		}
	}

	return servers, nil
}
func GetServer(addr string) (*Server, error) {
	objectPath := path.Join(DataPath, addr+".json")
	data, err := ioutil.ReadFile(objectPath)
	if err != nil {
		if os.IsExist(err) {
			log.Println("can't read file"+objectPath+":", err.Error())

			return nil, err
		}

		return nil, nil
	}

	server := &Server{}
	err = json.Unmarshal(data, server)
	if err != nil {
		log.Println("can't unmarshal content from "+objectPath+":", err.Error())

		return nil, err
	}

	return server, nil
}

func DeleteServer(addr string) (bool, error) {
	objectPath := path.Join(DataPath, addr+".json")
	err := os.Remove(objectPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
