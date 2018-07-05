package CAByGo

import (
	"math/rand"
	"time"
	"crypto/x509/pkix"
	"crypto/x509"
	"crypto/rsa"
	"math/big"
	cr "crypto/rand"
	"fmt"
	"os"
	"encoding/pem"
	"io/ioutil"
	"encoding/asn1"
)

/*
在go语言提供的系统包中包含了大量和数字证书有关的方法。在这些方法中就有私钥生成的方法、私钥解析的方法、证书请求生成的方法、
证书生成的方法等等。通过这些方法应该能够实现和openssl命令类似的功能。
仿照openssl生成证书的流程（从私钥的生成—>证书请求的生成—>证书的生成）用go语言进行模拟。
 */

func CAByGoTest() {

	baseinfo := &CertInformation{Country:[]string{"CN"},
	Organization:[]string{"WS"},
	IsCA:true,
	OrganizationalUnit:[]string{"work-stacks"},
	EmailAddress:[]string{"wozhuchenfu@qq.com"},
	Locality:[]string{"BeiJing"},
	Province:[]string{"BeiJing"},
	CommonName:"work-stacks",
	CrtName:"test_root.crt",
	KeyName:"test_root.key"}
	err := CreateCRT(nil,nil,baseinfo)
	if err!=nil {
		fmt.Println("Create crt error,Error info:", err)
		return
	}
	crtinfo := baseinfo
	crtinfo.IsCA = false
	crtinfo.CrtName = "test_server.crt"
	crtinfo.KeyName = "test_server.key"
	crtinfo.Names = []pkix.AttributeTypeAndValue{{asn1.ObjectIdentifier{2,1,3},"MAC_ADDR"}}//添加扩展字段用来做自定义使用
	crt,pri,err := Parse(baseinfo.CrtName,baseinfo.KeyName)
	if err!=nil {
		fmt.Println("Parse crt error,Error info:", err)
		return
	}
	err = CreateCRT(crt,pri,crtinfo)
	if err!=nil {
		fmt.Println("Create crt error,Error info:", err)
	}

	/*os.Remove(baseinfo.CrtName)
	os.Remove(baseinfo.KeyName)
	os.Remove(crtinfo.CrtName)
	os.Remove(crtinfo.KeyName)*/
}
//使用x509标准库创建自签名证书和签发名其他证书
/*
go使用时间作为种子生成随机数
设置时间种子使用time包
生成随机数需要math/rand包
打印输出使用fmt包

不设置时间种子的话，每次生成的rand值相同
 */
func init()  {
	rand.Seed(time.Now().UnixNano())
}

type CertInformation struct {
	Country	   []string
	Organization   []string
	OrganizationalUnit []string
	EmailAddress []string
	Province  []string
	Locality  []string
	CommonName  string
	CrtName,KeyName string
	IsCA bool
	Names  []pkix.AttributeTypeAndValue
}

func CreateCRT(RootCa *x509.Certificate,RootKey *rsa.PrivateKey,info *CertInformation) error{
	Crt := newCertificate(info)
	key,err := rsa.GenerateKey(cr.Reader,2048)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	var CAbufs []byte
	if RootCa == nil || RootKey == nil {
		//创建签名证书
		CAbufs,err = x509.CreateCertificate(cr.Reader,Crt,Crt,&key.PublicKey,key)
	}else {
		//使用根证书签名
		CAbufs,err = x509.CreateCertificate(cr.Reader,Crt,RootCa,&key.PublicKey,RootKey)
	}
	if err!=nil {
		return err
	}
	err = write(info.CrtName,"CERTIFICATE",CAbufs)
	if err!=nil {
		return err
	}
	CAbufs = x509.MarshalPKCS1PrivateKey(key)
	return write(info.KeyName,"PRIVATE KEY",CAbufs)
}

func newCertificate(info *CertInformation) *x509.Certificate {
		return &x509.Certificate{
			SerialNumber:big.NewInt(rand.Int63()),
			Subject:pkix.Name{
				Country:info.Country,
				Organization:info.Organization,
				OrganizationalUnit:info.OrganizationalUnit,
				Province:info.Province,
				CommonName:info.CommonName,
				Locality:info.Locality,
				ExtraNames:info.Names,
			},
			NotBefore:time.Now(),//证书开始时间
			NotAfter:time.Now().AddDate(20,0,0),//证书结束时间
			BasicConstraintsValid:true,
			IsCA:info.IsCA,//是否是根证书
			ExtKeyUsage:[]x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth,x509.ExtKeyUsageServerAuth},
			KeyUsage:x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
			EmailAddresses:info.EmailAddress,
		}
}

//编码写入文件
func write(fileName string,Type string,p []byte) error {
	File,err := os.Create(fileName)
	defer File.Close()
	if err!=nil {
		return err
	}
	var b *pem.Block = &pem.Block{Bytes:p,Type:Type}
	return pem.Encode(File,b)
}

func Parse(crtPath ,KeyPath string) (rootcertificate *x509.Certificate,rootPrivateKey *rsa.PrivateKey,err error) {
	rootcertificate,err = ParseCrt(crtPath)
	if err!=nil {
		return
	}
	rootPrivateKey,err = ParseKey(KeyPath)
	return
}

func ParseCrt(path string) (*x509.Certificate,error) {
	buf,err := ioutil.ReadFile(path)
	if err!=nil {
		return nil,err
	}
	p := &pem.Block{}
	p,buf = pem.Decode(buf)
	return x509.ParseCertificate(p.Bytes)
}

func ParseKey(path string) (*rsa.PrivateKey,error) {
	buf,err := ioutil.ReadFile(path)
	if err!=nil {
		return nil,err
	}
	p:=&pem.Block{}
	p,buf = pem.Decode(buf)
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}
