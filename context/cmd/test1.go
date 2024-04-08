package main

import (
	"context"
	"fmt"
	"time"
)

//олучение сигнала отмены контекста.
//Отмена контекста через 500 миллисекунд.

type DB struct {
}

type User struct {
	Name string
}

func (d *DB) SelectUser(ctx context.Context, email string) (User, error) {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		return User{Name: "Gosha"}, nil
	case <-ctx.Done():
		return User{}, fmt.Errorf("context canceled")
	}
}

type Handler struct {
	db *DB
}

type Request struct {
	Email string
}

type Response struct {
	User User
}

func (h *Handler) HandleAPI(ctx context.Context, req Request) (Response, error) {
	u, err := h.db.SelectUser(ctx, req.Email)
	if err != nil {
		return Response{}, err
	}

	return Response{User: u}, nil

}

type LabelError struct {
	Label string // метка должна быть в верхнем регистре
	Err   error
}

func (le *LabelError) Error() string {
	return fmt.Sprintf("[%v],%v", le.Label, le.Err)
}
func main() {
	db := DB{}
	handler := Handler{db: &db}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*500))

	// 2. допишите код, который отменяет контекст через 500 миллисекунд
	defer cancel()
	// когда код запустится и успешно выполнится,
	// попробуйте заменить длительность на 2000 миллисекунд

	req := Request{Email: "test@yandex.ru"}
	resp, err := handler.HandleAPI(ctx, req)
	defer fmt.Println(resp, err)
}
