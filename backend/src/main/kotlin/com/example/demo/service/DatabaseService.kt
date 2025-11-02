package com.example.demo.service

import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.listDirectoryEntries
import com.example.demo.model.MyData
import org.springframework.stereotype.Service
import org.springframework.beans.factory.annotation.Value
import redis.clients.jedis.UnifiedJedis;
import java.util.HashMap;
import java.util.UUID


@Service
class DatabaseService(
    @Value("\${IMAGE_ROOT_DIRECTORY}")
    private val imageRootDirectory: String,
) {
    val jedis = UnifiedJedis("redis://localhost:6379");

    fun getAllFilesFromDir(): List<String> {
        val output = mutableListOf<String>()
        val directory = File(imageRootDirectory)
        val fileTypes = listOf(
            "jpg",
            "jpeg", 
            "png", 
            "gif", 
            "bmp", 
            "tiff", 
            "webp",
        )
        if (directory.exists() && directory.isDirectory) {
            directory.walk()
                .filter { it.isFile } // Filter to include only actual files, not directories
                .forEach { file ->
                    if (fileTypes.contains(file.extension)) {
                        output.add(file.absolutePath)
                    }
                }
        } else {
            println("Directory not found or is not a directory: $imageRootDirectory")
        }
        return output
    }

    fun addToRedis(filename: String, data: ByteArray) {
        val key = filename.toByteArray()
        jedis.set(key, data);
        println(jedis.get(key)); // >>> OK
    }

    fun getRandomKey(): String{
        return jedis.randomKey()
    }

    fun getVal(key: String): ByteArray {
        return jedis.get(key.toByteArray())
    }
}