package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server up and listening on port 8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Client %v connected.\n", c.RemoteAddr())
	defer c.Close()

	// Create a file to save the image
	file, err := os.Create("./brain_image.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Copy the image from the client to the file
	_, err = io.Copy(file, c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Image received and saved successfully.")

	//Call the Python program
	cmd := exec.Command("python", "brain_MRI_testing.py")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Python program executed successfully.")

	err = sendFile(c, "result.txt")
	if err != nil {
		fmt.Println("Error sending file:", err)
	}
}

func sendFile(conn net.Conn, filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Log the file size
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	fmt.Printf("Sending file of size %d bytes\n", stat.Size())

	// Copy the file content to the connection
	_, err = io.Copy(conn, file)
	if err != nil {
		return err
	}

	fmt.Println("File sent successfully.")

	return nil
}
