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

@Service
class MyService(
    private var databaseService: DatabaseService,
) {
    val files: List<String> = databaseService.getAllFilesFromDir()
    fun getRandom(): MyData? {
        val randomFile = files.random()
        val file = File(randomFile)
        if (file.exists()) {
            return MyData(file)
        }
        return null
    }

    // @PostConstruct
    fun store() {
        var i = 0
        for(filename in files) {
            if (i==10){
                break
            }
            val file = File(filename)
            if (file.exists()) {
                val bytes = imageToByteArray(file)
                databaseService.addToRedis(filename, bytes)
            }
            i++
        }
    }

    // @PostConstruct
    // fun a() {
    //     val filePath = "/home/henry/Desktop/images/500+ Academic Male Poses/DSC_0593.jpg"
    //     val imageBytes = Files.readAllBytes(File(filePath).toPath())

    //     databaseService.addToRedis(filePath, imageBytes)
    //     println("Stored image '${filePath}' in Redis with key '$filePath'")

    //     val retrievedBytes = databaseService.getVal(filePath)
    //     val outPath = "output.jpg"
    //     Files.write(File(outPath).toPath(), retrievedBytes)
    //     println("Retrieved and wrote image to '$outPath'")
    // }


    fun imageToByteArray(imageFile: File): ByteArray {
        return Files.readAllBytes(imageFile.toPath())
    }
}