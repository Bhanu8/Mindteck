# Mindteck

Postman Collection

https://www.getpostman.com/collections/1188a1f98ea8502f80d3

endpoints:-
    1. upload :- POST Rest API
    2. retrieve :- GET Rest API
    3. search :- POST Rest API

1. Test Cases for Upload service

1.
    Request
        method: post
        body : valid img type file
        endpoint: upload
    Response
        [
            {
                "id": "5fdf97ff6269657d59f18d75",
                "status": 200,
                "message": "success"
            }
        ]
 2. 
    Request
        method: POST
        body : invalid img type file
        endpoint: upload
    Response
        [
            {
                "status": 201,
                "error_message": "Invalid Input File"
            }
        ]

3. 
    Request
        method: GET
        body : valid img type file
        endpoint: upload
    Response
        [
            {
                "status": 401,
                "error_message": "Method Not Allowed"
            }
        ]


Test cases for retrieving images

1. 
    Request
        method: GET
        params : limit
        endpoint: retrieve
    Response
        [
            {
                "image_name": "5fdf97ff6269657d59f18d75.png",
                "image_path": "/images/5fdf97ff6269657d59f18d75.png"
            },
            {
                "image_name": "5fdf974e62696577648cc2f8.png",
                "image_path": "/images/5fdf974e62696577648cc2f8.png"
            }
        ]

2. 

    Request
        method: POST
        params : limit
        endpoint: retrieve
    Response
        [
            {
                "status": 401,
                "error_message": "Method Not Allowed"
            }
        ]

Test cases for search

1. 
    Request
        method: POST
        body : id
        endpoint: search
    Response
        {
            "image_name": "5fdf974e62696577648cc2f8.png",
            "image_path": "/images/5fdf974e62696577648cc2f8.png"
        }

2.

    Request
        method: GET
        body : id
        endpoint: search
    Response
        [
            {
                "status": 401,
                "error_message": "Method Not Allowed"
            }
        ]