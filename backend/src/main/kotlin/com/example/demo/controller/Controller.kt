package com.example.demo.controller

import com.example.demo.model.MyData
import com.example.demo.service.ImageService
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
import javax.xml.crypto.Data
import java.io.File
import io.opentelemetry.api.trace.Span
import io.opentelemetry.api.trace.Tracer
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/api/mydata")
class MyDataController(
    private val imageService: ImageService,
    ) {

    val logger = LoggerFactory.getLogger(MyDataController::class.java);

    @GetMapping("/random", produces = [MediaType.IMAGE_JPEG_VALUE])
    fun getRandomImage(): ResponseEntity<ByteArray> {
        logger.info("This is a test log");
        val filename = imageService.getRandomKey()
            ?: return ResponseEntity.notFound().build()
        val data = imageService.getVal(filename)
            ?: return ResponseEntity.notFound().build()
        
        return ResponseEntity.ok()
            .header(HttpHeaders.CONTENT_DISPOSITION, "inline; filename=\"sample.jpg\"")
            .contentType(MediaType.IMAGE_JPEG)
            .body(data)
    }
}