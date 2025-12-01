package com.example.demo.service

import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.listDirectoryEntries
import com.example.demo.model.MyData
import org.springframework.stereotype.Service
import com.example.demo.service.DatabaseService
import org.springframework.core.io.ByteArrayResource
import org.springframework.core.io.Resource
import org.springframework.http.HttpHeaders
import org.springframework.http.MediaType
import org.springframework.http.MediaTypeFactory
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RestController
import java.nio.file.Files
import jakarta.annotation.PostConstruct;
import java.io.ByteArrayOutputStream
import javax.imageio.ImageIO
import java.awt.image.BufferedImage
import redis.clients.jedis.UnifiedJedis;

@Service
class ImageService(
    private var databaseService: DatabaseService,
) {
    val jedis = UnifiedJedis("redis://localhost:6379");

    val files: List<String> = databaseService.getAllFilesFromDir()
    fun getRandom(): MyData? {
        val randomFile = files.random()
        val file = File(randomFile)
        if (file.exists()) {
            return MyData(file)
        }
        return null
    }

    @PostConstruct
    fun storeImagesInCache() {
        for(filename in files) {
            val file = File(filename)
            if (file.exists()) {
                val bytes = imageToByteArray(file)
                addToRedis(filename, bytes)
            } else {
                println("Failed to store image in cache, file does not exist")
            }
        }
    }

    fun addToRedis(filename: String, data: ByteArray) {
        val key = filename.toByteArray()
        jedis.set(key, data);
    }

    fun getRandomKey(): String?{
        val cachedKey = jedis.randomKey()
        if (cachedKey != null)
            return jedis.randomKey()
        println("Failed to retrieve key from cache")

        if (files.size != 0)
            return files.random()
        println("No files available to retrieve")

        return null
    }

    fun getVal(key: String): ByteArray? {
        val data = jedis.get(key.toByteArray())
        if (data != null)
            return data
        println("Failed to retrieve image from cache")

        val file = File(key)
        if (file.exists()) {
            return imageToByteArray(file)
        }
        println("Failed to retrieve image from storage, file doesn't exist")
        return null
    }

    fun imageToByteArray(imageFile: File): ByteArray {
        return Files.readAllBytes(imageFile.toPath())
    }
}