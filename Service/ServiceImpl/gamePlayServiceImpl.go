package ServiceImpl

import (
	"context"
	"errors"
	"github.com/arohan007/gameModeStats/Models"
	"sync"
)

type GamePlayServiceImpl struct {
	GameModes   map[string]*Models.GameMode
	RedisClient *redis.Client
	Lock        sync.Mutex
}

func NewInitializeGamePlay() *GamePlayServiceImpl {

	//Initializing Game Modes
	mode1 := &Models.GameMode{Name: "Battle Royal"}
	mode2 := &Models.GameMode{Name: "Team Deathmatch"}
	mode3 := &Models.GameMode{Name: "Capture the flag"}

	var gameModes map[string]*Models.GameMode
	gameModes = make(map[string]*Models.GameMode, 0)
	gameModes["Battle Royal"] = mode1
	gameModes["Team Deathmatch"] = mode2
	gameModes["Capture the flag"] = mode3

	//Initializing Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//Initializing lock
	var lock sync.Mutex

	GamePlayService := GamePlayServiceImpl{
		GameModes:   gameModes,
		RedisClient: redisClient,
		Lock:        lock,
	}

	return &GamePlayService
}

func (g *GamePlayServiceImpl) EnterGamePlay(mode string, user *Models.User) error {

	if _, ok := g.GameModes[mode]; !ok {
		err := errors.New("game mode doesn't exist")
		return err
	}
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.RedisClient.HIncrBy(context.Background(), mode+":"+user.AreaCode, "count", 1)
	user.CurrentMode = mode
	return nil
}

func (g *GamePlayServiceImpl) ExitGamePlay(user *Models.User) error {

	if user.CurrentMode == "" {
		err := errors.New("user doesn't enter into any game mode")
		return err
	}
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.RedisClient.HIncrBy(context.Background(), user.CurrentMode+":"+user.AreaCode, "count",
		1) //todo decrease
	user.CurrentMode = ""
	return nil
}

func (g *GamePlayServiceImpl) CheckActivePlayingModeAreaWise(mode string, areaCode string) (int64, error) {
	_, err := g.RedisClient.HGetAll(context.Background(), mode+":"+areaCode).Result()
	return 0, err
}
