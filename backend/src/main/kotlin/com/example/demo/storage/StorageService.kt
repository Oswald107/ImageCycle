package com.example.demo.storage

import com.example.demo.model.MyData

interface StorageService {
	fun getImage(imageName: String): MyData?
	fun getRandomImageName(): String?
}
