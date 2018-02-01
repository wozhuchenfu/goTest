package goBlockChain

import (
"crypto/sha256"
"encoding/hex"
"time"
"net/http"
"github.com/gorilla/mux"//用来写web handler
"encoding/json"
"io"
"github.com/davecgh/go-spew/spew"//spew可以帮助我们在console中直接查看struct和slice两种数据结构
"os"
"log"
"github.com/joho/godotenv"//读取项目根目录中的.env配置文件，这样就不用将http port之类的配置硬编码进代码中
"fmt"
	//"strconv"
)

func BlockChainTest() {
	err := godotenv.Load()
	if err!=nil {
		log.Fatal(err)
	}
	go func() {
		t := time.Now()
		//genesisBlock为创世块，通过它来初始化区块链，第一个块的PresisBlock是空的。
		genesisBlock := Block{0,t.String(),0,"",""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain,genesisBlock)
	}()
	log.Fatal(run())

}

type Block struct {
	Index int //块在整个链中的位置
	Timestamp string //块生成时间
	BPM int //通过SHA256算法生产的散列值
	Hash string //代表一个块的SHA256散列值
	PrevHash string //每分钟心跳数。
}

//定义一个链
var Blockchain []Block

//我们使用散列算法（SHA256）来确定和维护链中块和块正确的顺序，
// 确保每一个块的 PrevHash 值等于前一个块中的 Hash 值，这样就以正确的块顺序构建出链：

/*
	为什么用散列？

	在节省空间的前提下去唯一标识数据。散列是用整个块的数据计算得出，在我们的例子中，将整个块的数据通过 SHA256 计算成一个定长不可伪造的字符串。

	维持链的完整性。通过存储前一个块的散列值，我们就能够确保每个块在链中的正确顺序。任何对数据的篡改都将改变散列值，同时也就破坏了链。
	以我们从事的医疗健康领域为例，比如有一个恶意的第三方为了调整“人寿险”的价格，而修改了一个或若干个块中的代表不健康的 BPM 值，那么整个链都变得不可信了。
 */
//计算给定数据的SHA256散列值
func calculateHaash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	bytes := []byte(record)
	h.Write(bytes)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block,BPM int) (Block,error) {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHaash(newBlock)
	return newBlock,nil

}

//校验块
func isBlockValid(newBlock,oldBlock Block) bool {
	if oldBlock.Index+1!=newBlock.Index {
		return false
	}
	if oldBlock.Hash!= newBlock.PrevHash {
		return false
	}
	if calculateHaash(newBlock)!=newBlock.Hash {
		return false
	}
	return true
}

func replaceChain(newBlocks []Block)  {
	if len(newBlocks)>len(Blockchain) {
		Blockchain = newBlocks
	}
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	fmt.Println(httpAddr)
	log.Println("Listening on",os.Getenv("ADDR"))
	s := &http.Server{
		Addr:":8080",
		Handler:mux,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	if err := s.ListenAndServe();err!=nil {
		return err
	}
	return nil
}


func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/",handleGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/",handleWriteBlock).Methods("POST")
	return muxRouter
}

//GET请求的handler
func handleGetBlockChain(w http.ResponseWriter,r *http.Request)  {
	bytes,err := json.MarshalIndent(Blockchain,"","")
	if err!=nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(w,string(bytes))
}

type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter,r *http.Request)  {

	//fmt.Println("POST")
	//r.ParseForm()
	//
	//s := r.PostFormValue("BPM")
	//fmt.Println(s)
	//BMP := r.Form["BPM"]
	//fmt.Println(BPM)

	//i,err := strconv.Atoi(s)
	//if err!=nil {
	//	log.Fatal(err.Error())
	//}
	var m Message
	//m.BPM = i
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m);err!=nil {
		respondWithJson(w,r,http.StatusInternalServerError,r.Body)
		return
	}
	defer r.Body.Close()
	newbLock,err := generateBlock(Blockchain[len(Blockchain)-1],m.BPM)
	if err!=nil {
		respondWithJson(w,r,http.StatusInternalServerError,m)
		return
	}
	if isBlockValid(newbLock,Blockchain[len(Blockchain)-1]) {
		newBlockChain := append(Blockchain,newbLock)
		replaceChain(newBlockChain)
		spew.Dump(Blockchain)
	}
	respondWithJson(w,r,http.StatusCreated,newbLock)
}

func respondWithJson(w http.ResponseWriter,r *http.Request,code int,payload interface{})  {
	response,err := json.MarshalIndent(payload,"","")
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}














