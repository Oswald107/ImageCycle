package com.example.demo.service

import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.listDirectoryEntries
import com.example.demo.model.MyData
import org.springframework.stereotype.Service
import org.springframework.beans.factory.annotation.Value
import java.util.HashMap;
import java.util.UUID


@Service
class DatabaseService(
    @Value("\${IMAGE_ROOT_DIRECTORY}")
    private val imageRootDirectory: String,
) {
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

}