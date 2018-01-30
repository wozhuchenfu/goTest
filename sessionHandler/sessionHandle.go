package sessionHandler

import (
	"time"
	"net/http"
	"fmt"
	"sync"
	"crypto/rand"
	"encoding/base64"
	"net/url"
)

/*type Cookie struct {
	Name 		string
	Value 		string
	Path  		string
	Domain  	string
	Expaires  	time.Time
	RawExpaires string
	MaxAge		int
	Secure 		bool
	HttpOnly	bool
	Raw			string
	Unparsed	[]string
}*/

func SetCookie(w http.ResponseWriter)  {
	expiration := time.Now()
	expiration = expiration.AddDate(0,0,1)//一天
	cookie := http.Cookie{Name:"username",Value:"qi",Expires:expiration}
	http.SetCookie(w,&cookie)
}

func GetCookie(r *http.Request)  {
	cookie,_ := r.Cookie("username")
	fmt.Println(cookie.Value)

	for _,cookie := range r.Cookies() {
		fmt.Println(cookie.Name+":"+cookie.Value)
	}
}

type SessionManager struct {
	cookieName	string		//private cookiename
	lock 		*sync.Mutex
	provider 	Provider
	maxLifeTime	int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var provides = make(map[string]Provider)

func NewManager(provideName,cookieName string,maxLifeTime int64) (*SessionManager,error) {
	provider,ok := provides[provideName]
	if !ok {
		return nil,fmt.Errorf("session:unkown provide %q (forgotten import?)",provideName)
	}
	return &SessionManager{provider:provider,cookieName:cookieName,maxLifeTime:maxLifeTime},nil

}
//创建全局唯一sessionid
func (manager *SessionManager) sessionId() string {
	b := make([]byte,32)
	if _,err := rand.Read(b);err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
//创建session
func (manager *SessionManager) SessionStart(w http.ResponseWriter,r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie,err := r.Cookie(manager.cookieName)
	if err!=nil || cookie.Value == "" {
		sid := manager.sessionId()
		session,_ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name:manager.cookieName,Value:url.QueryEscape(sid),Path:"/",HttpOnly:true,MaxAge:int(manager.maxLifeTime)}
		http.SetCookie(w,&cookie)
	} else {
		sid,_ := url.QueryUnescape(cookie.Value)
		session,_ = manager.provider.SessionRead(sid)
	}
	return
}

func (manager *SessionManager) SessionDestory(w http.ResponseWriter,r *http.Request) {
	cookie,err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name:manager.cookieName,Path:"/",HttpOnly:true,Expires:expiration,MaxAge:-1}
		http.SetCookie(w,&cookie)
	}
}

func (manager *SessionManager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime)*time.Second, func() {
		manager.GC()
	})

}










