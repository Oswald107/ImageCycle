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
    @GetMapping("/random")
    fun getRandomImage(): ResponseEntity<Resource> {
        val filename = databaseService.getRandomKey()
        // val data = databaseService.getVal(filename)
        // val resource = ByteArrayResource(data)

        // val data = myService.getRandom()
        // val filename = data!!.file.toPath().toString()
        val bytes = Files.readAllBytes(File(filename).toPath())
        val resource = ByteArrayResource(bytes)

        // return ResponseEntity.ok()
        //     .contentLength(10)
        //     .header(HttpHeaders.CONTENT_DISPOSITION, "inline; filename=\"${filename}\"")
        //     .contentType(MediaTypeFactory.getMediaType(filename).orElse(MediaType.APPLICATION_OCTET_STREAM))
        //     .body(resource)

        return ResponseEntity.ok()
            .header(HttpHeaders.CONTENT_DISPOSITION, "inline; filename=\"sample.jpg\"")
            .contentType(MediaType.IMAGE_JPEG)
            .body(resource)
    }
}