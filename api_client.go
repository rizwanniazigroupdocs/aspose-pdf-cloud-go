 /**
 *
 *   Copyright (c) 2018 Aspose.PDF Cloud
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
package asposepdfcloud

import (
    "bytes"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "errors"
    "io"
    "mime/multipart"
    "net/http"
    "net/url"
    "time"
    "os"
    "path/filepath"
    "reflect"
    "regexp"
    "strings"
    "unicode/utf8"
    "strconv"
)

var (
    jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
    xmlCheck = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Aspose.PDF Cloud API Reference API v2.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
    cfg     *Configuration
    common     service         // Reuse a single struct instead of allocating one for each service on the heap.

     // API Services
    PdfApi    *PdfApiService
}

type service struct {
    client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
    if cfg.HTTPClient == nil {
        cfg.HTTPClient = http.DefaultClient
    }

    c := &APIClient{}
    c.cfg = cfg
    c.common.client = c

    // API Services
    c.PdfApi = (*PdfApiService)(&c.common)

    return c
}

func atoi(in string) (int, error) {
    return strconv.Atoi(in)
}


// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
    if len(contentTypes) == 0 {
        return ""
    }
    if contains(contentTypes, "application/json") {
        return "application/json"
    }
    return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
    if len(accepts) == 0 {
        return ""
    }

    if contains(accepts, "application/json") {
        return "application/json"
    }

    return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
    for _, a := range haystack {
        if strings.ToLower(a) == strings.ToLower(needle) {
            return true
        }
    }
    return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
    // Make sure there is an object.
    if obj == nil {
        return nil
    }

    // Check the type is as expected.
    if reflect.TypeOf(obj).String() != expected {
        return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
    }
    return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
    var delimiter string

    switch collectionFormat {
    case "pipes":
        delimiter = "|"
    case "ssv":
        delimiter = " "
    case "tsv":
        delimiter = "\t"
    case "csv":
        delimiter = ","
    }

    if reflect.TypeOf(obj).Kind() == reflect.Slice {
        return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
    }

    return fmt.Sprintf("%v", obj)
}

// callAPI do the request. 
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
     return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath (path string) {
    c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest (
    path string, method string,
    postBody interface{},
    headerParams map[string]string,
    queryParams url.Values,
    formParams url.Values,
    fileName string,
    fileBytes []byte) (localVarRequest *http.Request, err error) {

    var body *bytes.Buffer

    // Detect postBody type and post.
    if postBody != nil {
        contentType := headerParams["Content-Type"]
        if contentType == "" {
            contentType = detectContentType(postBody)
            headerParams["Content-Type"] = contentType
        }

        body, err = setBody(postBody, contentType)
        if err != nil {
            return nil, err
        }
    }

    // add form parameters and file if available.
    if len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
        if body != nil {
            return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
        }
        body = &bytes.Buffer{}
        w := multipart.NewWriter(body)

        for k, v := range formParams {
            for _, iv := range v {
                if strings.HasPrefix(k, "@") { // file
                    err = addFile(w, k[1:], iv)
                    if err != nil {
                        return nil, err
                    }
                } else { // form value
                    w.WriteField(k, iv)
                }
            }
        }
        if len(fileBytes) > 0 && fileName != "" {
            body = bytes.NewBuffer(fileBytes)
            
            // Set the Boundary in the Content-Type
            headerParams["Content-Type"] = "application/octet-stream"
        }
        
        // Set Content-Length
        headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
        w.Close()
    }

    // Setup path and query parameters
    url, err := url.Parse(path)
    if err != nil {
        return nil, err
    }

    // Adding Query Param
    query := url.Query()
    for k, v := range queryParams {
        for _, iv := range v {
            query.Add(k, iv)
        }
    }

    // Encode the parameters.
    url.RawQuery = query.Encode()

    // Generate a new request
    if body != nil {
        localVarRequest, err = http.NewRequest(method, url.String(), body)
    } else {
        localVarRequest, err = http.NewRequest(method, url.String(), nil)
    }
    if err != nil {
        return nil, err
    }

    // add header parameters, if any
    if len(headerParams) > 0 {
        headers := http.Header{}
        for h, v := range headerParams {
            headers.Set(h, v)
        }
        localVarRequest.Header = headers
    }

    // Override request host, if applicable
    if c.cfg.Host != "" {
        localVarRequest.Host = c.cfg.Host
    }
    
    // Add the user agent to the request.
    localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)
    
    // Add auth
    err = c.addAuth(localVarRequest)
    if err != nil {
        return nil, err
    }

    for header, value := range c.cfg.DefaultHeader {
        localVarRequest.Header.Add(header, value)
    }
    
    return localVarRequest, nil
}


// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    part, err := w.CreateFormFile(fieldName, filepath.Base(path))
    if err != nil {
        return err
    }
    _, err = io.Copy(part, file)

    return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) (error) {
    return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
    if bodyBuf == nil {
        bodyBuf = &bytes.Buffer{}
    }

    if reader, ok := body.(io.Reader); ok {
        _, err = bodyBuf.ReadFrom(reader)
    } else if b, ok := body.([]byte); ok {
        _, err = bodyBuf.Write(b)
    } else if s, ok := body.(string); ok {
        _, err = bodyBuf.WriteString(s)
    } else if jsonCheck.MatchString(contentType) {
        err = json.NewEncoder(bodyBuf).Encode(body)
    } else if xmlCheck.MatchString(contentType) {
        xml.NewEncoder(bodyBuf).Encode(body)
    }

    if err != nil {
        return nil, err
    }

    if bodyBuf.Len() == 0 {
        err = fmt.Errorf("Invalid body type %s\n", contentType)
        return nil, err
    }
    return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
    contentType := "text/plain; charset=utf-8"
    kind := reflect.TypeOf(body).Kind()
    
    switch kind {
    case reflect.Struct, reflect.Map, reflect.Ptr:
        contentType = "application/json; charset=utf-8"
    case reflect.String:
        contentType = "text/plain; charset=utf-8"
    default:
        if b, ok := body.([]byte); ok {
            contentType = http.DetectContentType(b)
        } else if kind == reflect.Slice {
            contentType = "application/json; charset=utf-8"
        }
    }

    return contentType
}


// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
    cc := cacheControl{}
    ccHeader := headers.Get("Cache-Control")
    for _, part := range strings.Split(ccHeader, ",") {
        part = strings.Trim(part, " ")
        if part == "" {
            continue
        }
        if strings.ContainsRune(part, '=') {
            keyval := strings.Split(part, "=")
            cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
        } else {
            cc[part] = ""
        }
    }
    return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) (time.Time) {
    // Figure out when the cache expires.
    var expires time.Time
    now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
    if err != nil {
        return time.Now()
    }
    respCacheControl := parseCacheControl(r.Header)
    
    if maxAge, ok := respCacheControl["max-age"]; ok {
        lifetime, err := time.ParseDuration(maxAge + "s")
        if err != nil {
            expires = now
        }
        expires = now.Add(lifetime)
    } else {
        expiresHeader := r.Header.Get("Expires")
        if expiresHeader != "" {
            expires, err = time.Parse(time.RFC1123, expiresHeader)
            if err != nil {
                expires = now
            }
        }
    }
    return expires
}

func strlen(s string) (int) {
    return utf8.RuneCountInString(s)
}

// addAuth add Authorization header to request
func (a *APIClient) addAuth(request *http.Request) (err error) {
       if (a.cfg.AccessToken == "") {
               if err := a.RequestOauthToken(); err != nil {
                       return err
               }
       }
       request.Header.Add("Authorization", "Bearer " + a.cfg.AccessToken)
       return nil
}

// RequestOauthToken function for requests OAuth token
func (a *APIClient) RequestOauthToken() (error) {

       resp, err := http.PostForm(strings.Replace(a.cfg.BasePath, "/v2.0", "/oauth2/token", -1), url.Values{
               "grant_type": {"client_credentials"},
               "client_id": {a.cfg.AppSid},
               "client_secret": {a.cfg.AppKey}})

       if err != nil {
               return err
       }
       defer resp.Body.Close()

       var tr TokenResp
       if err = json.NewDecoder(resp.Body).Decode(&tr); err != nil {
               return err
       }
       a.cfg.AccessToken = tr.AccessToken
       a.cfg.RefreshToken = tr.RefreshToken
       return nil
}

// TokenResp represents data returned by GetAccessToken and RefreshToken as HTTP response body.
type TokenResp struct {
       AccessToken                         string `json:"access_token"`
       TokenType                           string `json:"token_type"`
       ExpiresIn                           int64  `json:"expires_in"`
       RefreshToken                        string `json:"refresh_token"`
       ClientID                            string `json:"client_id"`
       ClientRefreshTokenLifeTimeInMinutes string `json:"clientRefreshTokenLifeTimeInMinutes"`
       Issued                              string `json:".issued"`
       Expires                             string `json:".expires"`
}
