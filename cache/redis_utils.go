package cache

import (
	"github.com/go-redis/redis/v7"
	"gitlab.onedao.com.cn/onedao/backend/backend/utils-go/utils"
	"strconv"
	"time"
)

func RedisSetNx(redisClient *redis.Client, key string, value interface{}) bool {
	return redisClient.SetNX(key, value, 200*time.Second).Val()
}

func RedisSet(redisClient *redis.Client, key string, value interface{}, duration time.Duration) error {

	userByte, _ := utils.GobEncode(value)
	return redisClient.Set(key, userByte, duration).Err()
}

func RedisDel(redisClient *redis.Client, key string) error {
	return redisClient.Del(key).Err()
}

func RedisGet(redisClient *redis.Client, key string, result interface{}) error {
	value, err := redisClient.Get(key).Bytes()
	if err != nil {
		return err
	}
	err = utils.GobDeCode(value, result)
	return err
}

func RedisSetInt(redisClient *redis.Client, key string, value int64, duration time.Duration) error {
	return redisClient.Set(key, value, duration).Err()
}

func RedisGetInt(redisClient *redis.Client, key string) int64 {
	res := redisClient.Get(key).Val()
	resultInt, _ := strconv.ParseInt(res, 10, 64)
	return resultInt
}

func RedisIncrease(redisClient *redis.Client, key string, incr int64) {
	redisClient.IncrBy(key, incr)
}

func RedisDecrease(redisClient *redis.Client, key string, incr int64) {
	redisClient.DecrBy(key, incr)
}

func RedisExist(redisClient *redis.Client, key string) bool {
	i := redisClient.Exists(key).Val()
	return i > 0
}

func RedisSetMember(redisClient *redis.Client, key string, member interface{}) {
	redisClient.SAdd(key, member)
}

func RedisSetMemberExist(redisClient *redis.Client, key string, member interface{}) bool {
	return redisClient.SIsMember(key, member).Val()
}

func RedisSetMemberRemove(redisClient *redis.Client, key string, member interface{}) {
	redisClient.SRem(key, member)
}

//新增、更新有序集合中的元素
func RedisZSetAdd(redisClient *redis.Client, key string, member, score int64) {
	z := redis.Z{
		Score:  float64(score),
		Member: member,
	}
	redisClient.ZAdd(key, &z)
}

//根据score获取列表
func RedisZSetRange(redisClient *redis.Client, key string, start, stop int64) []redis.Z {
	//list := redisClient.ZRevRangeWithScores(key, start, stop).Val()
	//members := make([]interface{}, 0)
	//for _, li := range list {
	//
	//}

	return redisClient.ZRevRangeWithScores(key, start, stop).Val()
}
