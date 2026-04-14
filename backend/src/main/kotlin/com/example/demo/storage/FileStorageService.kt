package com.example.demo.storage

import com.example.demo.model.MyData
import org.springframework.stereotype.Service
import org.springframework.beans.factory.annotation.Value
import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.listDirectoryEntries
import java.util.HashMap;
import java.util.UUID

@Service
class FileStorageService(
    @Value("\${IMAGE_ROOT_DIRECTORY}")
    private val imageRootDirectory: String,
) : StorageService {
    val fileNames = mutableListOf<String>()
    init {
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
                        fileNames.add(file.absolutePath)
                    }
                }
        } else {
            println("Directory not found or is not a directory: $imageRootDirectory")
        }
    }
    
	override fun getImage(imageName: String): MyData? {
        val file = File(imageName)
        if (file.exists()) {
            return MyData(file)
        }
        return null
    }
    
	override fun getRandomImageName(): String {
        return fileNames.random()
    }
}
