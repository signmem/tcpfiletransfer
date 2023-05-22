package g

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"errors"
	"encoding/json"
)

func HTTPCheckContent(r *http.Request) ( fileInfo BackupPeriod, err error) {
	if r.ContentLength == 0 {
		msg := fmt.Sprintf("error: body is blank")
		return fileInfo, errors.New(msg)
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		msg := fmt.Sprintf("error: body not json format")
		return fileInfo, errors.New(msg)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		msg := fmt.Sprintf("error: body read error")
		return fileInfo, errors.New(msg)
	}

	err = json.Unmarshal(body, &fileInfo)

	if err != nil {
		msg := fmt.Sprintf("error: body json unmarshar error")
		return fileInfo, errors.New(msg)
	}

	if len(fileInfo.FSType) == 0  {
		msg := fmt.Sprintf("error: fstype empty")
		return fileInfo, errors.New(msg)
	}

	// fstype ?????????  ['temp', 'perm']

	if  fileInfo.FSType != "temp" && fileInfo.FSType != "perm" {
		msg := fmt.Sprintf("error: fstype not in temp or perm")
		return fileInfo, errors.New(msg)
	}

	//      "YYYYMMDD/binlog/backup_10.205.32.27_3307/file"
	//      "YYYYMMDD/mydumper/backup_10.205.54.133_3308"
	//      "YYYYMMDD/xtrabackup/backup_10.205.54.192_3309/"

	if len(fileInfo.FSPath) == 0 {
		msg := fmt.Sprintf("error: fspath empty")
		return fileInfo, errors.New(msg)
	}

	path := strings.Split(fileInfo.FSPath, "/")

	// path ? 2 ?????????? [ binlog, mydumper, xtrabackup ]

	if path[1] != "binlog" && path[1] != "mydumper" && path[1] != "xtrabackup" {
		msg := fmt.Sprintf("error: fspath not in binlog, mydumper, xtrabackup")
		return fileInfo, errors.New(msg)
	}

	if IsTimeFormat(path[0]) == false {
		msg := fmt.Sprintf("error: fspath not include datetime format")
		return fileInfo, errors.New(msg)
	}

	if len(fileInfo.GFSCluster) == 0 {
		msg := fmt.Sprintf("error: gfscluster empty")
		return fileInfo, errors.New(msg)
	}

	// token ?? rpc client ?? server ????,?? dba ????????? cephserver ???

	return fileInfo, nil
}


func GetClientIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", errors.New("No valid ip found")

}