package cashe

import (
	"context"
	"errors"

	"log"
	"sync"
	"time"

	"github.com/knmsh08200/Blog_test/internal/blog"
	"github.com/knmsh08200/Blog_test/internal/metrics"
	"github.com/knmsh08200/Blog_test/internal/model"
)

var (
	errNotFound = errors.New("id not found")
)

// или смысл в том, чтобы создать новый интерфейс и в кэше уже использовать новый - просто мысль, чтобы не забыть

type blogEntity struct {
	updated time.Time
	data    model.FindList
}

type BlogCashe struct {
	blogProvider blog.ListRepository // не понимаю что это делает - функция декоратора как я понимаю
	timeToUpdate time.Duration
	timeToDelete time.Duration
	ctx          context.Context
	cancelFunc   context.CancelFunc

	mx   sync.RWMutex
	data map[int]blogEntity
}

func NewCashe(ctx context.Context, timeUpdate time.Duration, timeDelete time.Duration, blogProvider blog.ListRepository) *BlogCashe {
	ctx, cancelFunc := context.WithCancel(ctx) // не на 100 процентов понима - нужно обсудить
	return &BlogCashe{
		data:         make(map[int]blogEntity),
		timeToUpdate: timeUpdate,
		timeToDelete: timeDelete,
		blogProvider: blogProvider,
		ctx:          ctx,
		cancelFunc:   cancelFunc,
	}

}

func (b *BlogCashe) StartCleaner() {
	go b.cleaner()
}

func (b *BlogCashe) Shutdown() {
	log.Println("stopping cache hit")
	b.cancelFunc()
}
func (b *BlogCashe) getBlogByID(id int) (*blogEntity, error) {
	b.mx.RLock()
	defer b.mx.RUnlock()
	blog, ok := b.data[id]
	if !ok {
		log.Printf("ID: %d not found in cache", id)
		return nil, errNotFound
	}

	return &blog, nil
}

// декоратор

func (b *BlogCashe) set(id int, blog model.FindList) {

	b.mx.Lock()
	defer b.mx.Unlock()
	log.Printf("Updating cashe for ID:%d", id)
	b.data[id] = blogEntity{
		updated: time.Now(),
		data:    blog,
	}
	log.Printf("Updated cashe for ID:%d", id)
}

func (b *BlogCashe) GetAllBlogs(ctx context.Context, limit, offset int) ([]model.ListResponse, model.Meta, error) {
	return b.blogProvider.GetAllBlogs(ctx, limit, offset)
}

func (b *BlogCashe) CreateBlog(ctx context.Context, list model.List) (int, error) {
	return b.blogProvider.CreateBlog(ctx, list)
}

func (b *BlogCashe) DeleteBlog(ctx context.Context, d int) (int64, error) {
	return b.blogProvider.DeleteBlog(ctx, d)
}

// работа с бд
func (b *BlogCashe) CounterUserBlog(userID int) (int, error) {
	return b.blogProvider.CounterUserBlog(userID)
}

func (b *BlogCashe) FindBlog(ctx context.Context, id int) (model.FindList, error) {
	start := time.Now()
	blog, err := b.getBlogByID(id)
	if err != nil || time.Since(blog.updated) > b.timeToUpdate {
		fetchedBlog, err := b.blogProvider.FindBlog(ctx, id)
		if err != nil {
			log.Printf("Error fetching data from DB for ID:%d, %v", id, err)
			return model.FindList{}, err // wrap -READABLE
		}
		b.set(id, fetchedBlog)
		metrics.ObserveCacheMiss(time.Since(start).Seconds())
		log.Printf("Fetched blog from DB and updated cashe for ID:%d", id)
		return fetchedBlog, nil
	}
	metrics.ObserveCacheHit(time.Since(start).Seconds())
	log.Printf("Cashe hit for ID:%d ", id)
	return blog.data, nil
}

// worker
func (b *BlogCashe) cleaner() {
	for { // зачем нам нужен for
		select {
		case <-b.ctx.Done():
			log.Println("cache cleaner stopped")
			return

		case <-time.After(3 * time.Hour): // отработать данную функцию
			b.mx.Lock()
			for id, blog := range b.data {
				if time.Since(blog.updated) > b.timeToDelete {
					delete(b.data, id)
					log.Printf("Блог '%d' удален из кэша из-за истечения TTL\n", id)
				}
			}
			b.mx.Unlock()
		}
	}

}
