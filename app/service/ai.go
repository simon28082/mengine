package service

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	openai "github.com/sashabaranov/go-openai"
	"github.com/simon/mengine/infrastructure/errors"
	"os"
	"sync"
)

var (
	DefaultAi *Ai
	Token     = os.Getenv("OPENAI_TOKEN")
)

const defaultUserId = "W3oUYJiAaEPV"
const defaultAnswer = `关于使用ChatGPT，未得到回复，可能因为微信的超时机制导致，请稍等一会，并重新 Copy 原问题即可得到答案。`

type Ai struct {
	token       string
	cache       map[string]*cacheValue
	gptClient   *openai.Client
	cacheLock   sync.RWMutex
	requestLock sync.Mutex
}

type cacheValue struct {
	Result string
	Valid  bool
}

func NewAi() *Ai {
	return &Ai{
		token:     Token,
		gptClient: openai.NewClient(Token),
		cache:     make(map[string]*cacheValue),
	}
}

func (a *Ai) Completion(ctx context.Context, req openai.ChatCompletionRequest) (string, error) {
	if len(req.Model) == 0 {
		req.Model = openai.GPT3Dot5Turbo
	}
	if req.MaxTokens == 0 {
		req.MaxTokens = 1024
	}
	if len(req.User) == 0 {
		req.User = defaultUserId
	}
	if len(req.Messages) == 0 {
		return ``, errors.NewDefault(`message length must required`)
	}

	resp, err := a.gptClient.CreateChatCompletion(
		ctx, req,
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ``, err
	}

	spew.Dump(resp.Choices[0].Message.Content)
	spew.Dump(resp)
	return defaultAnswer, nil
}

//
//func (a *Ai) reqHash(req gogpt.CompletionRequest) string {
//	key := strings.Join([]string{req.User, req.Prompt}, `,`)
//	hash := fmt.Sprintf("%x", sha1.Sum([]byte(key)))
//	return hash
//}
//
//func (a *Ai) readCache(req gogpt.CompletionRequest) (*cacheValue, bool) {
//	hash := a.reqHash(req)
//
//	a.cacheLock.RLock()
//	defer a.cacheLock.RUnlock()
//
//	v, ok := a.cache[hash]
//	if ok && v.Valid {
//		delete(a.cache, hash)
//	}
//	return v, ok
//}
//
//func (a *Ai) updateCache(req gogpt.CompletionRequest, result string, valid bool) {
//	a.cacheLock.Lock()
//	defer a.cacheLock.Unlock()
//
//	hash := a.reqHash(req)
//	a.cache[hash] = &cacheValue{
//		Result: result,
//		Valid:  valid,
//	}
//}
//
//func (a *Ai) deleteCache(req gogpt.CompletionRequest) {
//	a.cacheLock.Lock()
//	defer a.cacheLock.Unlock()
//
//	hash := a.reqHash(req)
//	delete(a.cache, hash)
//}
//
//func (a *Ai) Completion(ctx context.Context, req gogpt.CompletionRequest) (string, error) {
//	if len(req.Model) == 0 {
//		req.Model = gogpt.GPT3TextDavinci003
//	}
//	if req.MaxTokens == 0 {
//		req.MaxTokens = 2048
//	}
//	if len(req.User) == 0 {
//		req.User = defaultUserId
//	}
//	req.Stream = true
//	log.Printf("[Info] request user[%s], promot[%s]\n", req.User, req.Prompt)
//
//	result, ok := a.readCache(req)
//	if ok {
//		if result.Valid {
//			return result.Result, nil
//		} else {
//			return defaultAnswer, nil
//		}
//	}
//
//	a.updateCache(req, ``, false)
//
//	go func(ctx context.Context, req gogpt.CompletionRequest) {
//		log.Println(`[Info] will try request to openai completion stream api`)
//
//		stream, err := a.gptClient.CreateCompletionStream(ctx, req)
//		if err != nil {
//			log.Printf(`[Error] create completion stream failed, %s\n`, err)
//			return
//		}
//		defer stream.Close()
//
//		var str = strings.Builder{}
//		for {
//			response, err := stream.Recv()
//			if errors.Is(err, io.EOF) {
//				log.Println(`[Info] read completion stream completed`)
//				break
//			}
//
//			if err != nil {
//				log.Printf(`[Error] read completion stream failed, %s\n`, err)
//				break
//			}
//
//			if len(response.Choices) > 0 {
//				str.WriteString(response.Choices[0].Text)
//			}
//		}
//
//		rs := str.String()
//		if len(rs) > 0 {
//			a.updateCache(req, rs, true)
//		} else {
//			a.deleteCache(req)
//		}
//	}(ctx, req)
//
//	return defaultAnswer, nil
//}
