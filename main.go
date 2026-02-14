package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	HIGHLIGHT_COLOR = "\033[1;38;2;255;0;128m"
	RESET_COLOR     = "\033[0m"
)

var dir string

// ________________________________________________________________________FUNCS

func main() {
	port := 1024
	ip := getLocalIp()

	execPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err.Error())
	}

	dir = filepath.Dir(execPath)
	uploadedFolder := filepath.Join(dir, "uploaded")

	os.MkdirAll(uploadedFolder, 0755)

	http.HandleFunc("/", homeHandler)

	for ; port < 0xffff; port++ {
		addr := fmt.Sprintf("%s:%d", ip, port)

		fmt.Printf("Runing on :%shttp://%-25s%s\n",
			HIGHLIGHT_COLOR,
			addr,
			RESET_COLOR)

		if http.ListenAndServe(addr, nil) != nil {
			fmt.Printf("\033[1A")
		}
	}
}

//                                                                             |
// ----------------------------------------------------------------------------|
//                                                                             |

func getLocalIp() string {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ip := ipNet.IP.To4()
			if ip != nil {
				return ip.String()
			}
		}
	}
	return ""
}

//                                                                             |
// ----------------------------------------------------------------------------|
//                                                                             |

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, filepath.Join(dir, "index.html"))
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(1 << 62)

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer file.Close()

		ts := time.Now().Format("20060102_150405")
		newName := fmt.Sprintf("%s_%s", ts, header.Filename)
		dstPath := filepath.Join("uploaded", newName)

		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, file)
		w.Write([]byte("ok"))
	}
}
