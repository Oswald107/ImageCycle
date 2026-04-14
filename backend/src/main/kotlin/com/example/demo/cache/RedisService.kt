package com.example.demo.cache

import redis.clients.jedis.UnifiedJedis
import org.springframework.stereotype.Service

@Service
class RedisService() : CacheService 
{
    val jedis = UnifiedJedis("redis://localhost:6379");
    var curKey = "hello"

	override fun getImageName(): String? {
        return jedis.srandmember(curKey)
    }
	override fun storeImageName(imageName: String): Boolean {
        jedis.sadd(curKey, imageName)
        return true
    }
	override fun createNewSet(imageNames: List<String>): Boolean {
        var nextKey = "hello"
        if (curKey == "hello") {
            nextKey = "bye"
        }

        jedis.sadd(nextKey, *imageNames.toTypedArray())
        jedis.expire(nextKey, 2000)
        curKey = nextKey

        return true
    }
}