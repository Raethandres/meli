echo "test 1"
curl --location --request POST 'http://localhost:8080/topsecret/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "satellites": [
        {
            "name": "kenobi",
            "distance": 485.41,
            "message": ["este", "", "", "mensaje", ""]
        },
        {
            "name": "skywalker",
            "distance": 265.75,
            "message": ["", "es", "", "", "secreto"]
        },
        {
            "name": "sato",
            "distance": 600.52,
            "message": ["este", "", "un", "", ""]
        }
    ]
}'

echo "test 3"
curl --location --request POST 'http://localhost:8080/topsecret_split/kenobi' \
--header 'Content-Type: application/json' \
--data-raw '{
    "distance": 500.41,
    "message": ["dsfdg", "", "", "mensaje", ""]
}'

echo "test 4"
curl --location --request GET 'http://localhost:8080/topsecret_split/'