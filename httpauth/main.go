package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
)

var (
	Username string = "oliver"
	Password string = "mypass"
	Realm    string = "myrealm"
	HA1      string
)

func main() {

	HA1 = fmt.Sprintf("%x", md5.Sum([]byte(Username+":"+Realm+":"+Password)))

	http.HandleFunc("/basic_auth", HandleBasicAuth)
	http.HandleFunc("/digest_auth", HandleDigestAuth)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	}
}

func HandleBasicAuth(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != Username || password != Password {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s", username)
	return
}

func HandleDigestAuth(w http.ResponseWriter, r *http.Request) {
	param, ok := ParseDigestAuth(r)
	if !ok {
		RequireDigestAuth(w, Realm, RandomString(), RandomString())
		return
	}

	HA2 := fmt.Sprintf("%x", md5.Sum([]byte(r.Method+":"+param["uri"])))
	response := fmt.Sprintf("%x", md5.Sum([]byte(HA1+":"+param["nonce"]+":"+param["nc"]+":"+param["cnonce"]+":"+param["qop"]+":"+HA2)))

	if param["username"] != Username || response != param["response"] {
		RequireDigestAuth(w, Realm, RandomString(), RandomString())
		return
	}

	fmt.Fprintf(w, "Welcome %s", param["username"])
	return
}

func RequireDigestAuth(w http.ResponseWriter, realm string, nonce string, opaque string) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf("Digest realm=\"%s\",qop=\"auth\",nonce=\"%s\",opaque=\"%s\"", realm, nonce, opaque))
	w.WriteHeader(http.StatusUnauthorized)
}

func ParseDigestAuth(r *http.Request) (param map[string]string, ok bool) {
	const prefix = "Digest "
	requiredParams := []string{"username", "realm", "nonce", "uri", "qop", "nc", "cnonce", "response", "opaque"}

	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, prefix) {
		return
	}
	param = make(map[string]string)
	for _, pair := range strings.Split(auth[len(prefix):], ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			continue
		}
		param[strings.Trim(parts[0], "\" ")] = strings.Trim(parts[1], "\" ")
	}

	for _, k := range requiredParams {
		if _, exist := param[k]; !exist {
			fmt.Printf("Lack %s\n", k)
			return
		}
	}

	ok = true
	return
}

func RandomString() string {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", randBytes)
}
