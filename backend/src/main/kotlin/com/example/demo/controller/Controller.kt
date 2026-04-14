package com.example.demo.controller

import com.example.demo.cache.CacheService
import com.example.demo.storage.StorageService

import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import org.springframework.stereotype.Component
import org.springframework.core.io.ByteArrayResource
import org.springframework.core.io.Resource
import org.springframework.http.HttpHeaders
import org.springframework.http.MediaType
import org.springframework.http.MediaTypeFactory
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController
import java.nio.file.Files
import java.io.File
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import jakarta.annotation.PostConstruct
import org.springframework.core.io.FileSystemResource

@RestController
class MyDataController(
    private val storageService: StorageService,
    private val cacheService: CacheService
    ) {

    val logger = LoggerFactory.getLogger(MyDataController::class.java);

    @GetMapping("/random", produces = [MediaType.IMAGE_JPEG_VALUE])
    fun getImage(): ResponseEntity<Resource> {
        val filename = cacheService.getImageName()
            ?: return ResponseEntity.notFound().build()
        val data = storageService.getImage(filename)
            ?: return ResponseEntity.notFound().build()
        
        val resource: Resource = FileSystemResource(data.file)

        val contentType = MediaTypeFactory
            .getMediaType(data.file.name)
            .orElse(MediaType.APPLICATION_OCTET_STREAM)

        return ResponseEntity.ok()
            .header(HttpHeaders.CONTENT_DISPOSITION, "inline; filename=\"${data.file.name}\"")
            .contentType(contentType)
            .body(resource)
            
    }

    @PostConstruct
    @Scheduled(cron = "0 */30 * * * *")
    fun refreshCache() {
        val imageNames = mutableSetOf<String>()

        while (imageNames.size < 60) {
            val fileName = storageService.getRandomImageName()
            if (fileName != null) {
                imageNames.add(fileName)
            } else {
                logger.warn("Failed to fetch image name from storage")
            }
        }

        cacheService.createNewSet(imageNames.toList())
    }
}