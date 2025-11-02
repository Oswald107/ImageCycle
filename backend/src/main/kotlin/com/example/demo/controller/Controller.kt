package com.example.demo.controller

import com.example.demo.model.MyData
import com.example.demo.service.MyService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*

import org.springframework.stereotype.Component
import com.example.demo.service.DatabaseService
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

@RestController
@RequestMapping("/api/mydata")
class MyDataController(
    private val myService: MyService,
    private val databaseService: DatabaseService,
    ) {
    @GetMapping("/random", produces = [MediaType.IMAGE_JPEG_VALUE])
    fun getRandomImage(): ResponseEntity<ByteArray> {
        val filename = databaseService.getRandomKey()
        val data = databaseService.getVal(filename)
        
        return ResponseEntity.ok()
            .header(HttpHeaders.CONTENT_DISPOSITION, "inline; filename=\"sample.jpg\"")
            .contentType(MediaType.IMAGE_JPEG)
            .body(data)
    }
}