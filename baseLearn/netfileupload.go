package baseLearn

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"
	"os"
	"math/rand"
)

//go 文件断点续传
/*import (
	"os"
	"fmt"
	"http"
	"io"
)

const (
	UA = "Golang Downloader from Kejibo.com"
)


func UploadTest() {
	f,err := os.OpenFile("./uploadTest.txt",os.O_RDWR,0666)
	if err!=nil {
		fmt.Println("文件不存在")
	}
	stat,err2 := f.Stat()
	if err2 !=nil {
		fmt.Println("文件状态错误")
	}
	f.Seek(stat.Size(),0)//把文件指针指到文末，当然你说为何不直接用 O_APPEND 模式打开，没错是可以。我这里只是试验。

	url := "http://dl.google.com/chrome/install/696.57/chrome_installer.exe"
	var req http.Request
	req.Method = "GET"
	req.UserAgent = UA
	req.Close = true
	req.URL, err = http.ParseURL(url)
	if err != nil { panic(err) }
	header := http.Header{}
	header.Set("Range", "bytes=" + strconv.Itoa64(stat.Size) + "-")
	req.Header = header
	resp, err := http.DefaultClient.Do(&req)
	if err != nil { panic(err) }
	written, err := io.Copy(f, resp.Body)
	if err != nil { panic(err) }
	println("written: ", written)
}*/

func UploadTest()  {
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	//log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	http.ListenAndServe(":8080", nil)
}


const maxUploadSize = 10 * 1024 * 2014 // 10MB
const uploadPath = "./tmp"
func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		fmt.Println("===============")
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		//fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
		//byt := make([]byte,1024)
		//file.Read(byt)
		//w.Write(byt)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		fmt.Println(filetype)
		switch filetype {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(filetype)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
		fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}