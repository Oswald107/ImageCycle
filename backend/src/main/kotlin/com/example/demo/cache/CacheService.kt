package com.example.demo.cache

interface CacheService {
	fun getImageName(): String?
	fun storeImageName(imageName: String): Boolean
	fun createNewSet(imageNames: List<String>): Boolean
}
