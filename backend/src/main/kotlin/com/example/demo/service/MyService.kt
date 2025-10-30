package com.example.demo.service

import java.io.File
import kotlin.io.path.Path
import kotlin.io.path.listDirectoryEntries
import com.example.demo.model.MyData
import org.springframework.stereotype.Component
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

@Component
class MyService(
    private var databaseService: DatabaseService,
) {
    val files: List<String> = databaseService.getAllFilesFromDir()
    fun getRandom(): MyData? {
        val randomFile = files.random()
        // val randomFile = databaseService.getRandom()
        val file = File(randomFile)
        if (file.exists()) {
            return MyData(file)
        }
        return null
    }

    // @PostConstruct
    fun store() {
        for(filename in files) {
            val file = File(filename)
            if (file.exists()) {
                val bytes = Files.readAllBytes(file.toPath())
                databaseService.addToRedis(filename, bytes)
            }
        }
    }
}