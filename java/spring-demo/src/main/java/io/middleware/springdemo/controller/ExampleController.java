package io.middleware.springdemo.controller;

import io.github.middlewarelabs.agentapmjava.Logger;
import io.middleware.springdemo.dto.RequestBodyExample;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

/**
 * Example controller class.
 */
@RestController
public class ExampleController {

    /**
     * Example of post request.
     *
     * @param requestBodyExample Object of request body.
     * @return Result Object.
     */
    @PostMapping(value = "/post")
    public String postApi(@RequestBody RequestBodyExample requestBodyExample) {
        String requestBody = requestBodyExample.toString();
        Logger.debug(requestBody);
        return "Request body : %s".formatted(requestBody);
    }

    @GetMapping("/error")
    public String getApi(@RequestBody RequestBodyExample requestBodyExample) {
        throw new RuntimeException("Error occurred");
    }
}
