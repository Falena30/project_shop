package middleware

import (
	"fmt"
	"log"
	"net/http"
	"project/shop/data"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ExtractRedisYml(name string, location string) *data.RedisData {
	var redisData data.RedisData
	viper.SetConfigName(name)
	//beritahu lokasinya
	viper.AddConfigPath(location)
	//beri izin viper
	viper.AutomaticEnv()
	//beritahu formtnya
	viper.SetConfigType("yml")
	//check apakah terjadi error pada saat eksekusi
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	//set dauflt
	viper.SetDefault("REDIS.ADDRESS", "defaultHost")
	//tampung nilainya
	Network, ok := viper.Get("REDIS.NETWORK").(string)
	Addr, ok := viper.Get("REDIS.ADDRESS").(string)
	Keypari, ok := viper.Get("REDIS.KEYPAIR").(string)
	Password, ok := viper.Get("REDIS.PASSWORD").(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	//pmasukkan data kedalam struct QUERY
	redisData = data.RedisData{
		ADDRESS:  Addr,
		NETWORK:  Network,
		KEYPAIRS: Keypari,
		PASSWORD: Password,
	}
	//kembalikan nilainya
	return &redisData
}
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionsID := session.Get("id")
		sessionUsername := session.Get("username")
		if sessionsID == nil && sessionUsername == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"pesan": "Tidak Berwenang",
			})
			c.Abort()
		} else {

		}
	}
}
