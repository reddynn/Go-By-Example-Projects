```
## get all books

POST localhost:8080/books

## get a book

GET localhost:8080/books/1

## create book

POST localhost:8080/books

{
    "Name": "xyz",
    "Author": "xyz",
    "Publication": "xyz",
    "ID": "2"
}

## delete book

DELETE localhost:8080/books/3

## update book

POST localhost:8080/books/1
{
    "Name": "abcd",
    "Author": "xyz",
    "Publication": "xyz",
    "ID": "2"
}

```