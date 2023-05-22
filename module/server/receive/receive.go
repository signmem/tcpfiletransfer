package receive

import (
	"github.com/signmem/tcpfiletransfer/g"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const BUFFERSIZE = 1024

// https://mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket

func TcpListen() {
	serverAddr := g.Config().TcpServer.Server
	port := g.Config().TcpServer.Port
	tcpServer := serverAddr + ":" + port

	server, err := net.Listen( "tcp", tcpServer)
	if err != nil {
		g.Logger.Debugf("Listen %s err:%s", tcpServer, err)
		return
	}

	defer server.Close()


	g.Logger.Debug("server start, waiting for connections...")

	for {

		connection, err := server.Accept()

		if err != nil {
			g.Logger.Debugf("Listen Accept() %s err:%s", tcpServer, err)
			os.Exit(1)
		}

		g.Logger.Debug("client connectiond")

		go receiveFileFromTCP(connection)

	}
}

func receiveFileFromTCP(connection net.Conn) {

	defer connection.Close()

	/*
	connection.Write([]byte(fileSize))
	connection.Write([]byte(fileName))
	connection.Write([]byte(fileTempPath))
	connection.Write([]byte(fileTYpe))
	*/

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)
	bufferFileTempPath := make([]byte, 256)
	bufferFileType := make([]byte, 16)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	g.Logger.Debugf("filesize %s", fileSize)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName),":")
	g.Logger.Debugf("bufferFileName %s", string(bufferFileName))

	connection.Read(bufferFileTempPath)
	filePath := strings.TrimRight(strings.Trim(string(bufferFileTempPath),
		":"), "/")
	g.Logger.Debugf("bufferFileTempPath %s", string(bufferFileTempPath))

	connection.Read(bufferFileType)
	fileType := strings.Trim(string(bufferFileType), ":")
	g.Logger.Debugf("bufferFileType %s", string(bufferFileType))


	g.Logger.Debugf("filename:%s, size:%d, path:%s, type:%s", fileName,
		fileSize, filePath, fileType)


	destFile := g.Config().BackupPath + "/" + fileType + "/" + filePath +
		"/"	+ fileName
	
	g.CheckAndCreateDir(destFile)

	connection.Read(bufferFileSize)

	newFile, err := os.Create(destFile)
	if err != nil {
		g.Logger.Debugf("create file %s err:%s", destFile, err)
		return
	}

	defer newFile.Close()
	var receivedBytes int64
	for {
		if (fileSize - receivedBytes) < BUFFERSIZE {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BUFFERSIZE)
		receivedBytes += BUFFERSIZE
	}
	g.Logger.Debugf("Received file %s, completely!", fileName)
}


